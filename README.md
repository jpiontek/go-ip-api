# Goip

A simple wrapper around the ip-api.com API for IP geolocation information. 

## Usage

Import and create a new client.

```go
import "github.com/jpiontek/goip"

client := goip.NewClient()
```

Request your current public IP info

```go
result := client.GetLocation()
```

Or supply a specific IP address.

```go
result := client.GetLocationForIp("192.168.1.1")
```

Keep in mind that ip-api is limited to 150 requests per minute. If you exceed this limit your IP address is blacklisted from making further requests.
To correct this visit [this page](http://ip-api.com/docs/unban).

##License

MIT &copy; Jeremiah Piontek
