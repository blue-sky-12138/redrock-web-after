package main

import (
	"fmt"
)

func  main()  {
	var Date interface{}
	a:= "RedRock"			//在这里改变输入接受者函数的值
	Date=a
	Receiver(Date)
}

func Receiver(Date interface{}){
	switch Date.(type){
	case bool:
		fmt.Printf("This type of the date is：bool\n\n")
	case int:
		fmt.Printf("This type of the date is：int\n\n")
	case complex128:
		fmt.Printf("This type of the date is：complex128\n\n")
	case float64:
		fmt.Printf("This type of the date is：float64\n\n")
	case string:
		fmt.Printf("This type of the date is：string\n\n")
	case byte:
		fmt.Printf("This type of the date is：byte\n\n")

	}
}