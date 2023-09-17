package main

import (
	"fmt"
	"server/app/model"
)


func main() {

	h := model.House{}
	hlist,err:=h.FindList()
	if err!=nil{
		panic(err)
	}
	for i:=0;i<len(hlist);i++{
		fmt.Printf("hlist: %v\n", hlist)
	}
}
