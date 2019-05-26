package handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var fileCfg = "config.yml"

var data = `
keys: "xxxxx"
servers:
  app-server:
    hostname: "172.17.0.1"
  nginx-server:
    hostname: "172.17.0.1"
  tmp-server:
    hostname: "172.17.0.1"
`

// AppServerCfg : a set of servers
type AppServerCfg struct {
	Keys    string `yaml:"keys"`
	Servers struct {
		AppSrv struct {
			Hostname string `yaml:"hostname"`
		} `yaml:"app-server"`
		NgxSrv struct {
			Hostname string `yaml:"hostname"`
		} `yaml:"nginx-server"`
		TmpSrv struct {
			Hostname string `yaml:"hostname"`
		} `yaml:"tmp-server"`
	} `yaml:"servers"`
}

// InitCfg configuration file.
func InitCfg(fileCfg string) {
	data, err := ioutil.ReadFile(fileCfg)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal(err)
	}
	var cfg AppServerCfg
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Read configuration from file.", fileCfg)
	fmt.Println(&cfg.Servers)
}
