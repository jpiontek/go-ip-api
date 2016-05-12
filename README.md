# Goip

A simple wrapper around the  [ip-api.com](http://ip-api.com) API for IP geolocation information. 

## Usage

Import and create a new client.

```go
import "github.com/jpiontek/goip"

client := goip.NewClient()
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

## Todo

I current do not have a paid account at ip-api.com, this is mostly just a hobby library for somethign else I was working on. If anyone could fill
me in on how the request is made for a commercial account I'd be happy to add support for it - I'm not sure where the API key is supplied in the request. 

##License

MIT &copy; Jeremiah Piontek
