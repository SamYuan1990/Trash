package main

import (
	"net/http"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	/*Context("single run", func() {
		It("runs", func() {
			address0 := "1000"
			LPort := "1000"
			cmd := exec.Command(Bin, address0, LPort)
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

	Context("2 nodes", func() {
		It("runs", func() {
			address0 := "1000"
			LPort := "1001"
			cmd := exec.Command(Bin, address0, LPort)
			Session0, err = gexec.Start(cmd, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address1 := "1001"
			cmd1 := exec.Command(Bin, address1, LPort)
			Session1, err = gexec.Start(cmd1, nil, nil)
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(5 * time.Second)

			resp, err := http.Get("http://127.0.0.1:1000/data?data=123")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Eventually(Session0.Out).Should(Say("123"))
			Eventually(Session1.Out).Should(Say("123"))
		})
	})

	Context("3 nodes", func() {
		It("runs", func() {
			address0 := "1000"
			LPort := "1002"
			cmd := exec.Command(Bin, address0, LPort)
			Session0, err = gexec.Start(cmd, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address1 := "1001"
			cmd1 := exec.Command(Bin, address1, LPort)
			Session1, err = gexec.Start(cmd1, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address2 := "1002"
			cmd2 := exec.Command(Bin, address2, LPort)
			Session2, err = gexec.Start(cmd2, nil, nil)
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(5 * time.Second)

			resp, err := http.Get("http://127.0.0.1:1000/data?data=123")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Eventually(Session0.Out).Should(Say("123"))
			Eventually(Session1.Out).Should(Say("123"))
			Eventually(Session2.Out).Should(Say("123"))
		})
	})*/

	Context("5 nodes", func() {
		It("runs", func() {
			address0 := "1000"
			cmd := exec.Command(Bin, address0)
			Session0, err = gexec.Start(cmd, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address1 := "1001"
			cmd1 := exec.Command(Bin, address1)
			Session1, err = gexec.Start(cmd1, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address2 := "1002"
			cmd2 := exec.Command(Bin, address2)
			Session2, err = gexec.Start(cmd2, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address3 := "1003"
			cmd3 := exec.Command(Bin, address3)
			Session3, err = gexec.Start(cmd3, nil, nil)
			Expect(err).NotTo(HaveOccurred())

			address4 := "1004"
			cmd4 := exec.Command(Bin, address4)
			Session4, err = gexec.Start(cmd4, nil, nil)
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(5 * time.Second)

			resp, err := http.Get("http://127.0.0.1:1000/data?data=123")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Eventually(Session0.Out).Should(Say("123"))
			Eventually(Session1.Out).Should(Say("123"))
			Eventually(Session2.Out).Should(Say("123"))
			Eventually(Session3.Out).Should(Say("123"))
			Eventually(Session4.Out).Should(Say("123"))
		})
	})
})
