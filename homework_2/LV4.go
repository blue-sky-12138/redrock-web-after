package main

import (
	"fmt"
	"math"
)

type Painter interface {
	Paint()
}
type Heart struct {}
type Sin struct {}
type Circle struct {}

func PainterFocus(v Painter){
	v.Paint()
}


func main(){
	fmt.Printf("请选择你要打印的函数\n")
	fmt.Printf("1.心型	2.sin	3.圆\n")
	Select:=check(1,3)
	if Select==1{
		var heart Heart
		PainterFocus(heart)
	}else if Select==2{
		var sin Sin
		PainterFocus(sin)
	}else if Select==3{
		var circle Circle
		PainterFocus(circle)
	}
}



func (_ Heart)Paint(){
	for y:=1.5;y>-1.0;y=y-0.1{
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
func (_ Sin)Paint(){
	for y:=1.0;y>=0;y=y-0.1{
		for x:=-5.0;x<5;x=x+0.1{
			if math.Sin(x)>=y{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	for y:=0.0;y>=-1;y=y-0.1{
		for x:=-5.0;x<5;x=x+0.1{
			if math.Sin(x)<=y{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func (_ Circle)Paint(){
	for y:=1.5;y>-1.5;y=y-0.1{
		for x:=-1.5;x<1.5;x=x+0.05{
			if (x*x+y*y)<=1.5{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

//录入用户选择，判断是否在整数a与b之间，并自动排除错误
func check(a,b int) int {
	var (
		xuanZhe int
		ret int
	)
	for x:=0;x==0;{
		ret,_=fmt.Scanf("%d",&xuanZhe)
		//检查是否成功录入数据
		if ret==0{
			xuanZhe=0
			fmt.Printf("您的选择不在可选项中\n")
			fmt.Scanf("%s")
		}
		if xuanZhe >= a && xuanZhe <= b{
			x=1
		}else if xuanZhe!=0{
			fmt.Printf("您的选择不在可选项中\n")
		}
	}
	return xuanZhe
}