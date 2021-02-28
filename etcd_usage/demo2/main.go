package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putResp *clientv3.PutResponse
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

	// 写
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "112", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("reversion:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("PrevValue:", string(putResp.PrevKv.Value))
		}
	}
}

