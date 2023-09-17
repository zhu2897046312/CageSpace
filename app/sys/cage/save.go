package cage

import "server/app/model"

func SaveCage(c model.Cage)error{
	err:=c.Insert(c)
	if err!=nil{
		return err
	}
	return nil
}