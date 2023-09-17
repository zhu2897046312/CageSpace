package abnormalCondition

import "server/app/model"

func GetAbnormalConditionList() ([]model.AbnormalCondition, error) {
	var s model.AbnormalCondition
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}