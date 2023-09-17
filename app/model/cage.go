package model

import (
	"fmt"
	"server/app/core"
)

type Time struct{
	EggTimes     int
	CubTimes     int
	BadeggTimes  int
	AbandonTimes int			//5
	CreateTime   string      
}

type Number struct{
	EggNumber    int
	CubNumber    int		//4
	BadeggNumber int 
	Number       int 
}
type ID struct{
	Id int					//自动分配 可不填 查找 更新 删除 不要填
	CageID       int 		/* CageID HouseID 唯一标识 */
	HouseID      int 
}

type Cage struct {        
	Statu        string    
	StatuDays    int     //2
	Time Time
	Number Number
	ID ID
}

// 5+4+3+2 = 14
var cage_table string = "cage"


func (cage *Cage) Find(id ID) error {
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ",cage_table ,id.HouseID ,id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(&cage.ID.Id,&cage.ID.HouseID,&cage.ID.CageID,&cage.Number.Number,  // 4
				&cage.Statu,&cage.StatuDays,&cage.Time.EggTimes,&cage.Number.EggNumber,					//4
				&cage.Time.CubTimes,&cage.Number.CubNumber,	&cage.Time.BadeggTimes,	//3
				&cage.Number.BadeggNumber,	&cage.Time.AbandonTimes,&cage.Time.CreateTime,//3
	)
	if err != nil {
		return err
	}
	return nil
}
func (cage *Cage) Update(c Cage,oldID ID) error {
	db := core.GetSQLDB()
	err1 := cage.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", cage_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,number=%d,",c.ID.HouseID,c.ID.CageID,c.Number.Number)
	sql += fmt.Sprintf("statu='%s',statu_days=%d,egg_times=%d,egg_number=%d,",c.Statu,c.StatuDays,c.Time.EggTimes,c.Number.EggNumber)
	sql += fmt.Sprintf("cub_times=%d,cub_number=%d,badegg_times=%d, ",c.Time.CubTimes,c.Number.CubNumber,c.Time.BadeggTimes)
	sql += fmt.Sprintf("badegg_number=%d,abandon_times=%d,create_time='%s' ",c.Number.BadeggNumber,c.Time.AbandonTimes,c.Time.CreateTime)
	sql += fmt.Sprintf(" where house_id=%d and cage_id=%d",oldID.HouseID,oldID.CageID)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// err == nil
func (cage *Cage) Insert(c Cage) error {
	db := core.GetSQLDB()

	err := cage.Find(c.ID)
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s ", cage_table)
	sql += fmt.Sprintf("values(%d,%d,%d,%d,'%s',%d,%d,%d,%d,%d,%d,%d,%d,'%s')",
		cage.ID.Id,cage.ID.HouseID,cage.ID.CageID,cage.Number.Number,
		cage.Statu,cage.StatuDays,cage.Time.EggTimes,cage.Number.EggNumber,
		cage.Time.CubTimes,cage.Number.CubNumber,cage.Time.BadeggTimes,
		cage.Number.BadeggNumber,cage.Time.AbandonTimes,cage.Time.CreateTime,
	)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (cage *Cage) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := cage.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", cage_table)
	_, err := db.Exec(sql,id.HouseID,id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (cage *Cage)FindList()([]Cage,error){
	db := core.GetSQLDB()

	var s []Cage
	sql := fmt.Sprintf("select * from %s ",cage_table)
	fmt.Printf("sql: %v\n", sql)
	rows,err:=db.Query(sql)
	if err!=nil{
		return nil,err
	}
	
	for rows.Next(){
		temp :=Cage{}
		err := rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.Number.Number,  
			&temp.Statu,
			&temp.StatuDays,
			&temp.Time.EggTimes,
			&temp.Number.EggNumber,					
			&temp.Time.CubTimes,
			&temp.Number.CubNumber,	
			&temp.Time.BadeggTimes,	
			&temp.Number.BadeggNumber,
			&temp.Time.AbandonTimes,
			&temp.Time.CreateTime,
		)
		if err!=nil{
			return nil,err
		}
		s = append(s,temp)
	}
	return s,nil
}