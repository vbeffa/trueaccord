package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
	fmt.Println(string(respBody))

	type debt struct {
		Id     int     `json:"id"`
		Amount float32 `json:"amount"`
	}

	var debts []debt
	if err := json.Unmarshal(respBody, &debts); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", debts)

	resp, err = http.Get("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))

	type paymentPlan struct {
		Amount               float32 `json:"amount_to_pay"`
		DebtId               int     `json:"debt_id"`
		Id                   int     `json:"id"`
		InstallmentAmount    float32 `json:"installment_amount"`
		InstallmentFrequency string  `json:"installment_frequency"`
		StartDate            string  `json:"start_date"`
	}

	var paymentPlans []paymentPlan
	if err := json.Unmarshal(respBody, &paymentPlans); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", paymentPlans)

	resp, err = http.Get("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))

	type payment struct {
		Amount        float32 `json:"amount"`
		Date          string  `json:"date"`
		PaymentPlanId int     `json:"payment_plan_id"`
	}

	var payments []payment
	if err := json.Unmarshal(respBody, &payments); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", payments)
}
