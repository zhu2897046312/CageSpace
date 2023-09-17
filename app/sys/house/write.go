package house

import "server/app/model"

func writeHouse(a model.House) error {
	err := a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}