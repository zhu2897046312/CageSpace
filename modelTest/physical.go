package modeltest

// import "time"

type Setting struct {
	Id           int `gorm:"type:int"`
	LayeggTimes  int `gorm:"type:int"`
	EggDays      int `gorm:"type:int"`
	LayeggDays   int `gorm:"type:int"`
	AbandonTimes int `gorm:"type:int"`
	BadeggTimes  int `gorm:"type:int"`
}

type WhippingEgg struct {
	Id           uint   `gorm:"type:int"`
	EggNumber    int    `gorm:"type:int"`
	BadeggNumber int    `gorm:"type:int"`
	Time         string `gorm:"type:datetime"`
	CageID       int
	Cage         Cage `gorm:"foreignKey:CageID"`
	HouseID      int
	House        House `gorm:"foreignKey:HouseID"`
}

type Egg struct {
	Id      uint   `gorm:"type:int"`
	Time    string `gorm:"type:datetime"`
	Statu   int    `gorm:"type:int"`
	CageID  int
	Cage    Cage `gorm:"foreignKey:CageID"`
	HouseID int
	House   House `gorm:"foreignKey:HouseID"`
}

type Death struct {
	Id          int    `gorm:"type:int"`
	OldNumber   int    `gorm:"type:int"`
	YoungNumber int    `gorm:"type:int"`
	Time        string `gorm:"type:datetime"`
	CageID      int
	Cage        Cage `gorm:"foreignKey:CageID"`
	HouseID     int
	House       House `gorm:"foreignKey:HouseID"`
}

type Cub struct {
	Id         int    `gorm:"type:int"`
	CubNumber  int    `gorm:"type:int"`
	Time       string `gorm:"type:datetime"`
	UpdateTime string `gorm:"type:datetime"`
	Statu      int    `gorm:"type:int"`
	CageID     int
	Cage       Cage `gorm:"foreignKey:CageID"`
	HouseID    int
	House      House `gorm:"foreignKey:HouseID"`
}

type AbnormalCondition struct {
	Id      uint   `gorm:"type:int;primarykey;auto increment;comment:'主键'"`
	Refer   string `gorm:"type:varchar(500)"`
	Time    string `gorm:"type:datetime"`
	Statu   int    `gorm:"type:int"`
	CageID  int
	Cage    Cage `gorm:"foreignKey:CageID"`
	HouseID int
	House   House `gorm:"foreignKey:HouseID"`
}

type Abandon struct {
	Id      uint   `gorm:"type:int;primarykey;auto increment;comment:'主键'"`
	Time    string `gorm:"type:datetime"`
	CageID  int
	Cage    Cage `gorm:"foreignKey:CageID"`
	HouseID int
	House   House `gorm:"foreignKey:HouseID"`
}

type Cage struct {
	Id           int    `gorm:"auto increment"`
	CageID       int    `gorm:"index"`
	Number       int    `gorm:"type:int"`
	Statu        int    `gorm:"type:int"`
	StatuDays    int    `gorm:"type:int"`
	EggTimes     int    `gorm:"type:int"`
	EggNumber    int    `gorm:"type:int"`
	CubTimes     int    `gorm:"type:int"`
	CubNumber    int    `gorm:"type:int"`
	BadeggTimes  int    `gorm:"type:int"`
	BadeggNumber int    `gorm:"type:int"`
	AbandonTimes int    `gorm:"type:int"`
	CreateTime   string `gorm:"type:datetime"`
	HouseID      int
	House        House `gorm:"foreignKey:HouseID"`
}

type House struct {
	Id      int `gorm:"auto increment"`
	HouseID int `gorm:"index"`
	SiteID  int
	Site    Site `gorm:"foreignKey:SiteID"`
}

type Count struct {
	Id        int    `gorm:"type:int"`
	EggSum    int    `gorm:"type:int"`
	BadeggSum string `gorm:"type:varchar(255)"`
	CubSum    int    `gorm:"type:int"`
	DeathSum  int    `gorm:"type:int"`
	Time      string `gorm:"type:datetime"`
	SiteID    int
	Site      Site `gorm:"foreignKey:SiteID"`
}

type Site struct {
	Id   uint
	Name string
}
