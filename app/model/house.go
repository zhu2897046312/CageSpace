package model

import (
	"fmt"
	"server/app/core"
)

type House struct {
	ID      int //自动分配可不要  通过 HouseID 来查找
	HouseID int
	SiteID  int
}

var house_table string = "house"

func (h *House) Find(siteID int,houseID int) error {
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select (site_id,house_id)from %s where site_id = %d and house_id = %d ", house_table,siteID,houseID)

	r := db.QueryRow(sql)
	err := r.Scan(&h.ID,&h.SiteID,&h.HouseID)
	//err := r.Scan(&h.HouseID, &h.SiteID)
	if err != nil {
		return err
	}
	return nil
}
func (house *House) Update(houseID int, siteID int) error {
	db := core.GetSQLDB()
	var s Site
	err2 := s.FindToID(siteID)
	if err2 != nil {
		//没有找到 要改的siteid 所属行  抛异常
		fmt.Print("house update error:没有siteid 所属行")
	}
	err1 := house.Find(siteID, houseID)
	if err1 == nil {
		//找到了 houseid siteid相同的 行 抛异常
		fmt.Print("house update已有该数据,不允许再次插入")
	}

	sql := fmt.Sprintf("update %s set house_id=?, site_id=?", house_table)
	_, err := db.Exec(sql, houseID, siteID)
	if err != nil {
		return err
	}

	return nil
}

func (house *House) Insert(h House) error {
	db := core.GetSQLDB()

	err := house.Find(h.SiteID ,h.HouseID)
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
	}

	sql := fmt.Sprintf("insert into %s(house_id,site_id) values(?,?)", house_table)
	_, errExec := db.Exec(sql, h.HouseID, h.SiteID)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (house *House) Delete(houseID int, siteID int) error {
	db := core.GetSQLDB()

	err1 := house.Find(siteID, houseID)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where site_id=? and house_id=?", house_table)
	_, err := db.Exec(sql, siteID, houseID)
	if err != nil {
		return err
	}
	return nil
}

func (house *House)FindList()([]House,error){
	db := core.GetSQLDB()

	var s []House
	sql := fmt.Sprintf("select * from %s ",house_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=House{}
		err=rows.Scan(
			&temp.ID,
			&temp.SiteID,
			&temp.HouseID,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}
