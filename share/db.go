package share

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDB() {

	config := viper.New()
	config.SetConfigName("db")
	config.AddConfigPath(Path.Config)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		config.Get("database.user"),
		config.Get("database.pwd"),
		config.Get("database.host"),
		config.Get("database.dbname"),
		true,
		// "Asia%2FShanghai",  // 必须是 url.QueryEscape 的
		"Local",
	))
	if err != nil {
		//log.Fatalf("数据库连接失败. 数据库名字: %s. 错误信息: %s", name, err)
		fmt.Println(err)
	} else {
		//log.Fatalf("数据库连接成功, 数据库名字: %s", name)
		fmt.Println("连接成功")
	}
}

func GetDB() *gorm.DB {
	return db
}





