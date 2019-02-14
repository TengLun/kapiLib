# Library for the Kochava Reporting API

[![GoDoc](https://godoc.org/github.com/TengLun/kapiLib/reportsclient?status.svg)](https://godoc.org/github.com/TengLun/kapiLib/reportsclient) [![Go Report Card](https://goreportcard.com/badge/github.com/tenglun/kapilib)](https://goreportcard.com/report/github.com/tenglun/kapilib)

This is a privately created library to aid in using the Fraud Identification API, described here:

https://support.kochava.com/analytics-reports-api/api-v1-4-requesting-and-scheduling-reports/

## Import

Here is the import code:

```golang
import "github.com/tenglun/kfapi/reportsclient"
```

# Usage

Effort has been made to make this library simple and readable.

First, initialize the client using your API key, and account id. The system will automatically detect if you are using the library
as a marketer or network (which changes the calls).

```golang
client, err := reportsclient.NewClient("my_app_guid", "my_api_key")
```

### Client Options
Client Options can be passed in as functional options when the client is created. Stock functions are:
```golang
// SetClientHTTPClient allows a custom http client to be used by the client struct
SetClientHTTPClient(customClient *http.Client)

// SetEmailList configures which emails a client will send reports to
SetEmailList(emails []string)

// SetDeliverToS3Bucket configures report delivery to a custom s3 bucket
SetDeliverToS3Bucket(region, bucketName, accessKey, secretKey string)

// SetTimezone sets the report delivery timezone
SetTimezone(tz string)
```

## Retrieve Data


# Improvements Coming Soon
