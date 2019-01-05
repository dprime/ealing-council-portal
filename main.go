package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"net/url"
	"time"
	"net/http/cookiejar"

	"strings"
	"log"
	"strconv"
)

var mainPage = "https://www.ealing.gov.uk/site/custom_scripts/waste_collection/waste_collection.aspx"

/*
WARNING! This code does not work without modifications to net/http/cookie.go to allow broken cookie names.
Ealing council use some shitty MS stack that spits out cookies with { and } in their names
This issue prevents it from actually working, without bypassing the isCookieNameValid check:
https://github.com/golang/go/issues/29580
 */

func main() {
	jar, _ := cookiejar.New(nil)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Jar: jar,
	}

	resp, err := httpClient.Get(mainPage)

	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)

		//steal the tokens
		viewState, viewStateGenerator, eventValidation := extractFormInfo(string(content))

		fmt.Println(viewState, viewStateGenerator, eventValidation)

		addresses := loadAddresses("W5 1AA", viewState, viewStateGenerator, eventValidation, httpClient)

		log.Println(addresses)
	}
}

func loadAddresses(postcode, viewState, viewStateGenerator, eventValidation string, httpClient *http.Client) []string {

	urlValues := url.Values{
		"__VIEWSTATE": {viewState},
		"__VIEWSTATEGENERATOR": {viewStateGenerator},
		"__EVENTVALIDATION": {eventValidation},
		"hiddenLookup": {"LOOKUP"},
		"txtLookupPostCode": {postcode},
		"AddressLookup": {"Address lookup"},
		"selectLookup": {""},
	}

	encoded := urlValues.Encode()

	req, _ := http.NewRequest("POST", mainPage, strings.NewReader(encoded))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:64.0) Gecko/20100101 Firefox/64.0")


	resp, err := httpClient.Do(req)

	fmt.Println(resp)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
	//fmt.Println(len(content))

	return []string{}
}

func addressPost(postcode, viewState, viewStateGenerator, eventValidation string, httpClient *http.Client) {

	urlValues := url.Values{
		"__VIEWSTATE": {viewState},
		"__VIEWSTATEGENERATOR": {viewStateGenerator},
		"__EVENTVALIDATION": {eventValidation},
		"hiddenLookup": {"LOOKUP"},
		"txtLookupPostCode": {postcode},
		"AddressLookup": {"Address+lookup"},
		"selectLookup": {},
	}

	req, _ := http.NewRequest("POST", mainPage, strings.NewReader(urlValues.Encode()))

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	resp, err := httpClient.Do(req)

	fmt.Println(resp)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}

func extractFormInfo(content string) (viewState, viewStateGenerator, eventValidation string) {
	viewStateRegex := `id="__VIEWSTATE"\svalue="([^"]+)"\s\/>`
	viewStateGeneratorRegex := `id="__VIEWSTATEGENERATOR" value="([^"]+)" \/>`
	eventValidationRegex := `id="__EVENTVALIDATION" value="([^"]+)" \/>`

	viewState = getMG1FromRegex(content, viewStateRegex)
	viewStateGenerator = getMG1FromRegex(content, viewStateGeneratorRegex)
	eventValidation = getMG1FromRegex(content, eventValidationRegex)

	return
}

func getMG1FromRegex(haystack, regex string) string {
	r, _ := regexp.Compile(regex)

	matches := r.FindStringSubmatch(haystack)

	return matches[1]
}