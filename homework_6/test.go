package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main()  {
	fmt.Printf("%#v\n",time.Now().Date())
	fmt.Println(time.Now().Clock())
}
