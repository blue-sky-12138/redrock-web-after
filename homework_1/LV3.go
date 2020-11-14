package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	slice:=make([]int,0)
	fmt.Printf("请选择数组类型\n")
	fmt.Printf("1.自定义数组    2.随机数组\n")
	xuanZhe:=check(1,2)
	if xuanZhe==2{   //随机数组
		RandomSlice(&slice)
		fmt.Printf("已为您随机生成一个随机数组%v\n",slice)

	}else{			//自定义数组
		fmt.Printf("请输入一系列的整数(用空格来分隔数据，输入一个非数字来结束输入)：\n")
		var shuJu  int

		//实现重复录入并自动退出
		ret,_ :=fmt.Scanf("%d",&shuJu)
		for ret==1{
			slice=append(slice,shuJu)
			ret,_ =fmt.Scanf("%d",&shuJu)
		}
		fmt.Printf("输入结束\n")
		fmt.Printf("您的输入的数组如下%v\n",slice)
	}
	reorder(&slice,len(slice))
	result:=minFabs(&slice)
	fmt.Printf("数组中绝对值最小的数为%d",result)
}


//用于生成随机数组，已有函数只能生成非负数，故人为更改奇数项为负数
func RandomSlice(slice *[]int){
	rand.Seed(time.Now().Unix())
	//用系统时间来做种子，做到另一种意义上的随机
	for i:=0;i<10;i++{
		if i%2==0{
			*slice=append(*slice,rand.Intn(1000))
		}else{
			*slice=append(*slice,-rand.Intn(1000))
		}
	}
	return
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

//按大小重排列数组
func reorder(slice *[]int,len int)(result int){
	right:=0
	left:=0
	for{
		right=left+1
		for {  //交换元素使其重排列
			if (*slice)[right]<(*slice)[left]{
				(*slice)[right],(*slice)[left]=(*slice)[left],(*slice)[right]
			}
			if right==(len-1){
				break
			}
			right++    //继续对比之后的元素
		}
		if left==(len-2){
			break      //判断是否到达倒数第二个元素
		}
		left++
	}
	if (*slice)[right]<(*slice)[left]{
		(*slice)[right],(*slice)[left]=(*slice)[left],(*slice)[right]
	}
	return result
}




//计算数组中绝对值最小的数
func minFabs(x *[]int) int {
	var (
		right =len(*x) - 1
		left, middle int
	)
	//先判断二分后是否只剩两个元素，根据中间值元素正负号选定缩小范围
	for (right-left)!=1{
		middle=(right+left)/2
		if (*x)[middle]==0{
			right=middle
			break
		}else if (*x)[middle]<0{
			left=middle
		}else {
			right=middle
		}
	}
	if -(*x)[left]>(*x)[right]{
		return (*x)[right]
	}else {
		return -(*x)[left]
	}
}