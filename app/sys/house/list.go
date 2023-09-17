package house

import "server/app/model"

func GetHouseList() ([]model.House, error) {
	var s model.House
	list, err := s.FindList()
	if err != nil {
		return nil, err
	}
	return list, nil
}