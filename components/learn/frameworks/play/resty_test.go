package play_test

import (
	db2 "github.com/amanhigh/go-fun/models/fun-app/db"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net"
	"net/http"
	"time"
)

type BinAnyResponse struct {
	Headers map[string]string `json:"headers"`
	//Data db.Person `json:"data"`
	Method string `json:"method"`
}

var _ = Describe("Resty", func() {

	var (
		client = resty.New()
		err    error
		resp   *resty.Response
	)

	BeforeEach(func() {
		err = nil
		resp = nil

		//Custom Transport
		transport := http.Transport{
			DisableCompression: false,
			DisableKeepAlives:  false,
			DialContext: (&net.Dialer{
				Timeout: time.Second, // Connect Timeout
			}).DialContext,
			IdleConnTimeout:     time.Minute, //Idle Timeout Before Closing Keepalive Connection
			MaxIdleConnsPerHost: 10,
		}
		client.SetTransport(&transport)

		//client.SetDebug(true)
		//client.EnableTrace()
		client.SetHeader("Content-Type", "application/json")
		client.SetHeaderVerbatim("Foo", "bar")
		client.SetTimeout(2 * time.Second)
		client.SetHostURL("https://httpbin.org")

		//Try 2 times at interval of one second, max time 5 Seconds
		client.SetRetryCount(5).
			SetRetryWaitTime(300 * time.Millisecond).
			SetRetryMaxWaitTime(3 * time.Second)
	})

	It("should build", func() {
		Expect(client).To(Not(BeNil()))
	})

	It("should do get", func() {
		resp, err = client.R().SetQueryParam("foo", "bar").Get("/status/200")
		Expect(err).To(BeNil())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))
	})

	Context("Custom Request", func() {
		var (
			person      db2.Person
			binResponse BinAnyResponse
			headerKey   = "Myheader"
			headerValue = "MyHeaderValue"
		)

		BeforeEach(func() {
			binResponse = BinAnyResponse{}
			person = db2.Person{
				Name:   "Aman",
				Age:    18,
				Gender: "Male",
			}
		})

		It("should build custom Request", func() {
			resp, err = client.R().
				SetHeader(headerKey, headerValue).
				SetBody(person).
				SetResult(&binResponse).
				Post("/anything")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(binResponse.Method).To(Equal(http.MethodPost))
			Expect(len(binResponse.Headers)).To(BeNumerically(">", 2))
			Expect(binResponse.Headers).To(HaveKeyWithValue(headerKey, headerValue))
			Expect(binResponse.Headers).To(HaveKeyWithValue("Foo", "bar"), "Having Global Headers")
		})
	})

})
