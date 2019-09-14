package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//Config is used to restore loaded config in McAugur
type Config struct {
	Places        []Place `json:"Places"`
	GeneralEvents []Event `json:"GeneralEvents"`
	Group         int64   `json:"Group"`
}

//Place is a type to display places
type Place struct {
	Name        string  `json:"Name"`
	PlaceEvents []Event `json:"PlaceEvents"`
}

//Event is a what will happan and xiong or ji
type Event struct {
	Name  string `json:"Name"`
	Lucky bool   `json:"Lucky"`
}

//Loadconf is used to load Config
func Loadconf() Config {
	var Config Config
	data, err := ioutil.ReadFile(".\\conf\\McAugur.json")
	if err != nil {
		if os.IsNotExist(err) {
			if data, err = json.Marshal(Config); err != nil {
				cqp.AddLog(cqp.Fatal, "McAugur", fmt.Sprint(err))
			}
			if err = ioutil.WriteFile(".\\conf\\McAugur.json", []byte(data), 0666); err != nil {
				cqp.AddLog(cqp.Fatal, "McAugur", fmt.Sprint(err))
			}
			cqp.AddLog(cqp.Fatal, "McAugur", "找不到配置文件，已自动添加，请于酷Q目录下conf文件夹中编辑配置文件再开启酷Q")
		} else {
			cqp.AddLog(cqp.Fatal, "McAugur", fmt.Sprint(err))
		}
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		cqp.AddLog(cqp.Fatal, "McAugur", fmt.Sprint(err))
	}
	return Config
}
