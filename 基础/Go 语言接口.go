//Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
/* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//method_name2 [return_type]
//method_name3 [return_type]
//...
//method_namen [return_type]
//}
//
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
//	/* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
//	/* 方法实现*/
//}


package main

import "fmt"

type Phone interface {
	call()
}

type NoKia struct {
}

type IPhone struct{
}

func (nokiaPhone NoKia) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


func main() {
	var phone Phone
	phone = new(NoKia)
	phone.call()
	phone = new(IPhone)
	phone.call()
}
