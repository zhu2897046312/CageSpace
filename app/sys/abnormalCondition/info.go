package abnormalCondition

import "server/app/model"

func GetAbnormalConditionInfomation(id model.ID) (*model.AbnormalCondition, error) {
	c := new(model.AbnormalCondition)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}