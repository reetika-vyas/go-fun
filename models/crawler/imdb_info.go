package crawler

import (
	"errors"
	"fmt"
	helper2 "github.com/amanhigh/go-fun/common/helper"
	"github.com/fatih/color"
)

type ImdbInfo struct {
	Name     string
	Link     string
	Language string
	Rating   float64
	MyRating float64
	CutOff   int
}

func (self *ImdbInfo) Print() {
	if self.MyRating != -1 {
		color.White(fmt.Sprintf("%v: %.2f/%.2f - %v", self.Name, self.MyRating, self.Rating, self.Link))
	} else {
		color.White(fmt.Sprintf("%v: %.2f - %v", self.Name, self.Rating, self.Link))
	}
}

func (self *ImdbInfo) GoodBad() error {
	if self.Rating < float64(self.CutOff) {
		return errors.New(fmt.Sprintf("Subpar Rating %v < %v", self.Rating, self.CutOff))
	} else if self.MyRating != -1 {
		return errors.New(fmt.Sprintf("Movie already Rated: %v", self.MyRating))
	}
	return nil
}

func (self *ImdbInfo) ToUrl() []string {
	return []string{self.Link, helper2.YoutubeSearch(self.Name + " Trailer"), self.getDownloadLink()}
}

func (self *ImdbInfo) getDownloadLink() string {
	switch self.Language {
	case "en":
		return helper2.YtsSearch(self.Name)
	case "hi":
		return helper2.TSearch(self.Name)
	default:
		return helper2.YoutubeSearch(self.Name + " Full Movie")
	}

}
