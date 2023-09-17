package death

import "server/app/model"

func UpdateDeath(c model.Death,OldID model.ID) error {
	err:=c.Update(c,OldID)
	if err!=nil{
		return err
	}
	return nil
}