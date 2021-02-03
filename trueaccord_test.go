package trueaccord_test

import (
	"fmt"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrueAccord", func() {
	It("Should produce expected output", func() {
		cmd := exec.Command("go", "run", "cmd/trueaccord/main.go")
		stdoutStderr, err := cmd.CombinedOutput()
		Expect(err).To(BeNil())
		fmt.Printf("%s\n", stdoutStderr)
		lines := strings.Split(strings.TrimSpace(string(stdoutStderr)), "\n")
		Expect(lines).To(HaveLen(5))
		Expect(lines[0]).To(Equal(`{"id":0,"amount":123.46,"is_in_payment_plan":true,"remaining_amount":20.96,"next_payment_due_date":"2020-11-05"}`))
		Expect(lines[1]).To(Equal(`{"id":1,"amount":100,"is_in_payment_plan":true,"remaining_amount":50,"next_payment_due_date":"2020-08-15"}`))
		Expect(lines[2]).To(Equal(`{"id":2,"amount":4920.34,"is_in_payment_plan":true,"remaining_amount":607.6699,"next_payment_due_date":"2020-08-22"}`))
		Expect(lines[3]).To(Equal(`{"id":3,"amount":12938,"is_in_payment_plan":true,"remaining_amount":9247.745,"next_payment_due_date":"2020-08-22"}`))
		Expect(lines[4]).To(Equal(`{"id":4,"amount":9238.02,"is_in_payment_plan":false,"remaining_amount":9238.02,"next_payment_due_date":null}`))
	})
})
