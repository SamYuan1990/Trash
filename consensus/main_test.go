package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	Context("single run", func() {
		It("runs", func() {
			address0 := "0.0.0.0:1000"
			cmd := exec.Command(Bin, address0)
			Session0, err = gexec.Start(cmd, nil, nil)
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(5 * time.Second)
			resp, err := http.Get("http://127.0.0.1:1000/data?data=123")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Eventually(Session0.Out).Should(Say("123"))

			data := &ConsensusData{
				Data: "234",
			}
			d, err := json.Marshal(data)
			body := strings.NewReader(string(d))
			resp, err = http.Post("http://127.0.0.1:1000/consensus", "application/x-www-form-urlencoded", body)
			Eventually(Session0.Out).Should(Say("234"))
		})
	})
})
