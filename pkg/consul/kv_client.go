package consul

import "github.com/hashicorp/consul/api"

type KVClient struct {
	*api.KV
}
