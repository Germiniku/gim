package registry

import (
	"context"
	"errors"
	"time"
)

var (
	ErrDeRegisterFailed    = errors.New("service deregister failed")
	ErrRegisterFailed      = errors.New("service register failed")
	ErrNotFound            = errors.New("service not found")
	ErrMetaJsonFailed      = errors.New("service meta data json failed")
	ErrNoAvailedbleService = errors.New("no availble service")
)

type Registry interface {
	Register() error
	Discovry() error
	DeRegister() error
	ListServices() ([]*Service, error)
	GetService() (*Service, error)
}

type Service struct {
	Addr    string `json:"addr"`    // 服务地址
	Name    string `json:"name"`    // 服务名
	Version string `json:"version"` // 服务版本
}

// Service 服务元信息
type Options struct {
	Addr      string          `json:"addr"`    // 服务地址
	Name      string          `json:"name"`    // 服务名
	Version   string          `json:"version"` // 服务版本
	Endpoints []string        `json:"-"`       // 注册中心端口
	Context   context.Context `json:"-"`
	Timeout   time.Duration   `json:"-"`
	Mode      string          `json:"-"` // 负载均衡模式
}

type OptionFunc func(*Options)
