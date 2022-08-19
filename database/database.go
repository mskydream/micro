package database

// type DbInstance struct {
// 	Db *gorm.DB
// }

// var Database DbInstance

// func Init() {
// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(&models.Todo{})
// 	Database = DbInstance{
// 		Db: db,
// 	}
// }
