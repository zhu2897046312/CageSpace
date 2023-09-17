package cub

import "server/app/model"

func SaveAbandon(a model.Cub) error {
	err := a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}