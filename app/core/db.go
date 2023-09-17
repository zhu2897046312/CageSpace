package core
import(
	"gorm.io/gorm"
	"log"
	"fmt"
	"gorm.io/driver/mysql"
)

var db = InitMysql()

func GetDB() *gorm.DB {
	return db
}

func InitMysql() *gorm.DB{
	// 日志配置
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
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatal("initMysql gorm.Open err:", err)
	}
	  
	return db
}