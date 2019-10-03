package data

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var db *sql.DB

//Open database
func Open(source string) error {
	var err error
	db, err = sql.Open("mysql", source)
	return err
}

//QQGetID get id using qq
func QQGetID(qq int64) (uuid uuid.UUID, err error) {
	err = db.QueryRow("SELECT UUID FROM users WHERE QQ=?", qq).Scan(&uuid)
	return
}

//IDGetName is copied from sis
func IDGetName(uuid uuid.UUID) (string, error) {
	data, status, err := get("https://sessionserver.mojang.com/session/minecraft/profile/" + hex.EncodeToString(uuid[:]))
	if err != nil {
		return "", err
	}
	defer data.Close()

	// 检查返回码
	if status != 200 {
		err = fmt.Errorf("服务器状态码非200: %v", status)
	}

	var resp struct{ Name string }
	// 解析json返回值
	err = json.NewDecoder(data).Decode(&resp)
	if err != nil {
		return "", err
	}

	return resp.Name, nil
}

//QQGetName get name using QQ
func QQGetName(qq int64) (string, error) {
	id, err := QQGetID(qq)
	if err != nil {
		return "", fmt.Errorf("获取UUID失败:%v", err)
	}
	name, err := IDGetName(id)
	if err != nil {
		return "", fmt.Errorf("获取Name失败: %v", err)
	}
	return name, err
}

// 发送GET请求
func get(url string) (io.ReadCloser, int, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	// Golang默认的User-agent被屏蔽了
	request.Header.Set("User-agent", "SiS")

	// 发送Get请求
	resp, err := new(http.Client).Do(request)
	if err != nil {
		return nil, 0, err
	}

	return resp.Body, resp.StatusCode, nil
}
