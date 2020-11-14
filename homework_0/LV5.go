package main

import "fmt"

func main(){
	fmt.Printf("请输入一段字符串：\n")
	var ch string
	fmt.Scanf("%s",&ch)
	var(
		zhongXin =0
		changDu =0
		i=0
		j=0
		jiOuPanDuan=1
	)

	//中心对称情况
	for i=1;i<(len(ch)-1);{
		for j=1;ch[i-j]==ch[i+j];{
			if j>changDu{
				changDu=j
				zhongXin=i
			}
			if (i-j)==0 || (i+j)==(len(ch)-1){
				break
			}
			j++
		}
		i++
	}

	//轴对称情况
	for i=0;i<(len(ch)-1);{
		for j=0;ch[i-j]==ch[i+j+1];{
			if j>=changDu{
				changDu=j
				zhongXin=i
				jiOuPanDuan=0
			}
			if (i-j)==0 || (i+j+1)==(len(ch)-1){
				break
			}
			j++
		}
		i++
	}

	if changDu==0{
		fmt.Printf("没有回文子串\n")
		return
	}
	if jiOuPanDuan==1{
		fmt.Printf("最长的回文子串是从第%d个字符开始的%s\n",
			zhongXin-changDu+1,ch[zhongXin-changDu:zhongXin+changDu+1])
	}else{
		fmt.Printf("最长的回文子串是从第%d个字符开始的%s\n",
			zhongXin-changDu+1,ch[zhongXin-changDu:zhongXin+changDu+2])
	}
}

