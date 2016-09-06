package verifier

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	googleAndroidPublisherService "google.golang.org/api/androidpublisher/v2"
)

// VerifyPurchase uses the given purchase information and makes an API call against Google Play
// Developer API's Purchases.products:get endpoint in order to verify the purchase. Please refer
// to https://developers.google.com/android-publisher/api-ref/purchases/products/get for more
// information.
func VerifyPurchase(credentialsFilePath, bundleID, productID, purchaseToken string) error {
	configFileContent, err := getFileContent(credentialsFilePath)
	if err != nil {
		return err
	}

	config, err := google.JWTConfigFromJSON(configFileContent, googleAndroidPublisherService.AndroidpublisherScope)
	if err != nil {
		return err
	}

	service, err := googleAndroidPublisherService.New(config.Client(context.Background()))
	if err != nil {
		return err
	}

	rawResponse, err := service.Purchases.Products.Get(bundleID, productID, purchaseToken).Do()
	if err != nil {
		return err
	}

	_, err = rawResponse.MarshalJSON()
	if err != nil {
		return err
	}

	return nil
}
