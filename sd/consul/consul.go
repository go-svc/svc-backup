package consul

// 原始碼來自：https://github.com/go-kit/kit/blob/master/sd/consul/client.go

import (
	"github.com/hashicorp/consul/api"
)

// Client 包裹了一個 Consul API。
type Client interface {
	// 將此服務註冊到本地特工（Agent）。
	Register(r *api.AgentServiceRegistration) error

	// 將此服務從本地特工中移除。
	Deregister(r *api.AgentServiceRegistration) error

	// Service
	Service(service, tag string, passingOnly bool, queryOpts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error)

	// 回傳一個最底層的 Consul 客戶端。
	Client() *api.Client
}

type client struct {
	Consul *api.Client
}

// NewClient 會回傳一個實作 Client 介面的建構體，
// 這個建構體包裹了 Consul 客戶端。
func NewClient(conf *api.Config) (Client, error) {
	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &client{Consul: c}, nil
}

// Register 會將此服務註冊到本地特工（Agent）。
func (c *client) Register(r *api.AgentServiceRegistration) error {
	return c.Consul.Agent().ServiceRegister(r)
}

// Deregister 將此服務從本地特工中移除。
func (c *client) Deregister(r *api.AgentServiceRegistration) error {
	return c.Consul.Agent().ServiceDeregister(r.ID)
}

func (c *client) Service(service, tag string, passingOnly bool, queryOpts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.Consul.Health().Service(service, tag, passingOnly, queryOpts)
}

// Client 會回傳一個最底層的 Consul 客戶端。
func (c *client) Client() *api.Client {
	return c.Consul
}
