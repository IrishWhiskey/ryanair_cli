package ryanair

import (
	"time"
	"net/http"
	"io/ioutil"
)

var client = &http.Client{Timeout: 10 * time.Second}

func getUrl(departure string, arrival string, date string) string {
	var url = "https://desktopapps.ryanair.com/v4/en-gb/availability"
	url += "?ADT=1&CHD=0&TEEN=0&RoundTrip=false&ToUs=AGREED"
	url += "&DateIn=&DateOut=" + EncodeDate(date)
	url += "&Destination=" + arrival + "&Origin=" + departure
	url += "&FlexDaysOut=0&FlexDaysBeforeOut=0"
	return url
}

func getBody(url string) []byte {
	req, err := client.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	return []byte(body)
}

func Query(departure string, arrival string, date string) (*Travel_t, error) {
	url := getUrl(departure, arrival, date)

	body := getBody(url)

	ans := parseJson(body)

	return ans, nil
}
