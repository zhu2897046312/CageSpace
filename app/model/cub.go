package model

import(
	"fmt"
	"server/app/core"
)

type Cub struct {
	ID        ID //ID.id 没有自动分配 需任意随机数 但不能重复
	CubNumber  int   
	Time       string 
	UpdateTime string 
	Statu      int    
}

var cub_table string = "cub"

func (a *Cub) Find(id ID)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",cub_table ,id.HouseID ,id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.CubNumber,
		&a.Time,
		&a.UpdateTime,
		&a.Statu,
	)
	if err != nil {
		return err
	}
	return nil
}


func (a *Cub)Insert(c Cub)error{
	db := core.GetSQLDB()

	err := a.Find(ID{HouseID: c.ID.HouseID,CageID: c.ID.CageID})
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(id,house_id,cage_id,cub_number,time,update_time,statu)", cub_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,'%s','%s',%d)",
		c.ID.Id,
		c.ID.HouseID,
		c.ID.CageID,
		c.CubNumber,
		c.Time,
		c.UpdateTime,
		c.Statu,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Cub) Update(c Cub,oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", cub_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,time='%s',update_time='%s',statu=%d,cub_number=%d",
		c.ID.HouseID,
		c.ID.CageID,
		c.Time,
		c.UpdateTime,
		c.Statu,
		c.CubNumber,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *Cub) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", cub_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Cub)FindList()([]Cub,error){
	db := core.GetSQLDB()

	var s []Cub
	sql := fmt.Sprintf("select * from %s ",cub_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Cub{}
		err=rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.CubNumber,
			&temp.Time,
			&temp.UpdateTime,
			&temp.Statu,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}