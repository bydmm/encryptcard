package main

// CardPrototype 是card的原型
type CardPrototype struct {
	ID     int
	name   string
	Lines  string
	rarity int
}

// CardPrototypes 是卡片列表
// 卡片ID是算法是[0-9] + [0-9] + [0-9]
// 最小为0，最大为27
// 0和27最稀有，越靠近两边越稀有
var CardPrototypes = map[int]CardPrototype{
	0:  CardPrototype{ID: 0, name: "Zero", Lines: "我是开始，我是结束，我是阿赖耶，我是真理之门，我是一切的根源，我，是Zero", rarity: 5},
	27: CardPrototype{ID: 27, name: "42", Lines: "宇宙的奥秘，从此揭开", rarity: 5},

	1:  CardPrototype{ID: 1, name: "新桓结衣", Lines: "我不是你的老婆", rarity: 4},
	26: CardPrototype{ID: 26, name: "樱宁宁", Lines: "CPP又崩溃啦", rarity: 4},

	2: CardPrototype{ID: 2, name: "《计算机程序的构造和解释》(SICP)", Lines: "做完我的习题，再说你读过", rarity: 4},
	3: CardPrototype{ID: 3, name: "《黑客与画家》", Lines: "先实现一门语言，然后再开始实现功能。", rarity: 4},
	4: CardPrototype{ID: 4, name: "《代码大全》", Lines: "在挡子弹这件事情上，我很有自信", rarity: 4},
	5: CardPrototype{ID: 5, name: "《设计模式》", Lines: "四老外激动地站了起来", rarity: 4},
	6: CardPrototype{ID: 6, name: "《Unix网络编程》", Lines: "万物皆文件", rarity: 4},
	7: CardPrototype{ID: 7, name: "《TCP/IP详解》", Lines: "01111110", rarity: 4},
	8: CardPrototype{ID: 8, name: "《重构》", Lines: "写好测试，敏捷的重构你的微服务吧", rarity: 3},
	9: CardPrototype{ID: 9, name: "《编译原理技术和工具》", Lines: "屠龙之术不在乎有无龙可屠", rarity: 4},

	10: CardPrototype{ID: 10, name: "《C++ Primer》", Lines: "上个号称要七天精通C++的人造出了时光机", rarity: 3},
	11: CardPrototype{ID: 11, name: "《Python基础教程》", Lines: "人生苦短，我用大蟒蛇", rarity: 3},
	12: CardPrototype{ID: 12, name: "《Thinking in Java》", Lines: "老铁，来杯爪哇咖啡么", rarity: 3},
	13: CardPrototype{ID: 13, name: "《七天学会HTML》", Lines: "HTML是宇宙最好的语言", rarity: 3},
	14: CardPrototype{ID: 14, name: "《MYSQL从入门到跑路》", Lines: "DROP TABLE users;", rarity: 3},
	15: CardPrototype{ID: 15, name: "《React中文指南》", Lines: "尤雨溪给你多少钱?我马克扎波给你双倍", rarity: 3},
	16: CardPrototype{ID: 16, name: "《PHP和MySQL Web开发》", Lines: "我不是针对谁，我是说...", rarity: 3},
	17: CardPrototype{ID: 17, name: "《Web开发敏捷之道》", Lines: "听说硅谷的红宝石必须跑在轨道上", rarity: 3},
	18: CardPrototype{ID: 18, name: "《从0到1》", Lines: "作为村里唯一可以卖意大利炒面的餐馆，在几万亿的餐饮市场里我所向无敌", rarity: 3},
	19: CardPrototype{ID: 19, name: "《禅与摩托车维修艺术》", Lines: "你们程序员能不能不要再围观我修车了，该死，我说的是摩托车", rarity: 3},

	20: CardPrototype{ID: 20, name: "《复变函数》", Lines: "正在对你进行傅里叶展开", rarity: 4},
	21: CardPrototype{ID: 21, name: "《线性代数》", Lines: "He is the one", rarity: 4},
	22: CardPrototype{ID: 22, name: "《微积分学教程》", Lines: "抑制了房价快速上涨的趋势", rarity: 4},
	23: CardPrototype{ID: 23, name: "《数学分析》", Lines: "少年，你渴望力量吗？", rarity: 4},
	24: CardPrototype{ID: 24, name: "《实变函数》", Lines: "少年，你渴望力量吗？", rarity: 4},
	25: CardPrototype{ID: 25, name: "《泛函分析》", Lines: "少年，你渴望力量吗？", rarity: 4},
}
