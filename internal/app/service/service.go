package service

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/LIYINGZHEN/go-consul-example/configs"
	"github.com/go-redis/redis"
	consul "github.com/hashicorp/consul/api"
)

type Service struct {
	Name        string
	TTL         time.Duration
	RedisClient redis.UniversalClient
	ConsulAgent *consul.Agent
}

func New(addrs []string, ttl time.Duration) (*Service, error) {
	s := new(Service)
	s.Name = "go-consul-example"
	s.TTL = ttl
	s.RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: addrs,
	})

	ok, err := s.Check()
	if !ok {
		return nil, err
	}

	config := consul.DefaultConfig()
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	s.ConsulAgent = c.Agent()

	address := configs.C.Server.Host
	port, err := strconv.Atoi(configs.C.Server.Port)
	if err != nil {
		log.Fatalln(err)
	}

	serviceDef := &consul.AgentServiceRegistration{
		Name:    s.Name,
		Address: configs.C.Server.Host,
		Port:    port,
		Check: &consul.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/healthcheck", address, port),
			Interval: configs.C.Consul.Interval,
			Timeout:  configs.C.Consul.Timeout,
		},
	}
	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}

	return s, nil
}
