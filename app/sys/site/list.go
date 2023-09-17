package site

import "server/app/model"

func GetSiteList() ([]model.Site, error) {
	var s model.Site
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}