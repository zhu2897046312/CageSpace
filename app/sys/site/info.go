package site

import "server/app/model"

func GetCageInfomation(id int) (*model.Site, error) {
	c := new(model.Site)
	err := c.FindToID(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
