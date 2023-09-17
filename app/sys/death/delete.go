package death

import "server/app/model"

func DeleteDeath(id model.ID)error{
	c := model.Death{}
	err:=c.Delete(id)
	if err!=nil{
		return err
	}
	return nil
}