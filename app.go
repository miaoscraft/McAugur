package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
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

func onEnable() int32 {
	config = conf.Loadconf()
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if fromGroup != config.Group || msg != "算命" {
		return 0
	}
	//启动术式
	hash := md5.New()
	//将被占卜物代入天命术式
	io.WriteString(hash, fmt.Sprint(fromQQ, time.Now().Day()))
	//获取destiny 天命
	destiny := hash.Sum(nil)
	//将天命翻译为占卜语言
	buf := bytes.NewBuffer(destiny)
	var intdestiny int64
	binary.Read(buf, binary.LittleEndian, &intdestiny)
	addInfo(fmt.Sprint(intdestiny))
	//将翻译过的天命代入占卜公式
	rand.Seed(intdestiny)
	//占卜获得地点ID
	placeID := rand.Intn(len(config.Places))
	result := "在" + config.Places[placeID].Name
	//二次占卜获得事件ID
	events := config.GeneralEvents
	events = append(events, config.Places[placeID].PlaceEvents...)
	EventID := rand.Intn(len(events))
	if events[EventID].Lucky == true {
		result = "今日 吉：" + result + events[EventID].Name
	} else {
		result = "今日 凶：" + result + events[EventID].Name
	}
	cqp.SendGroupMsg(config.Group, result)
	return 0
}

func addInfo(info string) {
	cqp.AddLog(cqp.Info, "McAugur", info)
}
