package index

import (
	"dakesolo.com/gg/share"
	"dakesolo.com/gg/unit"
	"fmt"
)

type User struct {
	ID          int64 `gorm:"primary_key:id"`
	ExpressName string
	OrderNumber string
}

func GetIndex(b *unit.Context) string {

	/*if err == nil {
		fmt.Println(err)
	}*/
	var user = User{
		ExpressName: "99999999",
		OrderNumber: "123",
	}
	//database.DB.Self.AutoMigrate(&User{})
	share.GetDB().Create(&user)
	fmt.Println("111")
	return "GetIndex"
}

func GetNav(b *unit.Context) string {
	//fmt.Println(share.Path.Config)
	tables, err := share.TableStoreClient.ListTable()
	if err != nil {
		fmt.Println("Failed to list table")
	} else {
		fmt.Println("List table result is")
		for _, table := range tables.TableNames {
			fmt.Println("TableName: ", table)
		}
	}

	return "getNav"
}

func GetBanner(context *unit.Context) string {
	/*share.Log.WithFields(logrus.Fields{
		"request": context.Request,
	}).Info("A group of walrus emerges from the ocean")*/
	//fmt.Println(unique.GetMD5())
	return "getBanner"
}
