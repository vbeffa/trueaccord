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
}
