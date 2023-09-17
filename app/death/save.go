package death

import "server/app/model"

func SaveDeath(a model.Death) error {
	err := a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}