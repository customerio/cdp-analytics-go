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
