package main
import (
  "log"
  "github.com/FKouhai/cider-cli/methods"
  "github.com/FKouhai/cider-cli/config"
   "encoding/json"
)
type testStruct struct {
	Status string `json:"status"`
}
func main() {
	log.Println("Hello World")
	var c config.Config
	var t *testStruct
	setter := c.SetConfig()
	ep := "http://" + setter.Host + ":" + setter.Port + "/api/v1/playback/active"
	log.Println(ep)
	testing_ep, err := methods.Connect(ep,setter.Token,"GET"); if err != nil {
		log.Println(err)
	}
	st := json.Unmarshal(testing_ep,&t); if st != nil {
		log.Println(st)
	}
	log.Println(t.Status)
}
