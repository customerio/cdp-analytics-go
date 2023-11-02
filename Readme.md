Customer.io Data Pipelines analytics client for Go.

## Installation

```
go get github.com/customerio/cdp-analytics-go
```

## Usage

```go
package main

import (
    "os"

    "github.com/customerio/cdp-analytics-go"
)

func main() {
    client := analytics.New(os.Getenv("WRITE_KEY"))

    // Enqueues a track event that will be sent asynchronously.
    client.Enqueue(analytics.Track{
        UserId: "4",
        Event:  "order_complete",
    })

    // Flushes any queued messages and closes the client.
    client.Close()
}
```

## Other Regions

If you're using a [different data center](https://customer.io/docs/accounts-and-workspaces/data-centers/) such as our EU region, you can specify an alternate endpoint:

```go
package main

import (
    "os"

    "github.com/customerio/cdp-analytics-go"
)

func main() {
    client, err := analytics.NewWithConfig(os.Getenv("WRITE_KEY"), analytics.Config{
        Endpoint: "https://cdp-eu.customer.io",
    })
    if err != nil {
        panic(err)
    }

    // ...
}
```

## Documentation

The links below contain more detailed documentation on how to use this library:

* [Documentation](https://customer.io/docs/cdp/sources/connections/servers/go/)
* [Specs](https://customer.io/docs/cdp/sources/source-spec/source-events/)
* [godoc](https://pkg.go.dev/github.com/customerio/cdp-analytics-go)
