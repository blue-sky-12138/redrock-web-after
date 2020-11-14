package main

import (
	"fmt"
	"time"
)

var(
	Turn =make(chan int)
)

//双线程、相同内容
func print(choice int,Turn chan int){
	if choice==1{
		for i:=1;i<=100;i+=2{
			fmt.Printf("Print1's paint:%d\n",i)
			Turn<- choice
		}
	}else if choice==2{
		for i:=2;i<=100;i+=2{
			//测试中发现一个进程会打印两条语句，在此对其做延迟使其慢于进程1
			time.Sleep(time.Microsecond * 1)
			fmt.Printf("Print2's paint:%d\n",i)
			<- Turn
		}
	}
}

func main(){
	defer close(Turn)
	go print(1,Turn)
	go print(2,Turn)
	//保证协程完全进行
	time.Sleep(time.Second * 1)
}
