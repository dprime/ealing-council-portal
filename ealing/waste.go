package ealing

import (
	"regexp"
	"strings"
	"net/http/cookiejar"
	"net/http"
	"time"
	"io/ioutil"
	"net/url"
	"strconv"
	"log"
	"github.com/PuerkitoBio/goquery"
)

var mainPage = "https://www.ealing.gov.uk/site/custom_scripts/waste_collection/waste_collection.aspx"
var dateLayout = "02/01/06"


type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Jar: jar,
	}
	return &Client{
		httpClient,
	}
}

func (c *Client) LoadAddressesForPostCode(postcode string) []*Address{


	resp, err := c.httpClient.Get(mainPage)

	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)

		//steal the tokens
		viewState, viewStateGenerator, eventValidation := extractFormInfo(string(content))

		addresses := c.loadAddresses(postcode, viewState, viewStateGenerator, eventValidation)
		return addresses
	}
}

func (c *Client) loadAddresses(postcode, viewState, viewStateGenerator, eventValidation string) []*Address {

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

	addHeadersForPost(req, len(encoded))

	resp, err := c.httpClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)


	parsed := parseAddresses(string(content))

	return parsed
}



func (c *Client) LoadCollectionsForAddres(address *Address) map[string][]*Collection {

	/*
			hiddenLookup	COLLECTION
			hiddenUPRN	12005866++++++++
			hiddenAddress	FLAT+3+FIRST+FLOOR,+BURLINGTON+HOUSE,+1+3,+THE+COMMON,+ALL,+EALING,+W5
			hiddenTempScheduleAvailable
			hiddenCollectionCode
			hiddenDocPath	https://www.ealing.gov.uk/site/custom_scripts/waste_collection/docs/
			hiddenDocName	WC_Calendar_1218_
			hiddenDocExt	.pdf
			hiddenTimedCollectionCode
	 */

	urlValues := url.Values{
		"hiddenLookup": {"COLLECTION"},
		"hiddenUPRN": {address.UPRN},
		"hiddenAddress": {address.Address},
		//"hiddenTempScheduleAvailable": {""},
		//"hiddenCollectionCode": {""},
		//"hiddenDocPath": {"https://www.ealing.gov.uk/site/custom_scripts/waste_collection/docs/"},
		//"hiddenDocName": {"WC_Calendar_1218_"},
		//"hiddenDocExt": {".pdf"},
		//"hiddenTimedCollectionCode": {""},
	}

	encoded := urlValues.Encode()

	req, _ := http.NewRequest("POST", mainPage, strings.NewReader(encoded))

	addHeadersForPost(req, len(encoded))

	resp, err := c.httpClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	collections := make(map[string][]*Collection)

	// Find the review items
	doc.Find("#RetrieveAllDataGrid").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(i int, row *goquery.Selection) {
			if i == 0 {
				return //skip the title row
			}
			collectionType := ""
			dates := ""
			frequency := ""

			row.Find("td").Each(func(i int, cell *goquery.Selection) {
				if i == 0 {
					collectionType = cell.Text()
				} else if i == 1 {
					dates = cell.Text()
				} else if i == 2 {
					frequency = cell.Text()
				}
			})
			var collectionArr []*Collection

			if len(dates) % 8 != 0 {
				panic("date string invalid " + dates)
			} else {
				for i := 0; i < len(dates); i+= 8 {
					chunk := dates[i:i+8]
					t, _ := time.Parse(dateLayout, chunk)
					collectionArr = append(collectionArr, &Collection{
						Type: collectionType,
						Frequency: frequency,
						Date: t,
					})
				}
			}
			collections[collectionType] = collectionArr
		})
	})


	return collections

}

type Collection struct {
	Type string
	Frequency string
	Date time.Time
}

func extractFormInfo(content string) (viewState, viewStateGenerator, eventValidation string) {
	viewStateRegex := `id="__VIEWSTATE"\svalue="([^"]+)"\s\/>`
	viewStateGeneratorRegex := `id="__VIEWSTATEGENERATOR" value="([^"]+)" \/>`
	eventValidationRegex := `id="__EVENTVALIDATION" value="([^"]+)" \/>`

	viewState = getMatchesFromRegex(content, viewStateRegex)[1]
	viewStateGenerator = getMatchesFromRegex(content, viewStateGeneratorRegex)[1]
	eventValidation = getMatchesFromRegex(content, eventValidationRegex)[1]

	return
}


func getMatchesFromRegex(haystack, regex string) []string {
	r, _ := regexp.Compile(regex)

	matches := r.FindStringSubmatch(haystack)

	return matches
}

type Address struct {
	UPRN    string
	Address string
}

func parseAddresses(addressResponse string) []*Address{

	//need to find this line: <select name="selectLookup" id="selectLookup" onchange="getSelectedAddress()">
	//then keep going til we get to this line </select>
	//and build Address records from each line between

	lines := strings.Split(addressResponse, "\n")

	zone := false

	var res []*Address

	for _, line := range lines {
		if strings.Contains(line, `<select name="selectLookup" id="selectLookup" onchange="getSelectedAddress()">`) {
			zone = true
		}

		//look like this <option value="12002978        ">36, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>

		if zone {
			if strings.Contains(line, "<option value=") {
				uprn, address := ParseAddressOptionLine(line)
				res = append(res, &Address{UPRN: uprn, Address: address})
			}
		}

		if strings.Contains(line, `</select>`) {
			break
		}
	}

	return res
}

func ParseAddressOptionLine(line string) (uprn, address string) {
	r := `<option value="([\d\s]+)">([^<]+)<`

	matches := getMatchesFromRegex(line, r)

	if len(matches) == 3 {
		return matches[1], matches[2]
	}

	return "", ""
}

func addHeadersForPost(req *http.Request, length int) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(length))
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:64.0) Gecko/20100101 Firefox/64.0")
}