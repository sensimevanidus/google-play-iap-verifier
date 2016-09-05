# google-play-iap-verification
Server-side IAP verification for Google Play.

## Requirements

* You must download the JSON-formatted credentials file from the Google Developer Console for the application you'll be using.

## Request Fields

| Name | Key | Description |
| ---- | --- | ----------- |
| Bundle ID | bundleID | Bundle ID of the application |
| Product ID | productID | ID of the purchased product |
| Purchase Token | purchaseToken | Token provided by the Play Store API upon purchase completion on the client-side |

## Response Fields

| Name | Key | Values | Description |
| ---- | --- | ------ | ----------- |
| Is successful? | isSuccessful | true or false | Flag indicating whether the Play Store API verifies the purchase or not |
| Error details | errorDetails | (optional) string | Message that gives further information about the error |
