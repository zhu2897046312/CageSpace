package model

import (
	"fmt"
	"server/app/core"
)

type Count struct {
	Id        int    //自动分配
	SiteID    int
	EggSum    int   
	BadeggSum int
	CubSum    int    
	DeathSum  int   
	Time      string
}

var count_table string = "count"

func (a *Count) Find(siteID int)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where site_id = %d ",count_table ,siteID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.Id,
		&a.SiteID,
		&a.EggSum,
		&a.BadeggSum,
		&a.CubSum,
		&a.DeathSum,
		&a.Time,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Count)Insert(c Count)error{
	db := core.GetSQLDB()

	err := a.Find(c.SiteID)
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(site_id,egg_sum,badegg_sum,cub_sum,death_sum,time) ", count_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,%d,'%s')",
		c.SiteID,
		c.EggSum,
		c.BadeggSum,
		c.CubSum,
		c.DeathSum,
		c.Time,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Count) Update(c Count,siteID int) error {
	db := core.GetSQLDB()
	err1 := a.Find(siteID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", count_table)
	sql += fmt.Sprintf("site_id=%d,egg_sum=%d,badegg_sum=%d,cub_sum=%d,death_sum=%d,time='%s'",
		c.SiteID,
		c.EggSum,
		c.BadeggSum,
		c.CubSum,
		c.DeathSum,
		c.Time,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *Count) Delete(siteID int) error {
	db := core.GetSQLDB()

	err1 := a.Find(siteID)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where site_id=?", count_table)
	_, err := db.Exec(sql,siteID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Count)FindList()([]Count,error){
	db := core.GetSQLDB()

	var s []Count
	sql := fmt.Sprintf("select * from %s ",count_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Count{}
		err=rows.Scan(
			&temp.Id,
			&temp.SiteID,
			&temp.EggSum,
			&temp.BadeggSum,
			&temp.CubSum,
			&temp.DeathSum,
			&temp.Time,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}