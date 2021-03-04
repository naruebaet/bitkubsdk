package socket

import (
	"flag"
	"fmt"
	gwebsocket "github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = flag.String("addr", "api.bitkub.com", "ws service")

func GetWSSession(path, streamData string) *gwebsocket.Conn {
	p := fmt.Sprintf("%s/%s", path, streamData)
	u := url.URL{Scheme: "wss", Host: *addr, Path: p}
	c, _, err := gwebsocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return nil
	}

	return c

}
