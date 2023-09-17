package model

import(
	"fmt"
	"server/app/core"
)

type Egg struct {
	ID      ID   // ID.id 没有自动分配 需任意随机数 但不能重复
	Time    string 
	Statu   int    
}

var egg_table string = "egg"

func (a *Egg) Find(id ID)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",egg_table ,id.HouseID ,id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.Time,
		&a.Statu,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Egg)Insert(c Egg)error{
	db := core.GetSQLDB()

	err := a.Find(ID{HouseID: c.ID.HouseID,CageID: c.ID.CageID})
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(id,house_id,cage_id,statu,time)", egg_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,'%s')",
		c.ID.Id,
		c.ID.HouseID,
		c.ID.CageID,
		c.Statu,
		c.Time,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Egg) Update(c Egg,oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", egg_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,statu=%d,time='%s'",
		c.ID.HouseID,
		c.ID.CageID,
		c.Statu,
		c.Time,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *Egg) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", egg_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Egg)FindList()([]Egg,error){
	db := core.GetSQLDB()

	var s []Egg
	sql := fmt.Sprintf("select * from %s ",egg_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Egg{}
		err=rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.Time,
			&temp.Statu,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}