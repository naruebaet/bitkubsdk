# bitkub-sdk

> bitkub go sdk is bitkub public api wrapper.

## Usage

```go
package main

import (
	"github.com/naruebaet/bitkubsdk"
	"log"
)

func main() {
	bksdk := bitkubsdk.NewBitkub("xxx", "xxxx")
	resp, err := bksdk.GetSymbols()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
```

## Websocket connection

```go
package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/naruebaet/bitkubsdk"
	"log"
	"time"
)

const ColorRed = "\033[31m"
const ColorYellow = "\033[33m"

func main() {
	bksdk := bitkubsdk.NewBitkub("xxx", "xxx")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bksdk.WatchTicker(ctx, func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(ColorRed, "read : ", err)
				return
			}

			fmt.Println(ColorYellow, string(msg))
		}
	})
}
```

## public function
- GetServerStatus
- GetServerTime
- GetSymbols
- GetTicker
- GetTrades
- GetBids
- GetAsks
- GetBooks
- GetTradingview
- GetDepth

## websocket function
- WatchTicker
- WatchTrade
