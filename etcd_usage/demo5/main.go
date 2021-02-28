package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
		kvPair *mvccpb.KeyValue
	)

	config = clientv3.Config{
		Endpoints: []string{"106.75.130.240:2379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 读
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	}else {
		if len(delResp.PrevKvs) != 0 {
			for _, kvPair = range delResp.PrevKvs {
				fmt.Println("删除了：", string(kvPair.Key), string(kvPair.Value))
			}
		}
	}
}

