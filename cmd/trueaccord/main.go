package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type debt struct {
	Id     int     `json:"id"`
	Amount float32 `json:"amount"`
}

type paymentPlan struct {
	Amount               float32 `json:"amount_to_pay"`
	DebtId               int     `json:"debt_id"`
	Id                   int     `json:"id"`
	InstallmentAmount    float32 `json:"installment_amount"`
	InstallmentFrequency string  `json:"installment_frequency"`
	StartDate            string  `json:"start_date"`
}

type payment struct {
	Amount        float32 `json:"amount"`
	Date          string  `json:"date"`
	PaymentPlanId int     `json:"payment_plan_id"`
}

func main() {
	resp, err := http.Get("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(respBody))

	var debts []debt
	if err := json.Unmarshal(respBody, &debts); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", debts)

	resp, err = http.Get("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(respBody))

	var paymentPlans []paymentPlan
	if err := json.Unmarshal(respBody, &paymentPlans); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", paymentPlans)

	resp, err = http.Get("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(respBody))

	var payments []payment
	if err := json.Unmarshal(respBody, &payments); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", payments)

	type outputDebt struct {
		Id                 int     `json:"id"`
		Amount             float32 `json:"amount"`
		IsInPaymentPlan    bool    `json:"is_in_payment_plan"`
		RemainingAmount    float32 `json:"remaining_amount"`
		NextPaymentDueDate *string `json:"next_payment_due_date"`
	}

	for _, debt := range debts {
		od := outputDebt{
			Id:              debt.Id,
			Amount:          debt.Amount,
			IsInPaymentPlan: debt.isInPaymentPlan(paymentPlans),
			RemainingAmount: debt.Amount,
		}
		if od.IsInPaymentPlan {
			pp := debt.findPaymentPlan(paymentPlans)
			foundPayments := pp.findPayments(payments)
			od.RemainingAmount = debt.remainingAmount(pp, foundPayments)
			nextPDD := pp.nextPaymentDueDate(foundPayments)
			if !(time.Time{}).Equal(nextPDD) {
				npdd := nextPDD.Format("2006-01-02")
				od.NextPaymentDueDate = &npdd
			}
		}

		out, err := json.Marshal(od)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", string(out))
	}
}

func (d *debt) isInPaymentPlan(paymentPlans []paymentPlan) bool {
	return d.findPaymentPlan(paymentPlans) != nil
}

func (d *debt) findPaymentPlan(paymentPlans []paymentPlan) *paymentPlan {
	for _, pp := range paymentPlans {
		if d.Id == pp.DebtId {
			return &pp
		}
	}

	return nil
}

func (pp *paymentPlan) findPayments(allPayments []payment) []payment {
	var foundPayments []payment
	for _, p := range allPayments {
		if pp.Id == p.PaymentPlanId {
			foundPayments = append(foundPayments, p)
		}
	}

	return foundPayments
}

func (d *debt) remainingAmount(pp *paymentPlan, payments []payment) (remainingAmount float32) {
	remainingAmount = d.Amount
	for _, p := range payments {
		remainingAmount = remainingAmount - p.Amount
	}
	return
}

func mostRecentPayment(payments []payment) (mostRecentPayment time.Time) {
	for _, p := range payments {
		var t time.Time
		var err error
		if t, err = time.Parse("2006-01-02", p.Date); err != nil {
			log.Fatal(err)
		}
		if t.After(mostRecentPayment) {
			mostRecentPayment = t
		}
	}
	return
}

func (pp *paymentPlan) nextPaymentDueDate(payments []payment) time.Time {
	switch pp.InstallmentFrequency {
	case "WEEKLY":
		return mostRecentPayment(payments).Add(7 * 24 * time.Hour)
	case "BI_WEEKLY":
		return mostRecentPayment(payments).Add(14 * 24 * time.Hour)
	}
	return time.Time{} // should not happen
}
