package count

import "server/app/model"

func DeleteDeath(id int)error{
	c := model.Count{}
	err:=c.Delete(id)
	if err!=nil{
		return err
	}
	return nil
}