package main

import (
	"fmt"
)

//声明user类型
type user struct {
	username  string
	useremail string
}

//notify 使用值接受者实现一个方法
func (u user) notify() {
	fmt.Printf("sed email to %s <%s> \n", u.username, u.useremail)
}

//使用指针实现一个方法
func (u *user) changeemail(email string) {
	u.useremail = email
}

func main() {
	mytest := user{"zhangsan", "zhangsan@abcd.com"}
	mytest.notify()

	mytest1 := &user{"lisi", "lisi@abcd.com"}
	(*mytest1).notify()
	mytest1.changeemail("wangwu@abcd.com")
	mytest1.notify()

}
