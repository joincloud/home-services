package conf

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joincloud/home-platform/registry"
)

var (
	Configs = &Conf{}
)

type Platform struct {
	Addr          string `json:"addr" yaml:"addr"`
	RegisterRoute string `json:"register-route" yaml:"register-route"`
}

type Services struct {
	Node registry.Node `json:"node" yaml:"node"`
}

type Home struct {
	Platform Platform `json:"platform" yaml:"platform"`
	Services Services `json:"services" yaml:"services"`
}

type Conf struct {
	Home Home `json:"home" yaml:"home"`
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
}
