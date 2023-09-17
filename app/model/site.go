package model

import (
	"fmt"
	"server/app/core"
)


type Site struct {
	Id   int		// 自分配
	Name string
}
var site_table string = "site"


// 名字查找
func (site *Site) Find(name string) error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where name=?", site_table)

	r := db.QueryRow(sql,name)
	
	err := r.Scan(&site.Id, &site.Name)

	if err != nil {
		return err
	}
	return nil
}

//Id查找
func (site *Site) FindToID(id int) error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select * from %s where id = ?",site_table)
	r:=db.QueryRow(sql,id)
	
	err:=r.Scan(&site.Id,&site.Name)
	if err != nil {
		return err
	}
	return nil
}

func (site *Site) Insert(s Site) error {
	db := core.GetSQLDB()
	err1:=site.Find(s.Name)
	if err1==nil{
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
	}
	
	sql := fmt.Sprintf("insert into %s(id,name) values(?,?)", site_table)
	_, err := db.Exec(sql,s.Id,s.Name)
	if err != nil {
		return err
	}
	return nil
}

//update to name
func (site *Site) Update(old_name string, new_name string) error {
	db := core.GetSQLDB()

	err1 := site.Find(old_name)
	if err1 != nil {
		return err1
	}
	sql := fmt.Sprintf("update %s set name='%s' where name=?;", site_table, new_name)
	_, err := db.Exec(sql,old_name)

	if err != nil {
		return err
	}
	return nil
}

//delete to name
func (site *Site) Delete(name string) error {
	db := core.GetSQLDB()

	err1 := site.Find(name)
	if err1 != nil {
		return err1
	}
	sql := fmt.Sprintf("delete from %s where name=?;", site_table)
	_, err := db.Exec(sql,name)
	if err != nil {
		return err
	}
	return nil
}


//update to id
func (site *Site) UpdateToId(new_name string, old_id int) error {
	db := core.GetSQLDB()

	err1 := site.FindToID(old_id)
	if err1 != nil {
		return err1
	}
	sql := fmt.Sprintf("update %s set name='%s' where id=%d;", site_table, new_name,old_id)
	_, err := db.Exec(sql)

	if err != nil {
		return err
	}
	return nil
}

//delete to id
func (site *Site) DeleteToId(id int) error {
	db := core.GetSQLDB()

	err1 := site.FindToID(id)
	if err1 != nil {
		return err1
	}
	sql := fmt.Sprintf("delete from %s where id=?;", site_table)
	_, err := db.Exec(sql,id)
	if err != nil {
		return err
	}
	return nil
}

func (site *Site)FindList()([]Site,error){
	db := core.GetSQLDB()

	var s []Site
	sql := fmt.Sprintf("select * from %s ",site_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Site{}
		err=rows.Scan(&temp.Id,&temp.Name)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}