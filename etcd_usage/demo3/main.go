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
		getResp *clientv3.GetResponse
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
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1", clientv3.WithKeysOnly()); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(getResp.Kvs, getResp.Count)
	}

}

