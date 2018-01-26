// http://blog.csdn.net/tyq101010/article/details/52156960
package main
import (
	"net/rpc"
	"log"
	"net/http"
	"net"
)
type Echo int

func (t *Echo) Hi(args string, reply *string) error {
    *reply = "This is a message return from 1234 port:" + args
    return nil
}

func main() {
    rpc.Register(new(Echo))
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}