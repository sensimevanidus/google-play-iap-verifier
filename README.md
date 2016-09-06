# google-play-iap-verification
Server-side IAP verification for Google Play.

## Why this repo?

Recently, I had to implement an IAP (In-app Purchase) verification solution written in **go** for Google Play store. I searched for ready-to-use libraries; but couldn't find any. There were some, but they were implemented in other programming languages. That's why I created this repo.

This is a simple **go** package and does not provide any HTTP API. You'll need to wrap this with an HTTP API on your own.

## Requirements

* You must download the JSON-formatted credentials file from the Google Developer Console for the application you'll be using.

## Usage examples

```go
package main

import (
	"fmt"

	gplayVerifier "github.com/sensimevanidus/google-play-iap-verifier/verifier"
)

func main() {
	if err := gplayVerifier.VerifyPurchase("<credentialsFilePath>", "<bundleID>", "<productID>", "<purchaseToken>"); err != nil {
		fmt.Printf("Purchase is not valid. Error details: %v\n", err.Error())
	}

	fmt.Println("Purchase is valid.")
}
```
