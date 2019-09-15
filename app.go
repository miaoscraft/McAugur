package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/miaoscraft/McAugur/conf"
)

//go:generate cqcfg -c .
// cqp: 名称: McAugur
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: Msc定制mc算命插件
func main() { /*此处应当留空*/ }

var config conf.Config

func init() {
	cqp.AppID = "cn.miaoscraft.mcaugur"
	cqp.Enable = onEnable
	cqp.GroupMsg = onGroupMsg
}

// msc护符
func talisman() {
	if r := recover(); r != nil {
		cqp.AddLog(cqp.Fatal, "McAuger", fmt.Sprintf("%v\n\n%s", r, debug.Stack()))
	}
}

func onEnable() int32 {
	defer talisman()

	// 读取配置
	filename := filepath.Join(cqp.GetAppDir(), "conf.json")
	c, err := conf.LoadConf(filename)
	if err != nil {
		cqp.AddLog(cqp.Error, "McAuger", err.Error())
	} else {
		config = *c
	}

	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer talisman()

	if fromGroup == config.Group && msg == "算命" {
		return augur(fromQQ)
	}

	return 0
}

// 若出现错误则panic（正常情况永不panic）
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// 占卜术式
func augur(fromQQ int64) int32 {
	//启动术式
	hash := md5.New()

	//将被占卜物代入天命术式
	must(binary.Write(hash, binary.BigEndian, fromQQ)) // never returns a err
	y, m, d := time.Now().Date()
	must(binary.Write(hash, binary.BigEndian, y))
	must(binary.Write(hash, binary.BigEndian, m))
	must(binary.Write(hash, binary.BigEndian, d))

	//用天命开始占卜
	destiny := hash.Sum(nil) //获取destiny 天命
	rand.Seed(int64(binary.BigEndian.Uint64(destiny)))

	//占卜获得以太坐标
	placeID := rand.Intn(len(config.Places))
	result := "去" + config.Places[placeID].Name

	//占卜具体细节
	events := append(config.GeneralEvents, config.Places[placeID].PlaceEvents...)
	e := events[rand.Intn(len(events))]

	if e.Lucky {
		result = "今日 吉：" + result + e.Name
	} else {
		result = "今日 凶：" + result + e.Name
	}

	cqp.SendGroupMsg(config.Group, result)
	return 0
}
