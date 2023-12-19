package main

import (
	consulApi "consul-contest/pkg/consul"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	ConsulAddr  = ":8500"
	ServicePort = 8000
	ServiceHost = "0.0.0.0"
	TaskConfig  = "golang/task.yml"
)

func main() {
	cli, err := consulApi.NewClient(ConsulAddr)
	if err != nil {
		log.Fatalf("can't initiate consul client: %+v\n", err)
	}
	consulStore := consulApi.NewKVClientAndLocal(cli)

	if err := Run(consulStore); err != nil {
		log.Fatalf("error starting service: %+v\n", err)
	}
}

func Run(cs *consulApi.KVClient) error {
	http.HandleFunc("/health", healthcheck)
	http.HandleFunc("/param", func(writer http.ResponseWriter, request *http.Request) {
		v, err := cs.GetKVConfig(TaskConfig)
		if err != nil {
			log.Println("Problem with variables")
			writer.WriteHeader(500)
		} else {
			writer.WriteHeader(200)
			name := v.Tasks[0].Name
			_, err := writer.Write([]byte(name))
			if err != nil {
				return
			}
		}
	})

	srvString := fmt.Sprintf("%s:%d", ServiceHost, ServicePort)
	if err := http.ListenAndServe(srvString, nil); err != nil {
		return err
	}
	return nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ok!",
	})
	if err != nil {
		log.Fatal(err)
	}
}
