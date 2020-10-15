/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 10:18:13
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 13:10:24
 */
package etcd

import (
	"context"
	"encoding/json"
	"gim-server/common/registry"
	"gim-server/common/registry/loadbalance"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

type etcdRegistry struct {
	etcdCli          *clientv3.Client  // etcd 客户端
	availbleServices map[string]string // 当前可用的服务
	lock             sync.RWMutex
	opts             registry.Options
	lb               loadbalance.LoadBalance // 负载均衡器
}

// Init 初始化注册中心服务
func Init(opts ...registry.OptionFunc) registry.Registry {
	options := registry.Options{
		Context: context.Background(),
		Timeout: time.Millisecond * 1000,
	}
	for _, o := range opts {
		o(&options)
	}
	svc := new(etcdRegistry)
	svc.opts = options
	svc.availbleServices = make(map[string]string)
	// 初始化负载均衡器
	svc.lb = loadbalance.InitLoadBalance(svc.opts.Mode)
	// 初始化etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   options.Endpoints,
		DialTimeout: options.Timeout,
	})
	if err != nil {
		return nil
	}
	svc.etcdCli = cli
	return svc
}

// Register 服务注册
func (svc *etcdRegistry) Register() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	data, err := json.Marshal(&svc.opts)
	if err != nil {
		return registry.ErrMetaJsonFailed
	}
	_, err = svc.etcdCli.Put(ctx, svc.opts.Name, string(data))
	go WatchTimeOutCtx(ctx, cancel)
	if err != nil {
		return registry.ErrRegisterFailed
	}
	return
}

// watchServerChange 监听etcd服务的更新变化
func (svc *etcdRegistry) watchServerChange(resp *clientv3.GetResponse) {
	// 从GET对应的后续版本开始监听
	watchStartRevision := resp.Header.Revision + 1
	watchChan := svc.etcdCli.Watch(context.TODO(), svc.opts.Name, clientv3.WithPrefix(), clientv3.WithRev(watchStartRevision))
	// 处理监听事件
	for watchResp := range watchChan {
		for _, watchEvent := range watchResp.Events {
			switch watchEvent.Type {
			case mvccpb.PUT:
				// 判断当前服务列表当中是否已经拥有该服务
				// 如果没有则添加到服务列表当中
				if _, ok := svc.availbleServices[string(watchEvent.Kv.Key)]; !ok {
					svc.lock.Lock()
					svc.availbleServices[string(watchEvent.Kv.Key)] = string(watchEvent.Kv.Value)
					svc.lock.Unlock()
				}
			case mvccpb.DELETE:
				if _, ok := svc.availbleServices[string(watchEvent.Kv.Key)]; ok {
					svc.lock.Lock()
					delete(svc.availbleServices, string(watchEvent.Kv.Key))
					svc.lock.Unlock()
				}
			}
		}
	}
}

// Discovry 服务发现
func (svc *etcdRegistry) Discovry() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), svc.opts.Timeout)
	resp, err := svc.etcdCli.Get(ctx, svc.opts.Name, clientv3.WithPrefix())
	go WatchTimeOutCtx(ctx, cancel)
	if err != nil {
		return registry.ErrNotFound
	}
	// 遍历所有的服务添加到服务列表当中
	for _, item := range resp.Kvs {
		svc.availbleServices[string(item.Key)] = string(item.Value)
	}
	// 从该reversion向后监听变化事件
	go svc.watchServerChange(resp)
	return
}

// DeRegister 注销服务
func (svc *etcdRegistry) DeRegister() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), svc.opts.Timeout)
	resp, err := svc.etcdCli.Delete(ctx, svc.opts.Name)
	if err != nil {
		return registry.ErrDeRegisterFailed
	}
	go WatchTimeOutCtx(ctx, cancel)
	if resp.Deleted > 0 {
		return registry.ErrDeRegisterFailed
	}
	return
}

// ListServices 列出当前可用服务
func (svc *etcdRegistry) ListServices() (services []*registry.Service, err error) {
	for _, svc := range svc.availbleServices {
		var item *registry.Service
		if err = json.Unmarshal([]byte(svc), item); err != nil {
			continue
		}
		services = append(services, item)
	}
	if len(services) == 0 {
		return nil, registry.ErrNoAvailedbleService
	}
	return
}

func (svc *etcdRegistry) GetService() (service *registry.Service, err error) {
	services, err := svc.ListServices()
	if err != nil {
		return nil, registry.ErrNoAvailedbleService
	}
	service = svc.lb.GetService(services)
	return
}

// WatchTimeOutCtx 监听上下文超时
func WatchTimeOutCtx(ctx context.Context, cancelFunc context.CancelFunc) {
	select {
	case <-ctx.Done():
		cancelFunc()
	}
}
