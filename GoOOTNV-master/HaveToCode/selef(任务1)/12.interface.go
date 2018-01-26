package main

import "fmt"

/*interface被称为接口，是一种类型，其本质是一组抽象方法的集合。
凡是实现这些抽象方法的对象，都可以被称为“实现了这个接口”。
其存在意义是为了规定对象的一组行为。*/
type empty interface{}

type Connecter interface {
	Connect()
}
type USB interface {
	Name() string
	Connecter
}
type PhoneConnecter struct {
	name string
}
type TVConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func (tv TVConnecter) Connect() {
	fmt.Println("TVConnecter:", tv.name)
}

// func Disconnect(usb USB) { // 这里要求的是USB类型, 而USB是interface类型
// 	fmt.Println("Disconnect:", pc.Name())
// }

// func Disconnect(usb empty) { // 这里要求的是USB类型, 而USB是interface类型
// 	if pc, ok := usb.(PhoneConnecter); ok {
// 		fmt.Println("Disconnect:", pc.Name())
// 		return
// 	}
// 	fmt.Println("Unknown device")
// }

/*一个空interface变量可以存入任何值。
实际的用处是，当不确定传入函数的参数类型时，可以使用interface{}代替。
并且，我们有特定的语法可以判断具体存入interface{}变量的类型。
*/
func DisConnect(usb interface{}) { // 这里要求的是USB类型, 而USB是interface类型
	switch v := usb.(type) { // 注意:这种switch和val.(type)配合的语法是特有的，在switch以外的任何地方都不能使用类似于val.(type)这种形式。
	case PhoneConnecter:
		fmt.Println("DisConnect:", v.Name())
	default:
		fmt.Println("Unknown device.")
	}
}
func main() {
	pc := PhoneConnecter{name: "haha"}
	pc.Connect()
	DisConnect(pc)

	other := 1
	DisConnect(other)

	//tv := TVConnecter{name: "tv"}
	var con Connecter
	con = Connecter(pc) // 纯属复制品
	con.Connect()

	fmt.Println("--------------------")
	pc.name = "pc"
	// 注意观察下面输出
	// 发现name还是"PhoneConnector", 不是上面修改后的"pc", 说明con是pc的复制品, 不是引用
	con.Connect()
	fmt.Println("--------------------")
	// con.Name() // error 转换为Connector类型后就不能非Connector的方法了
	// 注意观察下面语句输出
	// 这里con.(type)还是识别类型为PhoneConnector类型,因为数据类型没变, 上面变得是接口类型
	DisConnect(con)

	var a interface{}
	fmt.Println("a == nil = ", a == nil)
	var p *int = nil
	a = p
	fmt.Println("a == nil = ", a == nil) // a虽然指向空指针, 但依然不是nil
}
