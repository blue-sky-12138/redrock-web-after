package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//base64加密解密
//在这里加密不了，目前仍无法解决
func main(){
	//打开文件
	contect := make([]byte, 0)
	nameKeyMap :=make(map[int]string)
	f,_:=os.OpenFile("base64-users.data" ,os.O_RDWR, os.ModePerm)
	defer f.Close()
	contect, err := ioutil.ReadFile("base64-users.data")
	if err != nil{
		fmt.Printf("Error!\n")
		return
	}

	//读取文件中的用户名和密码并存入map
	nameKey:=bytes.Split(contect,[]byte("\n"))
	for _,x:=range nameKey{
		if len(x)==0{
			break
		}
		if x[len(x)-1] == 13 {
			x=x[:len(x)-1]					//去除密码中的\n
		}
		nameKeySplit:=bytes.Split(x,[]byte(" "))	//分开用户名和密码

		//解密
		names,_:=base64.StdEncoding.DecodeString(string(nameKeySplit[0]))
		namesNumber,_:=strconv.Atoi(string(names))		//获取账号
		keys,_:=base64.StdEncoding.DecodeString(string(nameKeySplit[1])) 	//获取密码

		//存入map
		nameKeyMap[namesNumber]=string(keys)
	}
	fmt.Printf("*******************************\n")
	fmt.Printf("请选择您要使用的业务\n")
	fmt.Printf("1.登录  2.注册  3.退出系统\n")
	fmt.Printf("*******************************\n")
	xuanZhe:=check(1,3)
	if xuanZhe==3{
		fmt.Printf("欢迎再次使用\n")
		return
	}else if xuanZhe==1{
		logIn(nameKeyMap)
	}else if xuanZhe==2{
		newName,newKey:=register(nameKeyMap)

		//加密
		newNameEncoded := base64.StdEncoding.EncodeToString([]byte(newName))
		newKeyEncoded := base64.StdEncoding.EncodeToString([]byte(newKey))
		newInformation:=fmt.Sprintf("%s %s\n",base64.StdEncoding.EncodeToString([]byte(newKeyEncoded)),
			base64.StdEncoding.EncodeToString([]byte(newNameEncoded)))

		n,_:=f.Seek(0,2)
		f.WriteAt([]byte(newInformation),n)
	}
}

//登录
func logIn(nameKeyMap map[int]string)bool{
	//搜索用户名
	var name int
	for {
		fmt.Printf("*******************************\n")
		fmt.Printf("请输入你的用户名：\n")
		fmt.Scanf("%d",&name)
		_,pd :=nameKeyMap[name]
		if pd{
			break
		}else{
			fmt.Printf("该用户名不存在\n")
		}
	}

	//验证密码
	var key string
	for{
		fmt.Printf("*******************************\n")
		fmt.Printf("请输入你的密码：\n")
		fmt.Scanf("%s",&key)
		yz,_ :=nameKeyMap[name]
		if yz==key{
			fmt.Printf("欢迎回来%d\n",name)
			return true					//用于退出系统的指示变量
		}else{
			fmt.Printf("密码错误，请验证您的密码是否正确\n")
		}
	}
}

//注册
func register(nameKeyMap map[int]string)(string,string){
	//用户名输入
	var name string
	for{
		fmt.Printf("*******************************\n")
		fmt.Printf("请输入你的手机号：\n")
		fmt.Scanf("%s",&name)
		x,err:=strconv.Atoi(name)
		if err!=nil{
			fmt.Printf("手机号错误，请重新输入\n")
		}else{
			_,pd :=nameKeyMap[x]
			if x<=10000000000 || x>=100000000000{
				//简易判断手机号位数
				fmt.Printf("手机号错误，请重新输入\n")
			}else if pd==true{
				fmt.Printf("该用户名已存在\n")
			}else{
				break
			}
		}
	}

	//密码输入
	var (
		key string
		keyAgain string
	)
	for{
		fmt.Printf("*******************************\n")
		fmt.Printf("请输入你的密码：\n")
		fmt.Scanf("%s",&key)
		if strings.Contains(key,"/"){
			//简易判断密码非法符号
			fmt.Printf("密码有非法符号，请重新输入\n")
		}else if len(key)<=6{
			fmt.Printf("密码过短，请重新输入\n")
		}else{
			break
		}
	}
	for{
		fmt.Printf("*******************************\n")
		fmt.Printf("请再次输入你的密码：\n")
		fmt.Scanf("%s",&keyAgain)
		if keyAgain!=key{
			fmt.Printf("密码不一致\n")
		}else{
			break
		}
	}
	fmt.Printf("注册成功\n")
	return name,key
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