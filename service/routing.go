package service

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/kfcoding-container-api/configs"
	"github.com/kfcoding-container-api/model"
	"path"
)

// RoutingService, 接口，定义添加和删除规则
type RoutingService interface {
	AddRule(*model.RoutingBody) error
	DeleteRule(string) error
	DeleteRulePrefix(string) error
}

type RoutingTraefikService struct {
	etcdClient *EtcdService
}

// NewRoutingTraefikService, 返回新的RoutingTraefikService
func NewRoutingTraefikService(etcdClient *EtcdService) *RoutingTraefikService {
	return &RoutingTraefikService{
		etcdClient: etcdClient,
	}
}

// AddRule, 添加一条规则
func (service *RoutingTraefikService) AddRule(rule *model.RoutingBody) error {
	// set backend
	key := path.Join(configs.TraefikPrefix, "backends/", rule.Name, "/servers/1/url")
	value := rule.URL
	if _, err := service.etcdClient.EctdClientV3.Put(context.Background(), key, value); nil != err {
		return err
	}

	//set frontend
	key = path.Join(configs.TraefikPrefix, "frontends/", rule.Name, "/backend")
	value = rule.Name
	if _, err := service.etcdClient.EctdClientV3.Put(context.Background(), key, value); nil != err {
		return err
	}
	key = path.Join(configs.TraefikPrefix, "frontends/", rule.Name, "/routes/1/rule")
	value = rule.Rule
	if _, err := service.etcdClient.EctdClientV3.Put(context.Background(), key, value); nil != err {
		return err
	}

	return nil
}

// DeleteRule, 删除一条规则
func (service *RoutingTraefikService) DeleteRule(name string) error {
	// delete backend
	key := path.Join(configs.TraefikPrefix, "backends/", name, "/servers/1/url")
	if _, err := service.etcdClient.EctdClientV3.Delete(context.Background(), key); nil != err {
		return err
	}

	// delete frontend
	key = path.Join(configs.TraefikPrefix, "frontends/", name, "/backend")
	if _, err := service.etcdClient.EctdClientV3.Delete(context.Background(), key); nil != err {
		return err
	}
	key = path.Join(configs.TraefikPrefix, "frontends/", name, "/routes/1/rule")
	if _, err := service.etcdClient.EctdClientV3.Delete(context.Background(), key); nil != err {
		return err
	}
	return nil
}

// DeleteRulePrefix, 根据前缀删除规则
func (service *RoutingTraefikService) DeleteRulePrefix(name string) error {
	// 先检查 workspace-server 的转发规则存不存在
	key := path.Join(configs.TraefikPrefix, "frontends/", name, "/backend")
	exist, err := service.etcdClient.CheckExist(key)
	if err != nil {
		return err
	} else if !exist {
		return errors.New("Not found " + name)
	}

	// delete backend
	key = path.Join(configs.TraefikPrefix, "backends/", name)
	if _, err := service.etcdClient.EctdClientV3.Delete(context.Background(), key, clientv3.WithPrefix()); nil != err {
		return err
	}

	// delete frontend
	key = path.Join(configs.TraefikPrefix, "frontends/", name)
	if _, err := service.etcdClient.EctdClientV3.Delete(context.Background(), key, clientv3.WithPrefix()); nil != err {
		return err
	}
	return nil
}
