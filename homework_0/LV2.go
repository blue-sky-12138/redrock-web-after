package main

import(
	"fmt"
)

func main(){                         //先声明map,此时该map为nil
	var nameKey map[string]string    //再使用make函数创建一个非nil的map，nil map不能赋值
	nameKey =make(map[string]string) //也可以用:=代替，可省略第一步
	nameKey["France"]="Paris"
	nameKey["Italy"]="Rome"
	nameKey["Japan"]="Tokyo"
	var (
		x=1
		name string
		key string
	)
	for x==1{
		fmt.Printf("请输入你的用户名(输入q来退出)：\n")
		fmt.Scanf("%s",&name)
		if name=="q"{
			fmt.Printf("欢迎再次使用\n")
			return
		}
		_,pd :=nameKey[name]
		if pd{
			x=0
		}else{
			fmt.Printf("该用户名不存在\n")
		}
	}
	for i:=3;i>=0;i--{
		fmt.Printf("请输入你的密码(输入q来退出)：\n")
		fmt.Scanf("%s",&key)
		if key=="q"{
			fmt.Printf("欢迎再次使用\n")
			return
		}
		yz,_ :=nameKey[name]
		if yz==key{
			fmt.Printf("欢迎回来%s\n",name)
			return
		}else {
			fmt.Printf("密码错误，你还有%d次机会\n",i)
		}
	}
}
