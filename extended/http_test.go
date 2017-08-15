package extended_test

import (
	myHttp "github.com/hongkailiu/svt-go/http"
	"net/http"
	"math/rand"
	"gopkg.in/resty.v0"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
	"encoding/json"
	"github.com/onsi/ginkgo/config"
)

var (
	server myHttp.Server
	port int
)


var _ = BeforeSuite(func() {
	port = rand.Intn(100) + 10000 + config.GinkgoConfig.ParallelNode
	By(fmt.Sprintf("Starting http server with port %d ...", port))
	server := myHttp.Server{Port:port}
	go server.Run()
})

var _ = AfterSuite(func() {
	By("Stopping http server ...")
	server.Stop()
})

var _ = Describe("Http", func() {

	var (
		resp *resty.Response
		err error
	)

	BeforeEach(func() {


	})

	AfterEach(func() {

	})

	Describe("Root handler", func() {
		Context("Without parameters", func() {

			It(fmt.Sprintf("should return %d", http.StatusOK), func() {
				resp, err = resty.R().Get(fmt.Sprintf("http://localhost:%d", port))
				Expect(http.StatusOK).To(Equal(resp.StatusCode()))
				Expect(err).NotTo(HaveOccurred())
				c := make(map[string]interface{})
				err = json.Unmarshal(resp.Body(), &c)
				Expect(err).NotTo(HaveOccurred())
				Expect(c["version"]).To(ContainSubstring("no such file or directory"))
				Expect(c["now"]).NotTo(BeNil())
				Expect(c["ips"]).NotTo(BeEmpty())
				//Expect(c["ips"]).To(ContainElement("127.0.0.1"))
			})
		})

	})

	Describe("Log handler", func() {
		Context("Using get method", func() {

			It(fmt.Sprintf("should return %d", http.StatusNotFound), func() {

				resp, err = resty.R().Get(fmt.Sprintf("http://localhost:%d/logs", port))
				GinkgoWriter.Write(resp.Body())
				Expect(http.StatusNotFound).To(Equal(resp.StatusCode()))
				Expect(err).NotTo(HaveOccurred())

			})
		})

	})

})
