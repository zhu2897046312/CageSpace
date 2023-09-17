package site

import "server/app/model"

func SaveASite(a model.Site) error {
	err := a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}