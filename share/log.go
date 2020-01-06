package share

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/sirupsen/logrus"
	"os"
	"share/tool/unique"
)

var Log = logrus.New()

func InitLog() {
	// 设置日志格式为json格式
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	Log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	Log.Formatter = &logrus.JSONFormatter{}

	//Log.SetOutput(Log.Writer())

	// 添加自己实现的Hook
	Log.AddHook(&otsHook{})
}


type otsHook struct {}

func (hook *otsHook) Fire(entry *logrus.Entry) error {
	go func() {
		putRowRequest := new(tablestore.PutRowRequest)
		putRowChange := new(tablestore.PutRowChange)
		putRowChange.TableName = "log"
		putPk := new(tablestore.PrimaryKey)
		putPk.AddPrimaryKeyColumn("id", unique.GetMD5())
		putPk.AddPrimaryKeyColumn("service", "service")
		putPk.AddPrimaryKeyColumn("path", "path")
		putPk.AddPrimaryKeyColumn("level", entry.Level.String())

		putRowChange.PrimaryKey = putPk
		putRowChange.AddColumn("content", entry.Message)
		putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
		putRowRequest.PutRowChange = putRowChange
		_, err := TableStoreClient.PutRow(putRowRequest)

		if err != nil {
			fmt.Println("putrow failed with error:", err)
		} else {
			fmt.Println("putrow finished")
		}
	}()
	return nil
}

func (hook *otsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}



