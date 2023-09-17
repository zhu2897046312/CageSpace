package egg

import "server/app/model"

func SaveAEgg(a model.Egg) error {
	err := a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}