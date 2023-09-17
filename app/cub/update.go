package cub

import "server/app/model"

func UpdateCub(c model.Cub,oldId model.ID) error {
	err:=c.Update(c,oldId)
	if err!=nil{
		return err
	}
	return nil
}