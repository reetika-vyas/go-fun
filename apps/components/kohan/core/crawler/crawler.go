package crawler

import (
	"context"
	"fmt"
	util2 "github.com/amanhigh/go-fun/apps/common/util"
	"github.com/fatih/color"
	"github.com/wesovilabs/koazee"
	"io/ioutil"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	. "github.com/amanhigh/go-fun/apps/models/crawler"
	"github.com/amanhigh/go-fun/util"
)

const (
	GOOD_URL_FILE = "/tmp/good.txt"
	BAD_URL_FILE  = "/tmp/bad.txt"
	BUFFER_SIZE   = 512
)

type Crawler interface {
	GatherLinks(page *util2.Page, ch chan CrawlInfo)
	NextPageLink(page *util2.Page) (string, bool)
	PrintSet(good []CrawlInfo, bad []CrawlInfo) bool
	GetTopPage() *util2.Page
}

type CrawlerManager struct {
	Crawler    Crawler
	ctx        context.Context
	cancelFunc context.CancelFunc

	verbose bool

	/* Counts to track collected & required */
	collected int32
	required  int32

	infoChannel chan CrawlInfo
	goodInfo    []CrawlInfo
	badInfo     []CrawlInfo

	/* Concurrency Control */
	semaphoreChannel chan int
}

func NewCrawlerManager(crawler Crawler, requiredCount int, verbose bool) *CrawlerManager {
	return &CrawlerManager{
		Crawler:          crawler,
		required:         int32(requiredCount),
		infoChannel:      make(chan CrawlInfo, BUFFER_SIZE),
		verbose:          verbose,
		semaphoreChannel: make(chan int, runtime.NumCPU()),
	}
}

func (self *CrawlerManager) Crawl() {
	util.PrintYellow(fmt.Sprintf("Crawling RequiredLinks:%v Cores: %v", self.required, runtime.NumCPU()))

	/* Fire First Crawler */
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go self.crawlRecursive(self.Crawler.GetTopPage(), waitGroup)

	/* Collect & Organise Crawled Links */
	go self.BuildSet()

	/* Wait for all Crawlers to Return */
	waitGroup.Wait()
	close(self.infoChannel)

	/* Print Organised Links */
	self.PrintSet(self.goodInfo, self.badInfo)
}

func (self *CrawlerManager) BuildSet() {
	/* Fire Parallel Consumer to Separate Movies */
	for info := range self.infoChannel {
		if info.GoodBad() == nil {
			self.goodInfo = append(self.goodInfo, info)
			atomic.AddInt32(&self.collected, 1)
		} else {
			self.badInfo = append(self.badInfo, info)
		}
	}
}

func (self *CrawlerManager) PrintSet(good []CrawlInfo, bad []CrawlInfo) {
	/* Check if Crawler want us to print or already has printed required info */
	if ok := self.Crawler.PrintSet(good, bad); ok {
		/* Output Good/Bad Info in Separate Sections */
		color.Green("Passed Info: %v", len(good))
		self.printWriteCrawledInfo(good, GOOD_URL_FILE)

		color.Red("Failed Info: %v", len(bad))
		self.printWriteCrawledInfo(bad, BAD_URL_FILE)
	}
}

/**
Print Info using interface and write extracted links to
GOOD/BAD Files for Chrome Processing
*/
func (self *CrawlerManager) printWriteCrawledInfo(infos []CrawlInfo, filePath string) {
	var urls []string
	for _, info := range infos {
		urls = append(urls, info.ToUrl()...)
	}
	urls = koazee.StreamOf(urls).RemoveDuplicates().Out().Val().([]string)
	urlDump := strings.Join(urls, "\n")
	if self.verbose {
		fmt.Println(urlDump)
	}
	ioutil.WriteFile(filePath, []byte(urlDump), util.DEFAULT_PERM)
}

/**
Recursively Crawl Given Page moving to next if next link is available.
Write all Movies of current page onto channel
*/
func (self *CrawlerManager) crawlRecursive(page *util2.Page, waitGroup *sync.WaitGroup) {
	/* Aquire Grant */
	self.semaphoreChannel <- 1
	collected := atomic.LoadInt32(&self.collected)

	if collected < self.required {
		color.Yellow("Processing: %v Collected: %v", page.Document.Url.String(), collected)
		/* If Next Link is Present Crawl It */
		if link, ok := self.Crawler.NextPageLink(page); ok {
			waitGroup.Add(1)
			go self.crawlRecursive(util2.NewPage(link), waitGroup)
		}
		/* Find Links for this Page */
		self.Crawler.GatherLinks(page, self.infoChannel)
	}

	/* Release Grant */
	<-self.semaphoreChannel
	waitGroup.Done()
}
