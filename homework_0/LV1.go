package main

import(
	"fmt"
)

func main(){
	var(
		a float64
		b float64
		c string
	)
	fmt.Println("input:")
	fmt.Scanf("%f%1s%f",&a,&c,&b)
	if c=="+"{
		fmt.Printf("%g",a+b)
	}else if c=="-"{
		fmt.Printf("%g",a-b)
	}else if c == "*"{
		fmt.Printf("%g",a*b)
	}else if c=="/"{
		fmt.Printf("%g",a/b)
	}
}
