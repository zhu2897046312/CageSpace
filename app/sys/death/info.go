package death

import "server/app/model"

func GetDeathInfomation(id model.ID) (*model.Death, error) {
	c := new(model.Death)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}