package main

import (
	"fmt"
)

func main(){
	var(
		a,b,result float64
		sign string
	)
	for{
		fmt.Printf("input(input q to quit):\n")
		fmt.Scanf("%f%1s%f",&a,&sign,&b)

		result=SimpleCalculate(a,sign,b)
		fmt.Printf("result:\n")
		fmt.Printf("%g\n\n",result)
	}
}

func SimpleCalculate(a float64,sign string,b float64)(float64) {
	var result float64
	if sign=="+"{
		result= a+b
	}else if sign=="-"{
		result= a-b
	}else if sign=="/"{
		result= a/b
	}else if sign=="*" || sign=="X" || sign == "x"{
		result= a*b
	}
	return result

//方法二：用switch来完成
//	switch{
//	case sign=="+":
//		result= a+b
//	case sign=="-":
//		result= a-b
//	case sign=="/":
//		result= a/b
//	case sign=="*" || sign=="X" || sign == "x":
//		result= a*b
//	}
}
