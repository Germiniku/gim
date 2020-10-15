package loadbalance

import (
	"gim-server/common/registry"
	"math/rand"
)

/*
	负载均衡结构体
	封装负载均衡算法
*/

type LoadBalance interface {
	GetService(services []*registry.Service) (service *registry.Service)
}

type loadBalance struct {
	index int
	mode  string
}

// Init 初始化负载均衡器
func InitLoadBalance(mode string) LoadBalance {
	return &loadBalance{
		index: 0,
		mode:  mode,
	}
}

// GetService 根据不同的模式获取服务实例
func (lb *loadBalance) GetService(services []*registry.Service) (service *registry.Service) {
	switch lb.mode {
	case "random":
		return lb.randomSelector(services)
	case "poll":
		return lb.pollingSelector(services)
	default:
		return lb.randomSelector(services)
	}
}

// randomSelector 随机算法
func (lb *loadBalance) randomSelector(services []*registry.Service) (service *registry.Service) {
	index := rand.Intn(len(services) - 1)
	return services[index]
}

// pollingSelector 轮询算法
func (lb *loadBalance) pollingSelector(services []*registry.Service) (service *registry.Service) {
	if lb.index > len(services)-1 {
		lb.index = 0
	}
	service = services[lb.index]
	lb.index++
	return
}
