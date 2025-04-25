package discovery

import (
    "fmt"
    "log"

    consulapi "github.com/hashicorp/consul/api"
)

type ServiceRegistry struct {
    client *consulapi.Client
}

func NewServiceRegistry() (*ServiceRegistry, error) {
    config := consulapi.DefaultConfig()
    client, err := consulapi.NewClient(config)
    if err != nil {
        return nil, err
    }
