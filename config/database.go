package config

import (
	"go-gin/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:gjs199074@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// //配置MySQL连接参数
	// username := "root"      //账号
	// password := "gjs199074" //密码
	// host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	// port := 3306            //数据库端口
	// Dbname := "web3_test"   //数据库名
	// timeout := "10s"        //连接超时，10秒

	// //拼接dsn参数，这里使用Sprintf动态拼接
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// //连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("连接数据库失败, error=" + err.Error())
	// 	fmt.Println(db)
	// 	fmt.Println("连接成功")

	// }

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected successfully")
}
