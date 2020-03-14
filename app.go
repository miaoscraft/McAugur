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
	"github.com/miaoscraft/McAugur/rcon"
)

//go:generate cqcfg -c .
// cqp: 名称: McAugur
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: Msc定制mc算命插件
func main() { /*此处应当留空*/ }

var config conf.Config
var clear = make(chan bool)
var exit chan bool

func init() {

	cqp.AppID = "cn.miaoscraft.mcaugur"

	cqp.Enable = onEnable
	cqp.GroupMsg = onGroupMsg
	cqp.GroupMemberIncrease = onGroupMemberIncrease
	cqp.Disable = onDisable

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
	} else {
		cqp.AddLog(cqp.Info, "McAuger", "成功连接数据库")
	}
	if err = rcon.Open(config.Server, config.PassWd); err != nil {
		cqp.AddLog(cqp.Error, "McAuger", err.Error())
	} else {
		cqp.AddLog(cqp.Info, "McAuger", "成功连接RCON")
	}
	exit = make(chan bool)
	go rmeff(exit)
	return 0
}

func onDisable() int32 {
	exit <- true
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer talisman()

	switch {
	case fromGroup == config.Group && msg == "算命":
		return augur(fromQQ)
	case fromGroup == config.Group && msg == "/mcaugur reload" && fromQQ == 1098105012:
		onDisable()
		onEnable()
		cqp.SendGroupMsg(config.Group, "Reload Completely")

	case fromGroup == config.Group && msg == "/mcaugur clear" && fromQQ == 1098105012:
		clear <- true
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
	defer talisman()
	//启动术式
	hash := md5.New()
	name, err := data.QQGetName(fromQQ)
	if err != nil {
		cqp.AddLog(cqp.Info, "McAuger", err.Error())
		cqp.SendGroupMsg(config.Group, "算命大失败")
		return 0
	}
	if name == "" {
		cqp.SendGroupMsg(config.Group, "Wait!!!")
		return 0
	}
	//将被占卜物代入天命术式
	y, m, d := time.Now().Date()
	_, err = hash.Write([]byte(name)) // never returns a err
	must(err)
	must(binary.Write(hash, binary.BigEndian, int64(y)))
	must(binary.Write(hash, binary.BigEndian, int64(m)))
	must(binary.Write(hash, binary.BigEndian, int64(d)))

	//用天命开始占卜
	destiny := hash.Sum(nil) //获取destiny 天命
	rand.Seed(int64(binary.BigEndian.Uint64(destiny)))

	//占卜获得以太坐标与幸运指数
	luckindex := rand.Intn(100) + 1
	placeID := rand.Intn(len(config.Places))

	//占卜具体细节
	goodevents := append(config.GoodEvents, config.Places[placeID].GoodEvents...)
	badevents := append(config.BadEvents, config.Places[placeID].BadEvents...)
	result := fmt.Sprintf("%s今天的仙气值（1-100）为 %d\n", name, luckindex)
	switch {
	case luckindex <= 15:
		result += "今日 大凶\n"
		result1 := badevents[rand.Intn(len(badevents))]
		result2 := badevents[rand.Intn(len(badevents))]
		result += "去" + config.Places[placeID].Name + "不但" + result1.Name + "\n"
		result += "而且" + result2.Name
		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 15 && luckindex <= 45:
		result += "今日 凶\n"
		result1 := badevents[rand.Intn(len(badevents))]
		result2 := goodevents[rand.Intn(len(goodevents))]
		result += "去" + config.Places[placeID].Name + result1.Name + "\n"
		result += "不过呢" + result2.Name
		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 45 && luckindex <= 55:
		result += "今日 平，无特殊事件"

	case luckindex >= 55 && luckindex <= 85:
		result += "今日 吉\n"
		result1 := goodevents[rand.Intn(len(goodevents))]
		result2 := badevents[rand.Intn(len(badevents))]
		result += "去" + config.Places[placeID].Name + result1.Name + "\n"
		result += "但是要注意" + result2.Name

		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 85:
		result += "今日 大吉大利\n"
		result1 := goodevents[rand.Intn(len(goodevents))]
		result += "去" + config.Places[placeID].Name + result1.Name + "\n"
		result += "Today is your day!"
		runCmd(fmt.Sprintf(result1.Cmd, name))

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

func runCmd(cmd string) {
	cqp.AddLog(cqp.Info, "McAuger", cmd)
	resp, err := rcon.Cmd(cmd)
	if err != nil {
		cqp.AddLog(cqp.Error, "McAuger", err.Error())
		cqp.SendGroupMsg(config.Group, "连接服务器失败，算命效果可能无法实装")
	}
	cqp.AddLog(cqp.Info, "McAuger", resp)
}

//清除算命效果
func rmeff(exit chan bool) {
	next := time.Now().Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(time.Now()))
	for {
		select {
		case <-t.C:
			for _, v := range config.ResetCmds {
				runCmd(v)
			}
			cqp.SendGroupMsg(config.Group, "清除已有算命效果完成")
			time.Sleep(time.Minute)
			next := time.Now().Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t.Reset(next.Sub(time.Now()))
		case <-clear:
			for _, v := range config.ResetCmds {
				runCmd(v)
			}
			cqp.SendGroupMsg(config.Group, "清除已有算命效果完成")
		case <-exit:
			t.Stop()
			return
		}
	}
}
