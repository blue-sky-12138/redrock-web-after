package main

import "fmt"

type Personal struct { //一键四连
	Favourite int //点赞
	Coin      int //投币
	Collect   int //收藏
}
type PersonalCotrol interface{		//这里并没有用到接口
	Fav()
	Coi()
	Col()
}

func main(){
	var(
		RedRock Personal
		XuanZhe int
		)
	RedRock.Favourite=10
	RedRock.Coin=20
	RedRock.Collect=10
	fmt.Printf("让我们看看数据\n")
	fmt.Printf("%+#v\n",RedRock)
	fmt.Printf("你现在点开了你们的视频！\n")

	fmt.Printf("是否点赞？\n")
	fmt.Printf("1.是    2.否\n\n")
	XuanZhe=check(1,2)
	(&RedRock).Fav(XuanZhe)

	fmt.Printf("投几个硬币？\n")
	fmt.Printf("1.0    2.1    3.2\n\n")
	XuanZhe=check(0,2)
	if XuanZhe==1{
		fmt.Printf("成就：白嫖大师\n\n")
	}
	(&RedRock).Coi(XuanZhe)

	fmt.Printf("是否收藏？\n")
	fmt.Printf("1.是    2.否\n\n")
	XuanZhe=check(1,2)
	(&RedRock).Col(XuanZhe)

	fmt.Printf("让我们看看现在的数据\n")
	fmt.Printf("%+#v\n",RedRock)
	fmt.Printf("干得漂亮！\n")
}

//点赞
func (A *Personal)Fav(Check int){
	if Check==1{
		A.Favourite+=1
	}
}

//投币
func (A *Personal)Coi(number int){
	A.Coin+=number-1
}

//收藏
func (A *Personal)Col(Check int){
	if Check==1{
		A.Collect+=1
	}
}


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