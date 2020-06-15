package registry

import "net/http"

func Register() {
	for {
		resp, err := http.Post("http://127.0.0.1:8090")
		if err != nil {
			panic(err)
		}

		resp.Body.Close()
	}
}
