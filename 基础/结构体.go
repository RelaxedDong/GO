package main

import (
	"fmt"
	"unsafe"
)

type Books struct {
	title    string
	author   string
	subject  string
	book_id  int
	username string
}


func main() {
	// 创建一个新的结构体
	//fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	//// 也可以使用 key => value 格式
	//fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	//// 忽略的字段为 0 或 空
	//fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	//fmt.Printf( "Book 1 title : %s\n", Book1.title)
	//fmt.Printf( "Book 1 author : %s\n", Book1.author)
	//fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	//fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
	fmt.Println(Book2.username)
	fmt.Printf("=======", unsafe.Pointer(&Book2))
	fmt.Printf("author is ==",Book2.author)
	Book3 := printBook(&Book2)
	fmt.Printf("=======", unsafe.Pointer(&Book3))
	fmt.Printf(Book2.author)
}

func printBook( book *Books ) Books{
	fmt.Printf("=======", unsafe.Pointer(&*book))
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	book.author = "donghao"
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
	fmt.Printf( "Book book_id : %d\n", book.username);
	return *book
}