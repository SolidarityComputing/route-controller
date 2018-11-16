package service

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"strconv"
	"github.com/kfcoding-container-api/configs"
	"log"
	"sync"
)

type EtcdService struct {
	EctdClientV3 *clientv3.Client
}

var once sync.Once
var myEtcdClient *EtcdService

// NewEtcdService, 创建新的EtcdService
func NewEtcdService() *EtcdService {
	once.Do(func() {
		var err error
		var cfg clientv3.Config
		if configs.EtcdUsername != "" {
			cfg = clientv3.Config{
				Endpoints:   configs.EtcdEndPoints,
				DialTimeout: configs.EtcdTimeout,
				Username:    configs.EtcdUsername,
				Password:    configs.EtcdPassword,
			}
		} else {
			cfg = clientv3.Config{
				Endpoints:   configs.EtcdEndPoints,
				DialTimeout: configs.EtcdTimeout,
			}
		}
		ectdClientV3, err := clientv3.New(cfg)
		if err != nil {
			log.Fatal("NewEtcdService error: ", err)
		}

		myEtcdClient = &EtcdService{
			EctdClientV3: ectdClientV3,
		}
	})
	return myEtcdClient
}

// Put, etcd put
func (e *EtcdService) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.EtcdTimeout)
	resp, err := e.EctdClientV3.Put(ctx, key, val, opts...)
	cancel()
	return resp, err
}

// Get, etcd get
func (e *EtcdService) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.EtcdTimeout)
	resp, err := e.EctdClientV3.Get(ctx, key, opts ...)
	cancel()
	return resp, err
}

// Delete, etcd delete
func (e *EtcdService) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.EtcdTimeout)
	resp, err := e.EctdClientV3.Delete(ctx, key, opts ...)
	cancel()
	return resp, err
}

// CheckExist, etcd check if key exist
func (e *EtcdService) CheckExist(id string) (bool, error) {
	res, err := e.Get(id)
	if err == nil && res.Count > 0 {
		return true, nil
	}
	return false, err
}

// GetErrorType, etcd error type
func (e *EtcdService) GetErrorType(err error) int {
	arr := strings.Split(err.Error(), ":")
	if len(arr) <= 0 {
		return -1
	}
	code, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
	return code
}
