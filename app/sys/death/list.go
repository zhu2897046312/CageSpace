package death

import "server/app/model"

func GetDaethList() ([]model.Death, error) {
	var s model.Death
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}