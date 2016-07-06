# GoIpApi

A simple wrapper around the  [ip-api.com](http://ip-api.com) API for IP geolocation information written in Go (golang). Works for both the free API as well as the paid, commercial API if you have purchased a Pro license.

There's plenty of options if you need IP geolocation information, including directly using the MaxMind GeoIP2 database. However, I wasn't looking 
at standing up an entire service for making a few calls here and there. [ip-api.com](http://ip-api.com) has a great free tier service for hobby
projects and is quick to get up and running.

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

Keep in mind that the free account is limited to 150 requests per minute. If you exceed this limit your IP address is blacklisted from making further requests.
To correct this visit [this page](http://ip-api.com/docs/unban).

##License

MIT &copy; Jeremiah Piontek
