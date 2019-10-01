# McAugur
Minecraft服务器QQ群算命/占卜插件

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
                "可以白嫖铁块",
                "可以看到bs女装"
            ],
            "BadEvents": [
                "会被绑架去维修铁塔",
                "会被绑架去修分类机"
            ]
        },
        {
            "Name": "柚子村",
            "GoodEvents": [
                "会遇到DJ WCHF"
            ],
            "BadEvents": [
                "白嫖村民会被ZBin打死",
                "刷木头会弄坏树场"
            ]
        },
        {
            "Name": "花の岛",
            "BadEvents": [
                "会被离岛无条件拍死",
                "会被离岛抓去修分类机"
            ],
            "GoodEvents": [
                "嫖花不会被发现",
                "嫖钻石不会被发现"
            ]
        },
        {
            "Name": "螭水领",
            "GoodEvents": [
                "可以泡温泉"
            ],
            "BadEvents": [
                "会在雨林中迷路"
            ]
        }
    ],
    "GoodEvents": [
        "白嫖不会被发现"
    ],
    "BadEvents": [
        "白嫖会被抓现行",
        "会被雷劈死",
        "会被服务器延时卡进墙壁闷死"
    ],
    "Group": 000000000
    "Source":用户:密码@tcp(地址:端口)/库名
}
```
## 感谢
[BaiMeow](https://github.com/MscBaiMeow)（插件作者）
[Tnze](https://github.com/Tnze)（SDK开发者）
[Miaoscraft](https://miaoscraft.cn)（组织）