package count

import "server/app/model"

func GetCountList() ([]model.Count, error) {
	var s model.Count
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}