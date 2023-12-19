package consul

import (
	"consul-contest/pkg/kv"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"

	"github.com/hashicorp/consul/api"
)

var defaultConfig kv.TaskConfig

const (
	MaxAge      = 10 * time.Minute
	DefaultPath = "cfg/task.yml"
)

func NewClient(addr string) (*Client, error) {
	conf := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(conf)
	if err != nil {
		log.Println("error initiating new consul client: ", err)
		return &Client{}, err
	}

	return &Client{
		client,
	}, nil
}

func NewKVClient(c *Client) *KVClient {
	return &KVClient{
		c.KV(),
	}
}

func NewKVClientAndLocal(c *Client) *KVClient {
	// try to find local copy
	file, err := os.Open(DefaultPath)
	if err != nil {
		log.Fatalln("Couldn't find config.", err)
	}

	bb := make([]byte, 256)
	size, err := file.Read(bb)
	if err != nil {
		log.Fatalln("Couldn't read config.", err)
	}

	// remove extra bytes
	bb = bb[:size]
	cfg, err := unmarshal(bb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal config.", err)
	}

	defaultConfig = *cfg

	return &KVClient{
		c.KV(),
	}
}

func (k *KVClient) GetKVConfig(key string) (*kv.TaskConfig, error) {
	p, _, err := k.Get(key, &api.QueryOptions{
		UseCache: true,
		MaxAge:   MaxAge,
	})
	if err != nil {
		log.Println("error getting value from key: ", err)
		return nil, err
	}

	cfg, err := unmarshal(p.Value)
	if err != nil {
		log.Println("error parsing json: ", err)
		return nil, err
	}

	return cfg, nil
}

func unmarshal(b []byte) (*kv.TaskConfig, error) {
	var config kv.TaskConfig
	err := yaml.Unmarshal(b, &config)
	return &config, err
}
