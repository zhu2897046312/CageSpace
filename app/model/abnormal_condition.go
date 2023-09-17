package model

import (
	"fmt"
	"server/app/core"
)

type AbnormalCondition struct {
	ID    ID //这里的 ID.Id 没有默认值 第一次插入时需要设置
	Refer string
	Time  string
	Statu int
}

var abnormalCondition_table string = "abnormal_condition"

func (a *AbnormalCondition) Find(id ID) error {
	db := core.GetSQLDB()

	sql := fmt.Sprintf("select *from %s where house_id = %d and cage_id = %d ", abnormalCondition_table, id.HouseID, id.CageID)
	r := db.QueryRow(sql)

	err := r.Scan(
		&a.ID.Id,
		&a.ID.HouseID,
		&a.ID.CageID,
		&a.Refer,
		&a.Time,
		&a.Statu,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *AbnormalCondition) Insert(c AbnormalCondition) error {
	db := core.GetSQLDB()

	err := a.Find(ID{HouseID: c.ID.HouseID, CageID: c.ID.CageID})
	if err == nil {
		//抛异常
		fmt.Print("已有该数据,不允许再次插入")
		return err
	}
	sql := fmt.Sprintf("insert into %s(id,house_id,cage_id,time,refer,statu)", abnormalCondition_table)
	sql += fmt.Sprintf(" values(%d,%d,%d,'%s','%s',%d)", c.ID.Id, c.ID.HouseID, c.ID.CageID, c.Time, c.Refer, c.Statu)
	_, errExec := db.Exec(sql)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (a *AbnormalCondition) Update(c AbnormalCondition, oldID ID) error {
	db := core.GetSQLDB()
	err1 := a.Find(oldID)

	if err1 != nil {
		//没找到了 houseid siteid相同的 行 抛异常
		return err1
	}

	sql := fmt.Sprintf("update %s set ", abnormalCondition_table)
	sql += fmt.Sprintf("house_id=%d,cage_id=%d,time='%s',refer='%s',statu=%d ", c.ID.HouseID, c.ID.CageID, c.Time, c.Refer, c.Statu)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (a *AbnormalCondition) Delete(id ID) error {
	db := core.GetSQLDB()

	err1 := a.Find(id)
	if err1 != nil {
		//没有找到
		return err1
	}
	sql := fmt.Sprintf("delete from %s where house_id=? and cage_id=?", abnormalCondition_table)
	_, err := db.Exec(sql, id.HouseID, id.CageID)
	if err != nil {
		return err
	}
	return nil
}

func (a *AbnormalCondition) FindList() ([]AbnormalCondition, error) {
	db := core.GetSQLDB()

	var s []AbnormalCondition
	sql := fmt.Sprintf("select * from %s ", abnormalCondition_table)
	fmt.Printf("sql: %v\n", sql)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := AbnormalCondition{}
		err = rows.Scan(
			&temp.ID.Id,
			&temp.ID.HouseID,
			&temp.ID.CageID,
			&temp.Refer,
			&temp.Time,
			&temp.Statu,
		)
		if err != nil {
			return nil, err
		}
		s = append(s, temp)
	}
	return s, nil
}
