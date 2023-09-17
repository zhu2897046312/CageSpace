package model

import(
	"fmt"
	"server/app/core"
)

type WhippingEgg struct {
	ID ID  				// 这里的 ID.id 没有自动分配 需自行设置 ，但不能有所重复
	EggNumber    int    
	BadeggNumber int   
	Time         string 
}

var whipping_table string = "whipping_egg"

func (a *WhippingEgg) Find(id ID)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",whipping_table ,id.HouseID ,id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.EggNumber,
		&a.BadeggNumber,
		&a.Time,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *WhippingEgg)Insert(c WhippingEgg)error{
	db := core.GetSQLDB()

	err := a.Find(ID{HouseID: c.ID.HouseID,CageID: c.ID.CageID})
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(id,house_id,cage_id,egg_number,badegg_number,time)", whipping_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,%d,'%s')",
		c.ID.Id,
		c.ID.HouseID,
		c.ID.CageID,
		c.EggNumber,
		c.BadeggNumber,
		c.Time,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *WhippingEgg) Update(c WhippingEgg,oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", whipping_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,egg_number=%d,badegg_number=%d,time='%s'",
		c.ID.HouseID,
		c.ID.CageID,
		c.EggNumber,
		c.BadeggNumber,
		c.Time,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *WhippingEgg) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", whipping_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (a *WhippingEgg)FindList()([]WhippingEgg,error){
	db := core.GetSQLDB()

	var s []WhippingEgg
	sql := fmt.Sprintf("select * from %s ",whipping_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=WhippingEgg{}
		err=rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.EggNumber,
			&temp.BadeggNumber,
			&temp.Time,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}
