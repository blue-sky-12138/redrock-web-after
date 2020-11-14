package main

import (
	"fmt"
)

type Author struct {		//作者一栏
	Name string             //名字
	VIP bool                //大会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}

type Time struct {			//视频发布时间
	year int
	month int
	day int
	hour int
	minute int
	second int
}


type Topic struct{			//视频抬头
	VideoTpye	string		//视频类型
	Title string			//标题
	Play int64				//播放量
	Barrage int8			//弹幕量
	Time					//时间
	Others	string			//其他荣誉
}

type Body struct {			//视频主体
	Screen	string			//画面
	Setting string
	Schedule int			//进度
	BarrageInput string		//弹幕输入
}

type Personal struct {		//一键四连
	Favourite int			//点赞
	Coin int				//投币
	Collect int				//收藏
	Share int				//分享
	Complaint byte			//投诉
	Setting byte			//设置
}

type Comment struct {		//评论区
	Number int				//评论数
	Text string				//文本
	Experssion string		//表情
}

type Relative struct {		//相关推荐
	Advertisement string	//广告
	OthersVideo string	//推荐视频链接集
}

type Videoer struct {			//整个页面（总感觉给自己挖了个很大的坑）
	Author
	Topic
	Body
	Personal
	Comment
	Relative
}


func main(){
	var RedRock Videoer

	RedRock.Author.Name="BlueSky"
	RedRock.Author.Focus=1000000
	RedRock.Author.Icon="HAHAHAHAHA"
	RedRock.Author.VIP=true
	RedRock.Author.Signature="Welcome to RedRock"

	RedRock.Topic.VideoTpye="学习作品"
	RedRock.Topic.Title="我们红岩网校真是太强啦"
	RedRock.Topic.Play=0
	RedRock.Topic.Barrage=0
	RedRock.Topic.Others="年度最佳学习作品"
	//RedRock.Topic.Time我不初始化了┭┮﹏┭┮

	RedRock.Relative.Advertisement="https://www.yuque.com/cxyuts/gyq5k1/fslonp#dhFnh"
	RedRock.Relative.OthersVideo= "https://www.yuque.com/cxyuts"

	RedRock.Personal.Favourite=888888
	RedRock.Personal.Coin=0				//白嫖大成功
	RedRock.Personal.Collect=1000000	//收藏大于点赞+硬币
	RedRock.Personal.Share=50			//五十解君愁

	//是在没精力把这些全部初始化，放过我吧
	fmt.Printf("%+#v",RedRock)
}