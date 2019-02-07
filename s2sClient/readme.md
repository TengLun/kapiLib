# s2sClient - INCOMPLETE

This library is for the Kochava server-to-server api, which allows data to be sent to Kochava without using an SDK.

## Import

Here is the import code:

```golang
import "github.com/tenglun/kfapi/s2sclient"
```

# Usage

Effort has been made to make this library simple and readable.

First, initialize the client using your app guid:

```golang
client, err := 	s2sclient.CreateClient("my_app_guid")
```

## Sending Installs

Installs and events can be send either through using the "Set" functions, or by
constructing a S2SRequest object, and passing it as the first argument:

Using Set Functions:
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

Using S2SRequest Object:
```golang
install := S2SRequest{
  Data: Data{
    DeviceUa: "Generic Android 6.0",
    DeviceIds: map[string]string{
      "adid": "00000000-0000-0000-0000-000000000000",
    },
    OriginationIP: "127.0.0.1",
    Currency:      "USD",
  },
}

err := client.SendInstall(install)
```

# Improvements Coming Soon

More Set Functions will be added to make sending installs and events easier.

Stay tuned for more developments.
