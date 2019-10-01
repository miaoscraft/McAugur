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
	"github.com/miaoscraft/McAugur/data"
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
	cqp.GroupMemberIncrease = onGroupMemberIncrease
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
	if err = data.Open(config.Source); err != nil {
		cqp.AddLog(cqp.Error, "McAuger", err.Error())
	}
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer talisman()
	switch {
	case fromGroup == config.Group && msg == "算命":
		return augur(fromQQ)

	case fromGroup == config.Group:
		return 0
	}
	return 0
}

// 恐慌之神马斯特
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// 占卜术式
func augur(fromQQ int64) int32 {
	//启动术式
	hash := md5.New()
	name, err := data.QQGetName(fromQQ)
	if err != nil {
		cqp.AddLog(cqp.Info, "McAuger", err.Error())
		cqp.SendGroupMsg(config.Group, "算命大失败")
		return 0
	}
	if name == "" {
		cqp.SendGroupMsg(config.Group, "你莫得白名单算什么命")
		return 0
	}
	//将被占卜物代入天命术式
	y, m, d := time.Now().Date()
	must(binary.Write(hash, binary.BigEndian, fromQQ+8)) // never returns a err
	must(binary.Write(hash, binary.BigEndian, int64(y)))
	must(binary.Write(hash, binary.BigEndian, int64(m)))
	must(binary.Write(hash, binary.BigEndian, int64(d)))

	//用天命开始占卜
	destiny := hash.Sum(nil) //获取destiny 天命
	rand.Seed(int64(binary.BigEndian.Uint64(destiny)))

	//占卜获得以太坐标与幸运指数
	luckindex := rand.Intn(100) + 1
	result := fmt.Sprintf("%s今天的仙气值（1-100）为 %d\n", name, luckindex)
	placeID := rand.Intn(len(config.Places))

	//占卜具体细节
	goodevents := append(config.GoodEvents, config.Places[placeID].GoodEvents...)
	badevents := append(config.BadEvents, config.Places[placeID].BadEvents...)
	switch {
	case luckindex <= 15:
		result += "今日 大凶\n"
		result += "去" + config.Places[placeID].Name + "不但" + badevents[rand.Intn(len(badevents))] + "\n"
		result += "而且" + badevents[rand.Intn(len(badevents))]
	case luckindex > 15 && luckindex <= 45:
		result += "今日 凶\n"
		result += "去" + config.Places[placeID].Name + badevents[rand.Intn(len(badevents))] + "\n"
		result += "不过呢" + goodevents[rand.Intn(len(goodevents))]
	case luckindex > 45 && luckindex <= 55:
		result += "今日 平，无特殊事件"
	case luckindex >= 55 && luckindex <= 85:
		result += "今日 吉\n"
		result += "去" + config.Places[placeID].Name + goodevents[rand.Intn(len(goodevents))] + "\n"
		result += "但是要注意" + badevents[rand.Intn(len(badevents))]
	case luckindex > 85:
		result += "今日 大吉大利\n"
		result += "去" + config.Places[placeID].Name + goodevents[rand.Intn(len(goodevents))] + "\n"
		result += "Today is your day!"
	}

	cqp.SendGroupMsg(config.Group, result)
	return 0
}

//给新群员占卜刷一下存在感
func onGroupMemberIncrease(subType, sendTime int32, fromGroup, fromQQ, beingOperateQQ int64) int32 {
	if fromGroup != config.Group {
		return 0
	}

	hash := md5.New()
	must(binary.Write(hash, binary.BigEndian, int64(beingOperateQQ)))
	//用天命开始占卜
	destiny := hash.Sum(nil) //获取destiny 天命
	rand.Seed(int64(binary.BigEndian.Uint64(destiny)))

	cqp.SendGroupMsg(config.Group, "欢迎新大佬\n我看你骨骼精奇，适合加入"+config.Places[rand.Intn(len(config.Places))].Name)
	return 0
}

func addEvents() int32 {
	return 0
}
