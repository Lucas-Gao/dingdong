# 叮咚买菜助手

## 抢菜

### 通过抓包获取叮咚小程序的网络请求, 模仿接口调用, 配置文件中 header 和 form 里的必要参数需要根据自己抓包到的数据进行设置

### iphone (请在闲时可购买时抓包)
1. iphone上通过App Store安装Stream
2. 打开叮咚买菜微信小程序和Stream, Stream点击"开始抓包"
3. 回到叮咚买菜微信小程序, 商品加入购物车并结算提交(无需付款)使得小程序发起网络请求
4. 回到Stream停止抓包并查看抓包历史, 点击按域名, 会看到 maicai.api.ddxq.mobi 这个域名下的请求

[Stream 抓包教程(IOS)](https://www.jianshu.com/p/8a0fe2500f24)

[Charles 抓包教程(Mac)](https://www.jianshu.com/p/ff85b3dac157)

<img src="/assets/wechat.jpeg" width="300" alt="微信群" />

QQ群: 60768433 (golang & rust & blockchain & flutter 交流学习)

### 配置文件 config.json 设置

```js
snap_up; // 抢购 0关闭, 1 六点抢, 2 八点半抢, 3 六点和八点半都抢
pick_up_needed; // 闲时捡漏开关 false关闭 true打开 在抢购高峰期之外的时间捡漏 使用时需同时打开监视器
monitor_needed; // 监视器开关 监视是否有可配送时段
notify_needed; // 通知开关 发现有可配送时段时通知大家有可购商品 使用时需同时打开监视器
notify_interval; // 通知间隔 单位: 分钟
```

#### 通过接口修改配置文件

GET 请求 localhost:9999/set

| 参数            | 说明     |                       参数                       |
| :-------------- | :------- | :----------------------------------------------: |
| snap_up         | 抢购     | 0 关闭, 1 六点抢, 2 八点半抢, 3 六点和八点半都抢 |
| pick_up_needed  | 捡漏开关 |                  0 关闭 1 打开                   |
| monitor_needed  | 监视开关 |                  0 关闭 1 打开                   |
| notify_needed   | 通知开关 |                  0 关闭 1 打开                   |
| notify_interval | 通知间隔 |                 数字 单位: 分钟                  |

**例子**
| api | 说明 |
| :-----| :---- |
| localhost:9999/set?users=xxx,yyy | 添加需要通知的用户, 第一个是自己的 barkID, 其他为需要通知到的朋友(只能通知与你同属一个叮咚发货站点的用户) |
| localhost:9999/set?snap_up=1 | 六点抢购 |
| localhost:9999/set?pick_up_needed=1 | 打开捡漏,在抢购高峰期之外的时间捡漏(需同时打开监视器) |
| localhost:9999/set?monitor_needed=1 | 打开监视器 在抢购高峰期之外的时间监视是否可以配送 |
| localhost:9999/set?notify_needed=1 | 打开推送通知（需同时打开监视器） |
| localhost:9999/set?notify_interval=5 | 设置推送时间间隔(防止太过频繁) |

localhost:9999/set?users=xxx,yyy&snap_up=1&pick_up_needed=1&monitor_needed=1&notify_needed=1&notify_interval=5

## 可配送时段监听

<img src="/assets/effect.jpeg" width="300" alt="effect" />

### 当有可配送时段时, 发送通知到手机(只支持 ios)

### 注意事项

#### 1.安装 bark 得到自己的 barkID, 并将其写入配置文件中

<img src="/assets/user.jpeg" width="300" alt="user" />

#### 2.配置文件中 users 为一组需要通知的 bark userID

#### 3.bark 需打开允许通知

<img src="/assets/notify.jpeg" width="300" alt="notify" />

## 打包

打包到release目录下

```shell
make build # 默认打包为macOS darwin-amd64 Intel处理器
make build ARCH=arm64 # macOS M1处理器
make build OS=linux # linux
make build OS=windows # windows
```

## 执行程序

### macOS

在终端中执行

```ssh
./dingdong
```

### windows

在CMD中执行

```ssh
./dingdong.ext
```

### 版权说明

**本项目为 GPL3.0 协议，请所有进行二次开发的开发者遵守 GPL3.0 协议，并且不得将代码用于商用。**
