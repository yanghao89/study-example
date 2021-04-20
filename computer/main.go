package  main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	buf := make([]byte, 4096)

	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}

		inStr := strings.TrimSpace(string(buf[0:cnt]))

		inputs := strings.Split(inStr, " ")

		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			fmt.Printf("Unsupported command: %s\n", inputs[0])
		}
	}

	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}
type MPs map[string]int
func main() {
     maps := make(map[float64]MPs,0)
	MPs := map[string]int{
		"uid":1,
		"uid2":1,
	}
	maps[0.99] = MPs
 fmt.Println(maps)
}