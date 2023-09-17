package count

import "server/app/model"

func UpdateCount(c model.Count,oldSiteID int) error {
	err:=c.Update(c,oldSiteID)
	if err!=nil{
		return err
	}
	return nil
}