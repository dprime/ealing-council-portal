package main

import (
	"os"
	"github.com/dprime/ealing-council-portal/ealing"
	"bufio"
	"fmt"
	"strconv"
)


/*
WARNING! This code does not work without modifications to net/http/cookie.go to allow broken cookie names.
Ealing council use some shitty MS stack that spits out cookies with { and } in their names
This issue prevents it from actually working, without bypassing the isCookieNameValid check:
https://github.com/golang/go/issues/29580

Patch:

Index: cookie.go
IDEA additional info:
Subsystem: com.intellij.openapi.diff.impl.patch.CharsetEP
<+>UTF-8
===================================================================
--- cookie.go	(date 1546702070087)
+++ cookie.go	(date 1546702070087)
@@ -382,6 +382,9 @@
 }

 func isCookieNameValid(raw string) bool {
+	if true { //TODO DAVE HACK
+		return true
+	}
 	if raw == "" {
 		return false
 	}


 */

func main() {

	client := ealing.NewClient()

	postCode := getPostcode()
	if postCode == "" {
		fmt.Println("error no postcode supplied")
		os.Exit(1)
	}

	addresses := client.LoadAddressesForPostCode(postCode)

	if len(addresses) == 0 {
		fmt.Println("error no addresses found for postcode", postCode, ". Perhaps it's not an ealing postcode or it was mistyped or mis formatted? It needs to be upper case with a space between the two sections eg W1 2AA")
		os.Exit(1)
	}

	//pick an address

	address := getAddress(addresses)
	fmt.Println("picked: ", address.Address)

	collectionMap := client.LoadCollectionsForAddres(address)

	for collectionType, collections := range collectionMap {
		fmt.Println(collectionType)
		for _, collection := range collections {
			fmt.Println("\t", collection.Date, collection.Frequency)
		}
	}

}


func getAddress(addresses []*ealing.Address) *ealing.Address{
	fmt.Println("Please type the number of your address from the following list. If your address does not appear, please check your postcode is formatted correctly (eg W1 2AA)")
	for i, address := range addresses {
		fmt.Printf("[%d] %s\n", i, address.Address)
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	addressId, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		fmt.Println("error", text, "is not a valid integer")
		os.Exit(1)
	} else {
		if addressId < 0 || int(addressId) >= len(addresses) {
			fmt.Println("error", addressId, "is outside of the range of options available")
			os.Exit(1)
		} else {
			return addresses[addressId]
		}
	}

	return nil
}


func getPostcode() string{
	fmt.Print("Please type your postcode formatted like W1 2AA: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	return text
}
