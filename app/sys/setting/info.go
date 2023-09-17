package setting

import "server/app/model"

func GetSettingInfomation(id int)(*model.Setting,error){
	c := new(model.Setting)
	err := c.Find(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}