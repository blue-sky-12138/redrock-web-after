package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			over <- true	//更改2：将同步管道放入协程内
		}()
		//if i == 9 {
		//	over <- true	错误2：同步管道在协程外
		//}
		<-over			//更改1：将同步管道放入循环体内
	}
	//<-over			错误1：同步管道在循环外
	fmt.Println("over!!!")
}