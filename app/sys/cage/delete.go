package cage

import "server/app/model"

func DeteleCage(id model.ID)error{
	c := model.Cage{}
	err:=c.Delete(id)
	if err!=nil{
		return err
	}
	return nil
}