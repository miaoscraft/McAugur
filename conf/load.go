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
	Places     []Place  `json:"Places"`
	GoodEvents []Event  `json:"GoodEvents"`
	BadEvents  []Event  `json:"BadEvents"`
	ResetCmds  []string `json:"ResetCmds"`
	Group      int64    `json:"Group"`
	Source     string   `json:"Source"`
	Server     string   `json:"Server"`
	PassWd     string   `json:"PassWd"`
}

//Place is a type to display places
type Place struct {
	Name       string  `json:"Name"`
	GoodEvents []Event `json:"GoodEvents"`
	BadEvents  []Event `json:"BadEvents"`
}

//Event includes events and Cmd
type Event struct {
	Name string `json:"Name"`
	Cmd  string `json:"Cmd"`
}

//LoadConf is used to load Config
func LoadConf(filename string) (*Config, error) {
	c := new(Config)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) { // 若文件不存在
			if data, err = json.Marshal(c); err != nil {
				return nil, fmt.Errorf("生成配置文件失败: %v", err)
			}
			if err = ioutil.WriteFile(filename, data, 0666); err != nil {
				return nil, fmt.Errorf("创建配置文件失败: %v", err)
			}
			cqp.AddLog(cqp.InfoSuccess, "McAugur", "找不到配置文件，已自动添加，请编辑配置文件再开启酷Q")
		} else {
			return nil, fmt.Errorf("读取配置文件失败: %v", err)
		}
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return c, nil
}
