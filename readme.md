### Golang http server + consul KV integration

1. run docker
2. dun docker compose 
3. open http://localhost:8500
4. create file golang/task.yml
```yaml
tasks:
  - name: gen_ai_full_name
    limit: 10
```


---
Useful docs:  
[Consul KV](https://developer.hashicorp.com/consul/docs/dynamic-app-config/kv)  
[Consul github](https://github.com/hashicorp/consul/tree/main)
[Consul docker compose](https://developer.hashicorp.com/consul/tutorials/docker/docker-compose-datacenter)  
[Consul tls](https://developer.hashicorp.com/consul/tutorials/security/tls-encryption-secure)  
[Golang and yaml](https://betterprogramming.pub/parsing-and-creating-yaml-in-go-crash-course-2ec10b7db850
)

