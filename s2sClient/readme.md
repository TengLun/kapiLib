# s2sClient - INCOMPLETE

This library is for the Kochava server-to-server api, which allows data to be sent to Kochava without using an SDK.

## Import

Here is the import code:

```golang
import "github.com/tenglun/kfapi/s2sclient"
```

# Usage

Effort has been made to make this library simple and readable.

First, initiaze the client using your app guid:

```golang
client, err := 	s2sclient.CreateClient("my_app_guid")
```

## Sending Installs

Installs are sent as following:

```golang
err := client.SendInstall(nil,
                          SetDeviceUA("Generic Android 6.0"),
                          SetDeviceID(map[string]string{
                            "idfa":"12345-6789-0123",
                            "custom":"my_custom_id",
                            }),
                          SetIPAddress("127.0.0.2"),
                          SetDeviceCurrency("USD"))
```
