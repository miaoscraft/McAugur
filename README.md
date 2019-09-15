# McAugur
Minecraft服务器QQ群算命/占卜插件

## 开始
release里面有打包好的
自己拿去用就是了

## 配置文件
第一次使用会自动生成配置文件，请编辑这个配置文件添加想要的事件
这是一个示例配置文件
```
{
    "Places": [
        {
            "Name": "ero岛",
            "PlaceEvents": [
                {
                    "Name": "白嫖铁块",
                    "Lucky": true
                },
                {
                    "Name": "偷拿钻石💎被当场抓获",
                    "Lucky": false
                }
            ]
        },
        {
            "Name": "柚子村",
            "PlaceEvents": [
                {
                    "Name": "白嫖村民被ZBin打死",
                    "Lucky": false
                }
            ]
        }
    ],
    "GeneralEvents": [
        {
            "Name": "白嫖会被抓现行",
            "Lucky": false
        },
        {
            "Name": "白嫖",
            "Lucky": true
        },
        {
            "Name": "被雷劈死",
            "Lucky": false
        }
    ],
    "Group": 1008610010
}
```

## 感谢
[BaiMeow](https://github.com/MscBaiMeow)（插件作者）
[Tnze](https://github.com/Tnze)（SDK开发者）
[Miaoscraft](https://miaoscraft.cn)（组织）