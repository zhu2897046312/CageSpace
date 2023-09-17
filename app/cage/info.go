package cage

import (
	"server/app/model"
)

func GetCageInfomation(id model.ID) (*model.Cage,error){
	c := new(model.Cage)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c,nil
}

