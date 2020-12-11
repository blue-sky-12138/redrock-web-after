package database

import "fmt"

type Message struct {				//单条留言信息结构体
	Retract string					//缩进，用于显示时更好地看出是一个楼层
	ID int
	TimeDate string
	Author string
	Information string
	Like int
	authorityId int
}

//添加留言
func InsertMessage(dataTime string,author string,information string,previousFloor int,outsideName string,authorityId int){
	prepare:=fmt.Sprintf("insert message_history (time,user,information,previous_floor,outside_name,authority_id)value('%s','%s','%s',%d,'%s',%d)",
		dataTime,author,information,previousFloor,outsideName,authorityId)
	stmt,err:= DataBase.Prepare(prepare)
	defer stmt.Close()
	CheckError("InsertMessage error",err)
	stmt.Exec()
}

//权限查找
func MessageAuthority(authorityId int,memberId int)bool{
	if authorityId==0{
		return true
	}else{
		prepareQuery:=fmt.Sprintf("select ID from authority_relation where group_id=%d and member_id=%d",authorityId,memberId)
		stmt,errQuery:= DataBase.Query(prepareQuery)
		CheckError("MessageAuthority error",errQuery)
		defer stmt.Close()
		if stmt.Next(){
			return true
		}else {
			return false
		}
	}
}
//查找留言
func FindAllMessages(messages *[]Message,previousFloor int,retract int,memberId int) {
	var temMessage Message
	prepare:=fmt.Sprintf("select id,time,outside_name,information,authority_id from message_history where previous_floor=%d order by id desc",previousFloor)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindMessage error",err)
	defer stmt.Close()
	for stmt.Next(){
		stmt.Scan(&temMessage.ID,&temMessage.TimeDate,&temMessage.Author,&temMessage.Information,&temMessage.authorityId)
		if MessageAuthority(temMessage.authorityId,memberId){
			temMessage.Like=FindMessageFavoriteSum(temMessage.ID)
			for i:=0;i<retract;i++{
				temMessage.Retract+="____"	//缩进添加，用于显示时更好地看出是一个楼层（看着难受，抱歉就忍一下吧）
			}
			*messages=append(*messages,temMessage)
			FindAllMessages(messages,temMessage.ID,retract+1,memberId)
		}
	}
}

//寻找是否有该楼层ID
func FindFloorID(ID int)bool{
	prepare:=fmt.Sprintf("select * from message_history where ID=%d",ID)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindFloorID error",err)
	defer stmt.Close()
	if stmt.Next(){
		return true
	}
	return false
}

//点赞
func MessageFavorite(Floor int, like int, name string) {
	prepare:=fmt.Sprintf("insert message_favorite (floor,name,like)value(%d,'%s',%d)",Floor,name,like)
	stmt,err:= DataBase.Prepare(prepare)
	CheckError("Insert MessageFavorite error",err)
	defer stmt.Close()
	stmt.Exec()
}

//获取点赞量
func FindMessageFavoriteSum(Floor int) int {
	var like int
	prepare:=fmt.Sprintf("select sum(if_like) as sumlike from message_favorite where floor =%d",Floor)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindMessageFavoriteNumber error",err)
	defer stmt.Close()
	if stmt.Next(){
		stmt.Scan(&like)
	}
	return like
}

//判断是否有资格删除留言
func DeleteQualify(name string, floor int) bool {
	prepareQuery:=fmt.Sprintf("select ID from message_history where ID=%d and user='%s'",floor,name)
	stmt,errQuery:= DataBase.Query(prepareQuery)
	CheckError("DeleteQualify error",errQuery)
	defer stmt.Close()
	if stmt.Next(){
		return true
	}else {
		return false
	}
}

//删除留言
func DeleteMessage(Floor int){
	//寻找所有回复
	prepareQuery:=fmt.Sprintf("select ID from message_history where ID=%d",Floor)
	stmtQuery,errQuery:= DataBase.Query(prepareQuery)
	CheckError("DeleteMessage-FindFloor error",errQuery)
	defer stmtQuery.Close()
	if stmtQuery.Next(){
		//执行删除操作
		prepareDelete:=fmt.Sprintf("delete from message_history where id=%d",Floor)
		stmtDelete,errDelete:= DataBase.Prepare(prepareDelete)
		CheckError("DeleteMessage-Delete error",errDelete)
		defer stmtDelete.Close()
		stmtDelete.Exec()
	}
}