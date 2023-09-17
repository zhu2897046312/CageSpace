package abandon

import(
	"server/app/model"
)

func SaveAbandon(a model.Abandon)error{
	err:=a.Insert(a)
	if err != nil {
		return err
	}
	return nil
}