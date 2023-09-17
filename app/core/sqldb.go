package core

import (  
    "database/sql"  
    "fmt"  
    _ "github.com/go-sql-driver/mysql" //导入mysql包   
)

var DB = InitDB()

func GetSQLDB() *sql.DB{
	return DB
}


func InitDB() *sql.DB{
	host := "172.21.179.2"
	//host := "localhost"
	port := "3306"
	database := "pigeon"
	//database := "mybatis"
	username := "root"
	password := "xike614"
	//password := "0608"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := sql.Open("mysql", args)  
	if err != nil {  
  	 	fmt.Println("数据库链接错误", err)  
   		return nil
	} 
	return db
}

