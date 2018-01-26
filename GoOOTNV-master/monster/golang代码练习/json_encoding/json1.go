package main 
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
type Message struct {
    Name string
    Body string
    Time int64
}

func main(){
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err !=nil{

	}
	//b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	fmt.Println(string(b))


	c := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var n Message
	err2 := json.Unmarshal(c, &n)
	if err2!=nil{
		log.Fatal(err2)
	}
	fmt.Printf("%#v\n",n)


	cc := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err3 := json.Unmarshal(cc, &f)
	if err3!=nil{
		log.Fatal(err3)
	}
	fmt.Printf("%#v",f)
//	通过类型断言来访问
//	m := f.(map[string]interface{})
//	for k, v := range m {
//		switch vv := v.(type) {
//		case string:
//			fmt.Println(k, "is string", vv)
//		case int:
//			fmt.Println(k, "is int", vv)
//		case []interface{}:
//			fmt.Println(k, "is an array:")
//			for i, u := range vv {
//				fmt.Println(i, u)
//			}
//		default:
//			fmt.Println(k, "is of a type I don't know how to handle")
//		}
//	}


	//流式编解码
	//Encode writes the JSON encoding of v to the stream,
	// followed by a newline character.


	// Decode reads the next JSON-encoded value from its
	// input and stores it in the value pointed to by v.

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
