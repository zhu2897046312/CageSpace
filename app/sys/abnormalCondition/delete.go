package abnormalCondition

import "server/app/model"

func DeleteAbnormalCondition(id model.ID)error{
	c := model.AbnormalCondition{}
	err:=c.Delete(id)
	if err!=nil{
		return err
	}
	return nil
}