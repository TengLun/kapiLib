# Library for the Campaign Management API

![GoDoc](https://godoc.org/github.com/TengLun/kapiLib/campaignClient?status.svg)

This is a privately created library for the Kochava Campaign Management API.

Note that this library is not officially endorsed.
This is simply my library to help me interact with their API, described here:

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
client, err := campaignclient.CreateClient(AccountAccessor{	AppID: "myAppId", AuthKey "12345-abcde"})
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
