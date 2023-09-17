package site

import "server/app/model"

func DeleteSite(id int) error {
	d := model.Site{}
	err:=d.DeleteToId(id)
	if err!=nil{
		return err
	}

	return nil
}