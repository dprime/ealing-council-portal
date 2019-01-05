package main

import (
	"os"
	"github.com/dprime/ealing-council-portal/ealing"
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

	postCode := os.Args[1]

	addresses := client.LoadAddressesForPostCode(postCode)

	//pick an address

	address := addresses[48]

	client.LoadCollectionTimesForAddress(address)


}

