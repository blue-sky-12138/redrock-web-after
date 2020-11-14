package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Printf("您是使用一个随机整数组还是使用自定义整数组？\n")
	fmt.Printf("1.随机数组      2.自定义数组\n")

	//录入用户的选择
	var (
		xz int
		k int
	)
	for x:=0;x==0;{
		k,_=fmt.Scanf("%d",&xz)
		if k==0{
			xz=0
			fmt.Printf("您的选择不在可选项中\n")
			fmt.Scanf("%s")
		}
		if xz==1||xz==2{
			x=1
		}else if xz!=0{
			fmt.Printf("您的选择不在可选项中\n")
		}
	}


	if xz==1{    //随机数组：使用系统时间＋伪随机rand产生数组
		rand.Seed(time.Now().Unix())
		var random []int
		for i:=0;i<10;i++{
			random=append(random,rand.Intn(1000))
		}
		fmt.Printf("已为您随机生成一个随机数组%v\n",random)

		//打印结果
		fmt.Printf("该数组由大到小排列为%v\n",body(random))
		return
	}else {
		//用户自定义数组
		fmt.Printf("请输入一系列的整数(输入非数字来结束输入)：\n")
		var (
			shuJu  int
			zhiZhi = make([]int, 0)
		)

		//实现重复录入并自动退出
		ret,_ :=fmt.Scanf("%d",&shuJu)
		for ret==1{
			zhiZhi=append(zhiZhi,shuJu)
			ret,_ =fmt.Scanf("%d",&shuJu)
		}
		fmt.Printf("输入结束\n")
		fmt.Printf("您的输入的数组如下%v\n",zhiZhi)

		//打印结果
		fmt.Printf("该数组由大到小排列为%v\n",body(zhiZhi))
		return
	}
}
//输出按大小排列的数组
func body(x []int) []int{
	var(
		k=1
		v=1
		max=x[0]
		biaoJi=1
		ret=make([]int,0)
	)
	for len(x)!=1{   //判断不断切分后的数组是否只有一个元素
		for k,v=range x{   //获取数组中最大数的下标和数值
			if v>=max{
				max=v
				biaoJi=k
			}
		}
		//一个一个元素地重排列数组
		ret=append(ret,max)
		//除去数组中的最大数
		x=append(x[:biaoJi],x[biaoJi+1:]...)
		//重置最大数
		max=x[0]
	}
	ret=append(ret,x[:]...)
	return ret
}