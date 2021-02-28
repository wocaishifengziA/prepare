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
		getResp *clientv3.GetResponse
		watchStartRevision int64
		watcher clientv3.Watcher
		watchRespChan <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
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

	// 模拟变换
	go func() {
		for{
			kv.Put(context.TODO(), "/test/job/1", "adad")
			kv.Delete(context.TODO(), "/test/job/1")
			time.Sleep(1 * time.Second)
		}
	}()

	// 读
	if getResp, err = kv.Get(context.TODO(), "/test/job/1"); err != nil {
		fmt.Println(err)
	}else {
		if len(getResp.Kvs) != 0 {
			fmt.Println("当前值：", getResp.Kvs)
		}
	}

	watchStartRevision = getResp.Header.Revision + 1
	watcher = clientv3.NewWatcher(client)
	watcher = watcher
	fmt.Println("从该版本向后监听:", watchStartRevision)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5 * time.Second, func() {
		cancelFunc()
	})

	watchRespChan = watcher.Watch(ctx, "/test/job/1", clientv3.WithRev(watchStartRevision))
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}
}

