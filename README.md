# Goip

A simple wrapper around the ip-api.com API for IP geolocation information. 

## Usage

Usage

Create a new client and request your current IP's info.

```go
import "github.com/jpiontek/goip"

client := goip.NewClient()
result := client.GetLocation()
```

Or supply a specific IP address.

```go
import "github.com/jpiontek/goip"

client := goip.NewClient()
result := client.GetLocationForIp("192.168.1.1")
```

##License

MIT &copy; Jeremiah Piontek
