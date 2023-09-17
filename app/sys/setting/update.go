package setting

import "server/app/model"

func UpdateSetting(c model.Setting,oldID int) error {
	err:=c.Update(c,oldID)
	if err!=nil{
		return err
	}
	return nil
}