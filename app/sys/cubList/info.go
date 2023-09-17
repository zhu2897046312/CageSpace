package cublist

import "server/app/model"

func GetCubInfomation(id model.ID) (*model.Cub, error) {
	c := new(model.Cub)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}