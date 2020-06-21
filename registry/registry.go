package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joincloud/home-platform/home-services/conf"
	"io/ioutil"
	"net/http"
)

func Register() {
	url := conf.Configs.Home.Platform.Addr + conf.Configs.Home.Platform.RegisterRoute
	nodeJSON, _ := json.Marshal(conf.Configs.Home.Services.Node)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nodeJSON))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
