package site

import "server/app/model"

func UpdateSite(c model.Site, oldId int) error {
	err := c.UpdateToId(c.Name, oldId)
	if err != nil {
		return err
	}
	return nil
}