# Goip

A simple wrapper around the  [ip-api.com](http://ip-api.com) API for IP geolocation information. 

## Usage

Import and create a new client.

```go
import "github.com/jpiontek/goip"

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

Keep in mind that ip-api is limited to 150 requests per minute for a free account. If you exceed this limit your IP address is blacklisted from making further requests.
To correct this visit [this page](http://ip-api.com/docs/unban).

##License

MIT &copy; Jeremiah Piontek
