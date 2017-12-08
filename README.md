# encryptcard - 区块娘

使用区块链的计算力工作证明去挖SRR，会产生怎么样的火花?

## 基本结构

```json
{
	"version": "v0.0.1",
	"pubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2",
	"timestamp": 1974545345345,
	"randNumber": 6653,
	"cardBlock": "15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800000009004",
	"signature": "dsfsdf34515c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800"
}
```

## version

区块娘版本，不同版本的区块可能挖矿难度不一样，核心算法也不同

## pubkey

用户公钥(yue四声)

## timestamp

卡被挖出的时间戳

## randNumber

随机数，某个时间戳内为了多次重试，没有时间戳就无法挖卡了

## cardBlock

cardBlock = 哈希(version + pubkey + timestamp + randNumber + cardBlock)

任何人都可以通过这个验证这张卡的真实性，到底是不是挖出来的，还是随便乱写的。

这个也算是这个挖卡概念的技术核心，控制了出卡率只和用户的硬件水平有关。

假设cardBlock为：15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d680000000199

> 0000000 基础难度系数，0越多总体的难度提升（这个位数，不同的版本可能会变）

> 001 卡id，设定为0越多卡越稀有, 这是一张SSR

> 9  攻击, 纯属娱乐，未来对战系统可以自行实践

> 9  防御，纯属娱乐，未来对战系统可以自行实践

按照设想，挖卡程序只负责让卡有序产生，各种各样的应用可以根据这个卡的属性去自行设计任何游戏。

算是个神奇的开源社区的设想？

PS: 缺点是现在的hash算法sha256已经被ASIC矿机给优化。。但是我想那群挖币的应该不至于有空挖这个。。

PS2: 如果矿机挖这个，那么说明挖这卡的价值要大于挖比特币。。。

# signature

普通的数字签名

拥有者签名

signature = 签名函数(private_key, cardBlock)

交易者验证这张卡是不是真的来自于拥有者

cardBlock = 验证函数(pubkey, signature)

## 卡交易


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

signature = 签名函数(创造者的private_key, (ownerPubkey + cardBlock))

## 验证函数

ownerPubkey + cardBlock = 验证函数(pubkey, signature)

由于这个项目只是区块，不是链，也没有全局分布式账本，所以一张卡只允许交易一次了。。

因为第二次交易这张卡，很明显需要第一个用户的私钥，那不太现实。。。