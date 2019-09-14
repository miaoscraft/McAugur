# McAugur
Minecraft服务器QQ群算命/占卜插件

## 开始
release里面有打包好的
自己拿去用就是了

## 配置文件
第一次使用这个插件他会蹦你一次酷Q，然后给你整一个配置文件
这是一个实例配置文件
```
{
    "Places": [
        {
            "Name": "ero岛",
            "PlaceEvents": [
                {
                    "Name": "白嫖铁块",
                    "Lucky": true
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
按照实例修改就行了
## 工作原理
1. 把时间和qq号丢在一起，科学地计算md5
2. 把md5整成随机数源，算随机数
3. 根据随机数读取事件发生地点
4. 把该地点特有事件和通用事件丢一起再整一个随机数抽事件
5. 输出算命结果
## 感谢
[BaiMeow](https://github.com/MscBaiMeow)（插件作者）
[Tnze](https://github.com/Tnze)（SDK开发者）
[Miaoscraft](https://miaoscraft.cn)（组织）