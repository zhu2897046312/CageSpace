package house

import "server/app/model"

func readHouse(siteID int, houseID int) (*model.House,error) {
	c := new(model.House)
	err := c.Find(siteID,houseID)
	if err != nil {
		return nil, err
	}
	return c, nil
}