package methods

import (
	"io"
	"log"
	"net/http"
)

func Connect(endpoint string, token string, method string) ([]byte, error){
	client := http.Client{}
	var body []byte
	if method != "GET" {
		//resp, err := client.Post()
	}
		req, err := http.NewRequest("GET", endpoint, nil); if err != nil {
			log.Fatalln(err)
			return nil,err
		}
		req.Header.Add("apptoken", token)
		resp, err := client.Do(req); if err != nil {
			log.Fatalln(err)
			return nil,err
		}
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body); if err != nil {
			log.Fatalln(err)
			return nil,err
		}
		return body,nil

}
