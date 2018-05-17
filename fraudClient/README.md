# Library for the Fraud Identification API

This is a privately created library to aid in using the Fraud Identification API.

Note that this library is not officially endorsed.
This is simply my library to help me interact with their API, described here:

https://support.kochava.com/analytics-reports-api/fraud-api

## Import

Here is the import code:

```golang
import "github.com/tenglun/kfapi"
```

# Usage

Effort has been made to make this library simple and readable.

First, initiaze the client using your API key, and account id. The system will automatically detect if you are using the library
as a marketer or network (which changes the calls).

```golang
client, err := kfapi.GetAccount("my_api_key", "my_accound_id")
```

### Client Options
Client Options can be passed in as functional options when the client is created. Stock functions are:
```golang
// SetDebugTrue turns debug logic on; client returns spoof data instead of real data for testing
SetDebugTrue()

// SetHttpClient allows a custom http client to be utilized
SetHTTPClient(httpClient *http.Client)

// SetBaseURL allows a custom domain to be hit instead of the proper domain
SetBaseURL(baseURL string)
```

## Retrieve Data

Once the account has been initialized, data can be returned using the "List" endpoint, or the "GatherDataFrom" endpoint.

```golang
response, err := client.List.Apps("fraud type",start time, end time)
```

Example for List Endpoint:
```golang
response, err := client.List.Apps(kfapi.AnonymousInstall, time.Unix(1510000000,0), time.Now())
```

Example for Gather Endpoint:
```golang
response, err := client.GatherDataFrom.Networks(kfapi.AdStacking, time.Unix(1510000000,0), time.Now())
```

### Filters

Filters can be added to modify the data returned. Filters can be passed using the Filter() function
as follows:
```golang
response, err := client.Data.Networks(kfapi.Adstacking, time.Unix(1510000000,0), time.Now(),Filter(NetworkId,In,"1344","2522","133","6352"),Filter(AppId,In,"5667"))
```

## Response

All of the responses will be in an exported struct format called KFResponse. This format is being used to both unify the response
structure from the various calls, as well as to allow an easier "turnaround", sending data back to the API to add or remove from the blacklist.

# Constants

When gathering information,
it must be accessed for one particular fraud type at a time. For convenience, constants have been added and exported from the library.

The constants are directly accessible from the library:
```golang
kfapi.AdStacking
kfapi.AnonymousInstall
kfapi.DeviceHighClickVolume
...
etc.
```

# Improvements Coming Soon

Currently, functionality is being worked on that will allow the return to be ported directly back to the API to add suspicious actors to the
blacklist, after some threshold has been applied.

Stay tuned for more developments.
