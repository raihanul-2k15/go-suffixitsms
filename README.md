# go-mimsms

A simple, easy to use Go package to interact with [Suffix IT Bulk SMS](https://bulkmsg.suffixit.com) API, a Bangladeshi SMS Gateway

## Features

-   [x] Send SMS (single text to many recipients)
-   [ ] Check Balance - not supported by API
-   [ ] Check Delivery status - not supported by API
-   [x] Error messages returned by package does not leak the API key

## Install

```
go get github.com/raihanul-2k15/go-suffixitsms
```

## Usage

### Send SMS

```go
import "github.com/raihanul-2k15/go-suffixitsms/suffixitsms"

apiKey := "yourapikeyhere"

client := suffixitsms.NewClient(apiKey)

resp, err := client.SendMessage([]string{"01717171717"}, "Hello World, API Testing")
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(resp) // Success
```

### Set timeout for request

```go
client := mimsms.NewClient(...)
client.SetTimeout(30 * time.Second)
```

## Disclaimer

Author is not affiliated with Suffix IT
