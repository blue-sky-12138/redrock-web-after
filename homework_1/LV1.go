package main

import "fmt"

func main(){
	for y:=1.5;y>-1.5;y=y-0.1{
		for x:=-1.5;x<1.5;x=x+0.05{
			if ((x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y)<=0{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}