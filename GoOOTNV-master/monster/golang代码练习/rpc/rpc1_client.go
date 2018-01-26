// 客户端
package main

import (
	"log"
	"fmt"
	"net/rpc"
)
func main() {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var args = "hello rpc"
    var reply string
    err = client.Call("Echo.Hi", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("replay\t%s\n", reply)
}


// -------
// replay	This is a message return from 1234 port:hello rpc