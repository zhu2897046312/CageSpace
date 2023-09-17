package cage

import "server/app/model"

func UpdateCage(c model.Cage,OldID model.ID) error {
	err:=c.Update(c,OldID)
	if err!=nil{
		return err
	}
	return nil
}