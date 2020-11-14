package main

import (
	"fmt"
	"os"
)

func main(){
	fileName,_:=os.Create("proverb.txt")
	var p ="Don't communicate by sharing memory share memory by communicating."
	fileName.Write([]byte(p))
	fileContent :=make([]byte,len(p))
	file,_:=os.Open("../proverb.txt")
	defer file.Close()
	file.Read(fileContent)
	fmt.Printf("%s\n",fileContent)
}
