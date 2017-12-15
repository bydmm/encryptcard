# Encryptcard - 区块娘

使用区块链的计算力工作证明去挖SRR，会产生怎么样的火花?

## 项目简述

本项目验证了比特币区块的产生和验证技术，并将其应用到了喜闻乐见的抽卡环节中。

用户可以通过挂机的方式获得卡片，比比谁的CPU更强

PS: 看看代码你就会明白，其实工作证明也是有运气的成分的。即使是比特币，只要你特别欧洲，你也是有概率用普通电脑几分钟就挖到的。只是这个概率嘛。。。。

## 项目地址

https://github.com/bydmm/encryptcard

## 下载地址

新版本都会发布在releases里的：https://github.com/bydmm/encryptcard/releases

## 项目进度

* 计算力工作证明 ✓

* 私钥和公钥的产生 ✓

* 数字签名 ✓

* 交易（✗）

* 抽卡动画 ✓

* 登场台词 ✓

* 在线验证（✗）

## 项目展望

以后会不会有一种基于区块的分布式存档技术，用户存档不需要完全存在服务器端，而是也放一份在本地，服务器根据签名机制可以信任这份存档，并且加载数据。

这样做有一些可能的好处：

* 换代理商之后，用户可以自主上传存档，而不会出现被上家公司拒不交出存档的问题。

* 一些单机游戏为了防破解做了全程联网，这样很影响游戏体验。那能不能将存档放在本地，但是由服务器商签名，这样即使被破解了，也无法存档。

* 刀剑神域这样的跨游戏的存档实现，比如用户至少可以继承自己的人物名称和捏脸设定之类的，不过具体的继承内容的决定权在网游开发商。

## 基本结构

```json
{
  "Version": "v0.0.1",
  "PubKey": "-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANB.......pXBzV4QOMxBl5C\nrwIDAQAB\n-----END RSA PUBLIC KEY-----\n",
  "Timestamp": "1513263844677925384",
  "RandNumber": "725",
  "Hard": "4",
  "CardID": "35d859ed1f30d9e19b76b120ca7d706506edfdd35ed7c88feafccb0003601050",
  "Signature": "8c79aa73e105fad3479......eb5b0f2a1aa5e2493a1"
}
```

## Version

区块娘版本，不同版本的区块可能挖矿难度不一样，核心算法也不同

## PubKey

用户公钥(yue四声)

## Timestamp

卡被挖出的时间戳

## RandNumber

随机数，某个时间戳内为了多次重试，没有时间戳就无法挖卡了

## Hard

难度系数，

## CardID

CardID = 哈希(Version + PubKey + Timestamp + RandNumber + Hard)

任何人都可以通过这个验证这张卡的真实性，到底是不是挖出来的，还是随便乱写的。

这个也算是这个挖卡概念的技术核心，控制了出卡率只和用户的硬件水平有关。

假设CardID为：15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d680000000199

* 0000000 基础难度系数，0越多总体的难度提升（这个位数，不同的版本可能会变）

* 001 卡id，设定为0越多卡越稀有, 这是一张SSR

* 9  攻击, 纯属娱乐，未来对战系统可以自行实践

* 9  防御，纯属娱乐，未来对战系统可以自行实践

按照设想，挖卡程序只负责让卡有序产生，各种各样的应用可以根据这个卡的属性去自行设计任何游戏。

算是个神奇的开源社区的设想？

PS: 缺点是现在的hash算法sha256已经被ASIC矿机给优化。。但是我想那群挖币的应该不至于有空挖这个。。

PS2: 如果矿机挖这个，那么说明挖这卡的价值要大于挖比特币。。。

## Signature

普通的数字签名

首先是拥有者对卡（区块）签名

Signature = 签名函数(private_key, CardID)

交易者验证这张卡是不是真的来自于拥有者

CardID = 验证函数(PubKey, Signature)

## 卡交易（未实现）

```json
{
	"pubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2",
	"timestamp": 1974545345345,
	"randNumber": 6653,
	"cardBlock": "15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800000009004",
	"signature": "dsfsdf34515c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800",
	"ownerPubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2"
}
```

交易过的卡多一个字段：ownerPubkey

## 二次签名

signature = 签名函数(创造者的private_key, (ownerPubkey + CardID))

## 验证函数

ownerPubkey + cardBlock = 验证函数(pubkey, signature)

由于这个项目只是区块，不是链，也没有全局分布式账本，所以一张卡只允许交易一次了。。

因为第二次交易这张卡，很明显需要第一个用户的私钥，那不太现实。。。