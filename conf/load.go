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
	GoodEvents []string `json:"GoodEvents"`
	BadEvents  []string `json:"BadEvents"`
	Group      int64    `json:"Group"`
	Source     string   `json:"Source"`
}

//Place is a type to display places
type Place struct {
	Name       string   `json:"Name"`
	GoodEvents []string `json:"GoodEvents"`
	BadEvents  []string `json:"BadEvents"`
}

//Events includes two type of events
type Events struct {
	GoodEvents []string `json:"GoodEvents"`
	BadEvents  []string `json:"BadEvents"`
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
