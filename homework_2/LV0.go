package main

import "fmt"

func main(){
	type Web struct {
		front string
		after []string
	}
	type RedRock struct {
		Web
		Mobile float64
		VisualDesign complex128
		ProductPlanning bool
		SRE uint
	}

	var I RedRock
	I.after=append(I.after,"Yes")
	I.ProductPlanning=false

	You :=Web{
		front:"Yes",
		}
	fmt.Printf("I：%+v\n",I)
	fmt.Printf("And you ：%+v\n",You)
}