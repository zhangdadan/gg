package share

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/spf13/viper"
)

var TableStoreClient *tablestore.TableStoreClient

func InitOTS()  {
	config := viper.New()
	config.SetConfigName("aliyun")
	config.AddConfigPath(Path.Config)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	TableStoreClient = tablestore.NewClient(config.GetString("ots.host"), config.GetString("ots.instance"), config.GetString("access.id"), config.GetString("access.secret"))
}




