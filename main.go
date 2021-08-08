package main

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/MscBaiMeow/SimpleBot/bot"
	"github.com/MscBaiMeow/SimpleBot/driver"
	"github.com/MscBaiMeow/SimpleBot/handler"
	"github.com/miaoscraft/McAugur/conf"
	"github.com/miaoscraft/McAugur/data"
	"github.com/miaoscraft/McAugur/rcon"
)

var (
	b         *bot.Bot
	config    *conf.Config
	AugurData *data.AugurData
	clear     = make(chan bool)
)

func main() {
	enable()
	ctx := context.TODO()
	go cleanEff(ctx)
	b = bot.New(driver.NewWsDriver(config.Websocket, config.Token))
	b.Attach("message.group.normal", &handler.GroupMsgHandler{
		Priority: 1,
		F: func(MsgID int32, GroupID int64, FromQQ int64, Msg string) bool {
			if GroupID != config.Group || Msg != "算命" {
				return false
			}
			if _, err := b.SendGroupMsg(GroupID, Augur(FromQQ)); err != nil {
				log.Println(err)
			}
			return true
		},
	})
	b.Attach("message.group.normal", &handler.GroupMsgHandler{
		Priority: 2,
		F: func(MsgID int32, GroupID int64, FromQQ int64, Msg string) bool {
			if GroupID != config.Group || !isAdmin(FromQQ) {
				return false
			}
			switch Msg {
			case "/mcaugur reload":
				enable()
			case "/mcaugur clear":
				clear <- true
			}
			return true
		},
	})
}

func isAdmin(qq int64) bool {
	for _, v := range config.Admin {
		if v == qq {
			return true
		}
	}
	return false
}

func enable() {
	if c, err := conf.Load("conf.json"); err != nil {
		log.Fatal(err.Error())
	} else {
		config = c
	}
	if d, err := data.Load("data.json"); err != nil {
		log.Fatal(err)
	} else {
		AugurData = d
	}
	if err := data.Open(config.Source); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("成功连接数据库")
	}
	if err := rcon.Open(config.Server, config.PassWd); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("成功连接RCON")
	}
}

// 恐慌之神马斯特
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// 占卜术式
func Augur(fromQQ int64) string {
	//启动术式
	hash := md5.New()
	name, err := data.QQGetName(fromQQ)
	if err != nil {
		log.Println(err.Error())
		return "算命大失败"
	}
	if name == "" {
		return "Wait!!!"
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
	placeID := rand.Intn(len(AugurData.Places))

	//占卜具体细节
	goodevents := append(AugurData.GoodEvents, AugurData.Places[placeID].GoodEvents...)
	badevents := append(AugurData.BadEvents, AugurData.Places[placeID].BadEvents...)
	result := fmt.Sprintf("%s今天的仙气值（1-100）为 %d\n", name, luckindex)
	switch {
	case luckindex <= 15:
		result += "今日 大凶\n"
		result1 := badevents[rand.Intn(len(badevents))]
		result2 := badevents[rand.Intn(len(badevents))]
		result += "去" + AugurData.Places[placeID].Name + "不但" + result1.Name + "\n"
		result += "而且" + result2.Name
		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 15 && luckindex <= 45:
		result += "今日 凶\n"
		result1 := badevents[rand.Intn(len(badevents))]
		result2 := goodevents[rand.Intn(len(goodevents))]
		result += "去" + AugurData.Places[placeID].Name + result1.Name + "\n"
		result += "不过呢" + result2.Name
		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 45 && luckindex <= 55:
		result += "今日 平，无特殊事件"

	case luckindex >= 55 && luckindex <= 85:
		result += "今日 吉\n"
		result1 := goodevents[rand.Intn(len(goodevents))]
		result2 := badevents[rand.Intn(len(badevents))]
		result += "去" + AugurData.Places[placeID].Name + result1.Name + "\n"
		result += "但是要注意" + result2.Name

		runCmd(fmt.Sprintf(result1.Cmd, name))
		runCmd(fmt.Sprintf(result2.Cmd, name))

	case luckindex > 85:
		result += "今日 大吉大利\n"
		result1 := goodevents[rand.Intn(len(goodevents))]
		result += "去" + AugurData.Places[placeID].Name + result1.Name + "\n"
		result += "Today is your day!"
		runCmd(fmt.Sprintf(result1.Cmd, name))

	}
	return result
}

func runCmd(cmd string) {
	log.Println(cmd)
	resp, err := rcon.Cmd(cmd)
	if err != nil {
		log.Fatal(err.Error())
		b.SendGroupMsg(config.Group, "连接服务器失败，算命效果可能无法实装")
	}
	log.Println(resp)
}

//清除算命效果
func cleanEff(goctx context.Context) {
	now := time.Now()
	clean := func() {
		for _, v := range AugurData.ResetCmds {
			runCmd(v)
		}
		b.SendGroupMsg(config.Group, "清除已有算命效果完成")
	}
	// 等待12点
	select {
	case <-time.After(time.Until(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()))):
		clean()
	case <-goctx.Done():
		return
	}

	timer := time.NewTicker(time.Hour * 24)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			clean()
		case <-goctx.Done():
			return
		}
	}
}
