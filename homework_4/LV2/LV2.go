package main

import (
	"fmt"
	"time"
)

func main(){
	//作业要求的计时器
	go timerClock(2,0,0,"没有困难的工作，只有勇敢的打工人！",true)
	go timerClock(8,0,0,"早安，打工人！",true)
	homeworkDuration,_:=time.ParseDuration("1h")
	go pastTimeClock(homeworkDuration,"芜湖！起飞！",true)

	//附加：自定义计时器
	fmt.Printf("请选择您要设置的闹钟类型\n")
	fmt.Printf("1.延时闹钟   2.定时闹钟\n")
	xuanZhe:=check(1,2)
	if xuanZhe==1{
		fmt.Printf("请输入延时时长(h,m,s)各数据用空格隔开\n")
		timeInformation:=timeCheck()
		go pastTimeClock(timeInformation,"\r您的闹钟响啦！\n",false)
	}else{
		fmt.Printf("请输入定时时间(h,m,s)各数据用空格隔开\n")
		var h,m,s int

		func(){		//录入数据
			var ret int
			for{
				ret, _ = fmt.Scanf("%d%d%d",&h,&m,&s)
				//检查是否成功录入数据
				if ret != 3 {
					fmt.Printf("输入错误，请重新输入\n")
					fmt.Scanf("%s")
				}else{
					break
				}
			}
		}()

		go timerClock(h,m,s,"\n您的闹钟响啦！\n",false)
	}
	for{
		time.Sleep(1*time.Second)
		fmt.Printf("\r现在时间：%d-%d-%d %d:%d:%d",time.Now().Year(),time.Now().Month(),time.Now().Day(),
			time.Now().Hour(),time.Now().Minute(),time.Now().Second())
	}
}

//录入延时时间时间数据
func timeCheck() time.Duration{
	var(
		ret,h,m,s int
	)
	for {
		ret, _ = fmt.Scanf("%d%d%d",&h,&m,&s)
		//检查是否成功录入q全部数据
		if ret != 3 {
			fmt.Printf("输入错误，请重新输入\n")
			fmt.Scanf("%s")
		}else{
			break
		}
	}
	information:=fmt.Sprintf("%dh%dm%ds",h,m,s)
	timeInformation, _ :=time.ParseDuration(information)
	return timeInformation
}

//经过一定时间闹钟
func pastTimeClock(timeInFormation time.Duration,text string,continueOrNot bool){
	for{
		time.Sleep(timeInFormation)
		fmt.Println(text)
		if continueOrNot!=true{
			break
		}
	}
}

//定时闹钟
func timerClock(h,m,s int,text string,continueOrNot bool){
	for{
		timeNow:=time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day(),h,m,s,0,time.Now().Location())
		if timeNow.Unix()<time.Now().Unix(){	//通过时间戳查看闹钟时间是否早于当今时间，若否则将闹钟设置在明天
			timeNow.Add(24*time.Hour)
		}
		timeInFormation:=time.Now().Sub(timeNow)
		time.Sleep(timeInFormation)
		fmt.Println(text)
		if continueOrNot==false{		//是否为一次性闹钟
			break
		}
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