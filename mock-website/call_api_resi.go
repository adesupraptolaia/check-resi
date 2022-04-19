package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func cekResi() (resp []byte) {

	url := os.Getenv("URL_CHECK_RESI")
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf("_token=%s&resi=%s&text=%s",
			os.Getenv("TOKEN_CHECK_RESI"),
			os.Getenv("RESI_CEHCK_RESI"),
			os.Getenv("TEXT_CHECK_RESI")),
	)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", os.Getenv("COOKIE_CHECK_RESI"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp = body

	return
}
