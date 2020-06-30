package conf

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joincloud/home-platform/registry"
	log "github.com/sirupsen/logrus"
)

var (
	Configs        = &Conf{}
	bootstrapNodes []string
)

type Platform struct {
	Addr          string `json:"addr" yaml:"addr"`
	RegisterRoute string `json:"register-route" yaml:"register-route"`
}

type Services struct {
	Node           registry.Node     `json:"node" yaml:"node"`
	BootstrapNodes map[string]string `json:"bootstrap-nodes" yaml:"bootstrap-nodes"`
}

type Home struct {
	Platform Platform        `json:"platform" yaml:"platform"`
	Services Services        `json:"services" yaml:"services"`
	Files    map[string]File `json:"files" yaml:"files"`
}

type Conf struct {
	Home Home `json:"home" yaml:"home"`
}

type File struct {
	IsTemp bool   `json:"isTmp" yaml:"isTmp"`
	Dir    string `json:"dir" yaml:"dir"`
}

func Init(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, Configs)
	if err != nil {
		panic(err)
	}

	// region log config
	for s, f := range Configs.Home.Files {
		log.Infof("file dir: %s: %s", s, f.Dir)
	}

	// endregion
}

func GetBootStrapNodes() []string {
	if bootstrapNodes == nil {
		for _, ns := range Configs.Home.Services.BootstrapNodes {
			bootstrapNodes = append(bootstrapNodes, ns)
		}
	}

	return bootstrapNodes
}
