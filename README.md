# Go language tokyo-rs client SDK

Go Client SDK to play with server https://github.com/unrealhoang/tokyo-rs

## Quick start

Quick connect to tokyo-rs server

```go
package main

import tokyo "github.com/ledongthuc/tokyo_go_sdk"


func main() {
  client := tokyo.NewClient("ws://host/socket", "YBOUNCEMEN", "Top 1 Number one")
  log.Fatal(client.Listen())
}
```

## Examples

[Check here](examples)
