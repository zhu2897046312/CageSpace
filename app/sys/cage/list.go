package cage

import "server/app/model"

func GetCageList() ([]model.Cage, error) {
	var s model.Cage
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}