package count

import "server/app/model"

func GetCountInfomation(id int) (*model.Count, error) {
	c := new(model.Count)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}