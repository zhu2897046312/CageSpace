package model

import(
	"fmt"
	"server/app/core"
)

type Abandon struct {  
	Time    string 
	ID ID
}

var abandon_table string = "abandon"

func (a *Abandon) Find(houseID int,cageID int)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",abandon_table ,houseID ,cageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.Time,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Abandon)Insert(abandon Abandon)error{
	db := core.GetSQLDB()

	err := a.Find(a.ID.HouseID, a.ID.CageID)
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(house_id,cage_id,time)", abandon_table)
	sql += fmt.Sprintf(" values(%d,%d,'%s')",abandon.ID.HouseID,abandon.ID.CageID,abandon.Time)
	_, errExec := db.Exec(sql)
	defer db.Close()
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Abandon) Update(c Abandon,oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(c.ID.HouseID,c.ID.CageID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", abandon_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,time='%s'",c.ID.HouseID,c.ID.CageID,c.Time)
	_, err := db.Exec(sql)
	defer db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (a *Abandon) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id.HouseID,id.CageID)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", abandon_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	defer db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (a *Abandon)FindList()([]Abandon,error){
	db := core.GetSQLDB()

	var s []Abandon
	sql := fmt.Sprintf("select * from %s ",abandon_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Abandon{}
		err=rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.Time,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}
