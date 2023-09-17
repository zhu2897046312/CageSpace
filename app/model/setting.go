package model

import(
	"fmt"
	"server/app/core"
)


type Setting struct {
	Id           int // 此字段 为 唯一标识
	LayeggTimes  int 
	EggDays      int 
	LayeggDays   int 
	AbandonTimes int 
	BadeggTimes  int 
}

var setting_table string = "setting"

func (a *Setting) Find(id int)error{
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where id = %d ",setting_table ,id)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.Id,
		&a.LayeggTimes,
		&a.EggDays,
		&a.LayeggDays,
		&a.AbandonTimes,
		&a.BadeggTimes,
	)
	if err != nil {
		return err
	}
	return nil
}


func (a *Setting)Insert(c Setting)error{
	db := core.GetSQLDB()

	err := a.Find(c.Id)
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(id,layegg_times,egg_days,layegg_days,abandon_times,badegg_times) ", setting_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,%d,%d,%d)",
		c.Id,
		c.LayeggTimes,
		c.EggDays,
		c.LayeggDays,
		c.AbandonTimes,
		c.BadeggTimes,
	)
	fmt.Printf("sql: %v\n", sql)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *Setting) Update(c Setting,oldID int) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", setting_table)
	sql += fmt.Sprintf("layegg_times=%d,egg_days=%d,layegg_days=%d,abandon_times=%d,badegg_times=%d",
		c.LayeggTimes,
		c.EggDays,
		c.LayeggDays,
		c.AbandonTimes,
		c.BadeggTimes,
	)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *Setting) Delete(id int) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where id=?", setting_table)
	_, err := db.Exec(sql,id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Setting)FindList()([]Setting,error){
	db := core.GetSQLDB()

	var s []Setting
	sql := fmt.Sprintf("select * from %s ",setting_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Setting{}
		err=rows.Scan(
			&temp.Id,
			&temp.LayeggTimes,
			&temp.EggDays,
			&temp.LayeggDays,
			&temp.AbandonTimes,
			&temp.BadeggTimes,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}
