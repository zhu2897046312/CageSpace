package cublist
import "server/app/model"
func GetCubList() ([]model.Cub, error) {
	var s model.Cub
	list,err:=s.FindList()
	if err!=nil{
		return nil,err
	}
	return list,nil
}