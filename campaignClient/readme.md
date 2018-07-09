# Library for the Campaign Management API

[![GoDoc](https://godoc.org/github.com/TengLun/kapiLib/campaignClient?status.svg)](https://godoc.org/github.com/TengLun/kapiLib/campaignClient)[![Go Report Card](https://goreportcard.com/badge/github.com/TengLun/kapiLib)](https://goreportcard.com/report/github.com/TengLun/kapiLib)


This is a privately created library for the Kochava Campaign Management API.


https://support.kochava.com/analytics-reports-api/premium-publisher-api-campaign-management

## Import

Here is the import code:

```golang
import "github.com/tenglun/kfapi/campaignclient"
```

# Usage

Effort has been made to make this library simple and readable.

First, initiaze the client using your API key and app id.

```golang
a := campaignclient.AccountAccessor{
  AppID:   "myAppId",
  AuthKey: "12345-12345-12345-12345-12345",
}
client, err := 	campaignclient.CreateClient(a)
```

### Client Options
Client Options can be passed in as functional options when the client is created. Stock functions are:
```golang
// SetDebugTrue turns debug logic on; client returns spoof data instead of real data for testing
SetDebugTrue()

// SetHttpClient allows a custom http client to be utilized
SetHTTPClient(httpClient *http.Client)

```

## Using the API

Once the account has been initialized, the client provides access to all the available methods.

Example:

```golang
trackers, err := client.GetTrackers()
```

# Improvements Coming Soon

Stay tuned for more developments.
