package model

import(
	"fmt"
	"server/app/core"
)

type Death struct {
	ID          ID    
	OldNumber   int    
	YoungNumber int   
	Time        string 
}

var death_table string = "death"

func (a *Death) Find(id ID)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",death_table ,id.HouseID ,id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.OldNumber,
		&a.YoungNumber,
		&a.Time,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Death)Insert(c Death)error{
	db := core.GetSQLDB()

	err := a.Find(ID{HouseID: c.ID.HouseID,CageID: c.ID.CageID})
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(house_id,cage_id,old_number,young_number,time)", death_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,'%s')",
		c.ID.HouseID,
		c.ID.CageID,
		c.OldNumber,
		c.YoungNumber,
		c.Time,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Death) Update(c Death,oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", death_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,old_number=%d,young_number=%d,time='%s'",
		c.ID.HouseID,
		c.ID.CageID,
		c.OldNumber,
		c.YoungNumber,
		c.Time,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *Death) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", death_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Death)FindList()([]Death,error){
	db := core.GetSQLDB()

	var s []Death
	sql := fmt.Sprintf("select * from %s ",death_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Death{}
		err=rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.OldNumber,
			&temp.YoungNumber,
			&temp.Time,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}