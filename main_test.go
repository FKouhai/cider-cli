package main
import (
	"testing"
	"log"
	"encoding/json"
	"github.com/FKouhai/cider-cli/config"
	"github.com/FKouhai/cider-cli/methods"
)
type testStructs struct {
	Status string `json:"status"`
}
func TestActive(tes *testing.T){
	var c config.Config
	var t *testStructs
	setter := c.SetConfig()
	ep := "http://" + setter.Host + ":" + setter.Port + "/api/v1/playback/active"
	testing_ep, err := methods.Connect(ep,setter.Token,"GET"); if err != nil {
		log.Println(err)
	}
	st := json.Unmarshal(testing_ep,&t); if st != nil {
		log.Println(st)
	}
	log.Println(t.Status)
	got := t.Status
	if got != "ok" {
		tes.Errorf("Expected ok, got: %v", got)

	}

}
