package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type AugurData struct {
	Places     []Place  `json:"Places"`
	GoodEvents []Event  `json:"GoodEvents"`
	BadEvents  []Event  `json:"BadEvents"`
	ResetCmds  []string `json:"ResetCmds"`
}

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

func Load(filename string) (*AugurData, error) {
	data := new(AugurData)

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("McAugur", "找不到算命数据文件")
		} else {
			return nil, fmt.Errorf("读取算命数据失败: %v", err)
		}
	}

	err = json.Unmarshal(f, data)
	if err != nil {
		return nil, fmt.Errorf("解析算命数据失败: %v", err)
	}
	return data, nil
}
