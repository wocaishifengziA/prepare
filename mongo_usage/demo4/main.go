package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64	`bson:"startTime"`
	EndTime int64	`bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName string	`bson:"jobName"` // 任务名
	Command string `bson:"command"` // shell命令
	Err string `bson:"err"` // 脚本错误
	Content string `bson:"content"`// 脚本输出
	TimePoint TimePoint `bson:"timePoint"`// 执行时间点
}

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"`	// JobName赋值为job10
}

func main() {
	// mongodb读取回来的是bson, 需要反序列为LogRecord对象
	var (
		client *mongo.Client
		err error
		database *mongo.Database
		collection *mongo.Collection
		cond *FindByJobName
		cursor *mongo.Cursor
		record *LogRecord
	)
	// 1, 建立连接
	option := options.Client().ApplyURI("mongodb://106.75.130.240:27017")
	option.SetConnectTimeout(5 * time.Second)
	if client, err = mongo.Connect(context.TODO(), option); err != nil {
		fmt.Println(err)
		return
	}

	// 2, 选择数据库my_db
	database = client.Database("cron")

	// 3, 选择表my_collection
	collection = database.Collection("log")

	// 4, 按照jobName字段过滤, 想找出jobName=job10, 找出5条
	cond = &FindByJobName{JobName: "job10"}	// {"jobName": "job10"}

	//// 5, 查询（过滤 +翻页参数）
	//if cursor, err = collection.Find(context.TODO(), cond); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	limit := int64(2)
	skip := int64(0)
	if cursor, err = collection.Find(context.TODO(), cond, &options.FindOptions{Skip: &skip, Limit: &limit}); err != nil {
		fmt.Println(err)
		return
	}

	// 延迟释放游标
	defer cursor.Close(context.TODO())

	// 6, 遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义一个日志对象
		record = &LogRecord{}

		// 反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		// 把日志行打印出来
		fmt.Println(*record)
	}
}
