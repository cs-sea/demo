package discover

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"
)

type ConsulService struct {
	client *api.Client
}

func NewConsulService() *ConsulService {
	client, _ := api.NewClient(api.DefaultConfig())

	return &ConsulService{client}
}

func (s *ConsulService) Register(name, ip string, port int) error {
	// check service
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d", ip, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}

	rand.Seed(time.Now().Unix())
	// reg service
	registration := &api.AgentServiceRegistration{
		ID:                strconv.Itoa(rand.Int()),
		Name:              name,
		Tags:              []string{name},
		Port:              port,
		Address:           ip,
		EnableTagOverride: false,
		Check:             check,
	}

	return s.client.Agent().ServiceRegister(registration)
}

func (s *ConsulService) Get(name string) ([]*api.AgentService, error) {

	validService := make([]*api.AgentService, 0)
	services, err := s.client.Agent().Services()
	if err != nil {
		return nil, err
	}

	for _, v := range services {
		if v.Service == name {
			validService = append(validService, v)
		}
	}

	if len(validService) == 0 {
		return nil, errors.New("not valid services")
	}

	return validService, nil
}

func (s *ConsulService) Deregister(id string) error {
	return s.client.Agent().ServiceDeregister(id)
}

func (s *ConsulService) DeregisterAll() error {
	services, err := s.client.Agent().Services()

	if err != nil {
		return err
	}

	for _, service := range services {
		err = s.client.Agent().ServiceDeregister(service.ID)
		if err != nil {
			fmt.Errorf("error%v", err)
		}
	}

	return nil
}
