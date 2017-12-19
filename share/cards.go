package share

// CardPrototype 是card的原型
type CardPrototype struct {
	ID     int
	Name   string
	Lines  string
	Rarity int
}

// CardPrototypes 是卡片列表
// 卡片ID是算法是[0-9] + [0-9] + [0-9]
// 最小为0，最大为27
// 0和27最稀有，越靠近两边越稀有
var CardPrototypes = map[int]CardPrototype{
	0:  CardPrototype{ID: 0, Name: "Zero", Lines: "我是开始，我是结束，我是阿赖耶，我是真理之门，我是一切的根源，我，是Zero", Rarity: 5},
	27: CardPrototype{ID: 27, Name: "42", Lines: "宇宙的奥秘，从此揭开", Rarity: 5},

	1:  CardPrototype{ID: 1, Name: "新桓结衣", Lines: "我不是你的老婆", Rarity: 4},
	26: CardPrototype{ID: 26, Name: "樱宁宁", Lines: "CPP又崩溃啦", Rarity: 4},

	2: CardPrototype{ID: 2, Name: "《计算机程序的构造和解释》(SICP)", Lines: "做完我的习题，再说你读过", Rarity: 4},
	3: CardPrototype{ID: 3, Name: "《黑客与画家》", Lines: "先实现一门语言，然后再开始实现功能。", Rarity: 4},
	4: CardPrototype{ID: 4, Name: "《代码大全》", Lines: "在挡子弹这件事情上，我很有自信", Rarity: 4},
	5: CardPrototype{ID: 5, Name: "《设计模式》", Lines: "四老外激动地站了起来", Rarity: 4},
	6: CardPrototype{ID: 6, Name: "《Unix网络编程》", Lines: "万物皆文件", Rarity: 4},
	7: CardPrototype{ID: 7, Name: "《TCP/IP详解》", Lines: "01111110", Rarity: 4},
	8: CardPrototype{ID: 8, Name: "《重构》", Lines: "写好测试，敏捷的重构你的微服务吧", Rarity: 3},
	9: CardPrototype{ID: 9, Name: "《编译原理技术和工具》", Lines: "屠龙之术不在乎有无龙可屠", Rarity: 4},

	10: CardPrototype{ID: 10, Name: "《C++ Primer》", Lines: "上个号称要七天精通C++的人造出了时光机", Rarity: 3},
	11: CardPrototype{ID: 11, Name: "《Python基础教程》", Lines: "人生苦短，我用大蟒蛇", Rarity: 3},
	12: CardPrototype{ID: 12, Name: "《Thinking in Java》", Lines: "老铁，来杯爪哇咖啡么", Rarity: 3},
	13: CardPrototype{ID: 13, Name: "《七天学会HTML》", Lines: "HTML是宇宙最好的语言", Rarity: 3},
	14: CardPrototype{ID: 14, Name: "《MYSQL从入门到跑路》", Lines: "DROP TABLE users;", Rarity: 3},
	15: CardPrototype{ID: 15, Name: "《React中文指南》", Lines: "尤雨溪给你多少钱?我马克扎波给你双倍", Rarity: 3},
	16: CardPrototype{ID: 16, Name: "《PHP和MySQL Web开发》", Lines: "我不是针对谁，我是说...", Rarity: 3},
	17: CardPrototype{ID: 17, Name: "《Web开发敏捷之道》", Lines: "听说硅谷的红宝石必须跑在轨道上", Rarity: 3},
	18: CardPrototype{ID: 18, Name: "《从0到1》", Lines: "作为村里唯一可以卖意大利炒面的餐馆，在几万亿的餐饮市场里我所向无敌", Rarity: 3},
	19: CardPrototype{ID: 19, Name: "《禅与摩托车维修艺术》", Lines: "你们程序员能不能不要再围观我修车了，该死，我说的是摩托车", Rarity: 3},

	20: CardPrototype{ID: 20, Name: "《复变函数》", Lines: "正在对你进行傅里叶展开", Rarity: 4},
	21: CardPrototype{ID: 21, Name: "《线性代数》", Lines: "He is the one", Rarity: 4},
	22: CardPrototype{ID: 22, Name: "《微积分学教程》", Lines: "抑制了房价快速上涨的趋势", Rarity: 4},
	23: CardPrototype{ID: 23, Name: "《数学分析》", Lines: "少年，你渴望力量吗？", Rarity: 4},
	24: CardPrototype{ID: 24, Name: "《实变函数》", Lines: "少年，你渴望力量吗？", Rarity: 4},
	25: CardPrototype{ID: 25, Name: "《泛函分析》", Lines: "少年，你渴望力量吗？", Rarity: 4},
}
