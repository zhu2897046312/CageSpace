package abnormalCondition

import "server/app/model"

func UpdateAbnormalCondition(c model.AbnormalCondition,OldID model.ID) error {
	err:=c.Update(c,OldID)
	if err!=nil{
		return err
	}
	return nil
}