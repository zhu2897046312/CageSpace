package cublist

import "server/app/model"

func UpdateCub(c model.Cub,OldID model.ID) error {
	err:=c.Update(c,OldID)
	if err!=nil{
		return err
	}
	return nil
}