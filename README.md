# McAugur
Minecraft服务器QQ群算命/占卜插件

**敬告:这是专门写给Miaoscraft的插件，不考虑其他用户，所以配置可能会显得非常繁琐**
**有些关于本人的权限直接写代码里了，要用一定要改**
## 功能
- 通过QQ读取玩家ID，算命并返回
- 通过RCON执行命令实装算命效果
- 在一天结束后清除上一天的算命效果
## 开始
这是一个SiS的附属插件，用到了SiS的数据库（仅支持MySQL）
如果要体验完整功能首先要安装SiS然后自行编译
release里的是独立版本的算命不需要SiS，但是功能有一定的残缺
## 配置文件
第一次使用会自动生成配置文件，请编辑这个配置文件添加想要的事件
这是一个示例配置文件
```
{
    "Places": [
        {
            "Name": "ero岛",
            "GoodEvents": [
                {
                    "Name": "可以白嫖铁块"
                },
                {
                    "Name": "可以看到bs女装",
                    "Cmd": "execute as %s run title @a title [{\"text\":\"BS女装！！！\",\"color\":\"red\",\"bold\":true,\"italic\":true,\"underlined\":false,\"strikethrough\":false,\"obfuscated\":false}]"
                }
            ],
            "BadEvents": [
                {
                    "Name": "会被绑架去维修铁塔"
                },
                {
                    "Name": "会被绑架去修分类机"
                }
            ]
        },
        {
            "Name": "柚子村",
            "GoodEvents": [
                {
                    "Name": "会遇到DJ WCHF"
                }
            ],
            "BadEvents": [
                {
                    "Name": "白嫖村民会被ZBin打死"
                },
                {
                    "Name:": "刷木头会弄坏树场"
                }
            ]
        },
        {
            "Name": "花の岛",
            "BadEvents": [
                {
                    "Name": "会被离岛无条件拍死"
                },
                {
                    "Name": "会被离岛抓去修分类机"
                }
            ],
            "GoodEvents": [
                {
                    "Name": "嫖花不会被发现"
                },
                {
                    "Name": "嫖钻石不会被发现"
                }
            ]
        },
        {
            "Name": "螭水领",
            "GoodEvents": [
                {
                    "Name": "可以泡温泉"
                }
            ],
            "BadEvents": [
                {
                    "Name": "会在雨林中迷路"
                },
                {
                    "Name": "会被温泉自动门夹死",
                    "Cmd": "tag %s add csl_doorkiller"
                }
            ]
        },
        {
            "Name": "出生点",
            "GoodEvents": [
                {
                    "Name": "可以在坟前蹦迪",
                    "Cmd": ""
                },
                {
                    "Name": "招揽新人会招到肝帝"
                },
                {
                    "Name": "能找到隐藏宝藏"
                }
            ],
            "BadEvents": [
                {
                    "Name": "会被甜梅刺死"
                },
                {
                    "Name": "会碰到十八级袭击而死"
                }
            ]
        }
    ],
    "GoodEvents": [
        {
            "Name": "白嫖不会被发现"
        }
    ],
    "BadEvents": [
        {
            "Name": "白嫖会被抓现行"
        },
        {
            "Name": "会被雷劈死",
            "Cmd": "scoreboard players set %s lightning_bolt 8"
        },
        {
            "Name": "会被服务器延时卡进墙壁闷死"
        }
    ],
    "Group": ***********,
    "Source": "**********",
    "Server": "**********",
    "PassWd": "********",
    "ResetCmds": [
        "scoreboard objectives remove lightning_bolt",
        "scoreboard objectives add lightning_bolt dummy"
    ]
}
```
由于本人不熟悉mc的一些指令，实例可能会出错
## 感谢
[BaiMeow](https://github.com/MscBaiMeow)（插件作者）
[Tnze](https://github.com/Tnze)（SDK开发者）
[Miaoscraft](https://miaoscraft.cn)（组织）