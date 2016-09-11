# go-ip-api

A simple wrapper around the  [ip-api.com](http://ip-api.com) API for IP geolocation information written in Go (golang). Works for both the free API as well as the paid, commercial API if you have purchased a Pro license.

## Usage

Import and create a new client.

```go
import "github.com/jpiontek/go-ip-api"

client := goip.NewClient()
```

If you have an api key for a paid account then use

```go
client := goip.NewClientWithApiKey("my-api-key")
```

Request your current public IP info.

```go
result := client.GetLocation()
```

Or supply a specific IP address.

```go
result := client.GetLocationForIp("127.0.0.1")
```

Keep in mind that the free account is limited to 150 requests per minute. If you exceed this limit your IP address is blacklisted from making further requests.
To correct this visit [this page](http://ip-api.com/docs/unban).

##License

MIT &copy; Jeremiah Piontek
