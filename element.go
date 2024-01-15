package doushu

import (
	"log"
)

type Element int

func (e Element) Category() (num int, first Element) {
	if e == Unknown {
		return 0, Unknown
	}
	if e >= Zi && e <= Hai {
		return int(Hai-Zi) + 1, Zi
	}
	if e >= Jia && e <= Gui {
		return int(Gui-Jia) + 1, Jia
	}

	switch {
	case e >= End:
		if e > End {
			log.Printf("invalid element: %d", e)
		}
		return 0, e
	case e >= Ziwei:
		return int(End - Ziwei), Ziwei
	case e >= Miao:
		return int(Xian-Miao) + 1, Miao
	case e >= Shui2ju:
		return int(Huo6ju-Shui2ju) + 1, Shui2ju
	case e >= Shengong:
		return 1, Shengong
	case e >= Minggong:
		return int(Xiongdi-Minggong) + 1, Minggong
	case e >= Chuyi:
		return int(Sanshi-Chuyi) + 1, Chuyi
	case e >= Zhengyue:
		return int(Layue-Zhengyue) + 1, Zhengyue
	case e >= Zi:
		return int(Hai-Zi) + 1, Zi
	case e >= Jia:
		return int(Gui-Jia) + 1, Jia
	case e >= Shu:
		return int(Zhu-Shu) + 1, Shu
	case e >= Mu:
		return int(Shui-Mu) + 1, Mu
	case e >= Nan:
		return int(Nv-Nan) + 1, Nan
	case e >= Yang:
		return int(YinXing-Yang) + 1, Yang
	default:
		log.Printf("invalid element: %d", e)
		return int(e), e
	}
}

func (e Element) Value() int {
	if e >= Zi && e <= Hai {
		return int(e - Zi)
	}
	if e >= Jia && e <= Gui {
		return int(e - Jia)
	}

	_, first := e.Category()
	return int(e - first)
}

func (e Element) String() string {
	return names[e]
}

func (e Element) Next(n int) Element {
	if n == 0 {
		return e
	}
	num, first := e.Category()
	return Element((e.Value()+n)%num) + first
}

func (e Element) Pre(n int) Element {
	if n == 0 {
		return e
	}
	num, first := e.Category()
	return Element((e.Value()-n+num*12)%num) + first
}

func (e Element) NextTo(n int) []Element {
	ret := make([]Element, n)
	for i := 0; i < n; i++ {
		ret[i] = e.Next(i)
	}
	return ret
}

func (e Element) PreTo(n int) []Element {
	ret := make([]Element, n)
	for i := 0; i < n; i++ {
		ret[i] = e.Pre(i)
	}
	return ret
}

const (
	Unknown Element = iota

	Yang    // 阳
	YinXing // 阴

	Nan // 男
	Nv  // 女

	Mu   // 木
	Huo  // 火
	Tu   // 土
	Jin  // 金
	Shui // 水

	Shu      // 鼠
	Niu      // 牛
	Hu       // 虎
	TuXiao   // 兔
	Long     // 龙
	She      // 蛇
	Ma       // 马
	YangXiao // 羊
	Hou      // 猴
	JiXiao   // 鸡
	Gou      // 狗
	Zhu      // 猪

	// 十天干
	Jia   // 甲
	Yi    // 乙
	Bing  // 丙
	Ding  // 丁
	WuGan // 戊
	Ji    // 己
	Geng  // 庚
	Xin   // 辛
	Ren   // 壬
	Gui   // 癸

	// 十二地支
	Zi   // 子
	Chou // 丑
	Yin  // 寅
	Mao  // 卯
	Chen // 辰
	Si   // 巳
	Wu   // 午
	Wei  // 未
	Shen // 申
	You  // 酉
	Xu   // 戌
	Hai  // 亥

	// 农历月
	Zhengyue // 正月
	Eryue    // 二月
	Sanyue   // 三月
	Siyue    // 四月
	Wuyue    // 五月
	LiuYue   // 六月
	Qiyue    // 七月
	Bayue    // 八月
	Jiuyue   // 九月
	Shiyue   // 十月
	Dongyue  // 冬月
	Layue    // 腊月

	// 农历日
	Chuyi    // 初一
	Chuer    // 初二
	Chusan   // 初三
	Chusi    // 初四
	Chuwu    // 初五
	Chuliu   // 初六
	Chuqi    // 初七
	Chuba    // 初八
	Chujiu   // 初九
	Chushi   // 初十
	Shiyi    // 十一
	Shier    // 十二
	Shisan   // 十三
	Shisi    // 十四
	Shiwu    // 十五
	Shiliu   // 十六
	Shiqi    // 十七
	Shiba    // 十八
	Shijiu   // 十九
	Ershi    // 二十
	Ershiyi  // 二十一
	Ershier  // 二十二
	Ershisan // 二十三
	Ershisi  // 二十四
	Ershiwu  // 二十五
	Ershiliu // 二十六
	Ershiqi  // 二十七
	Ershiba  // 二十八
	Ershijiu // 二十九
	Sanshi   // 三十

	// 妻财子禄十二宫
	Minggong // 命宫
	Fumu     // 父母
	Fude     // 福德
	Tianzhai // 田宅
	Guanlu   // 官禄
	Jiaoyou  // 交友
	Qianyi   // 迁移
	Jie      // 疾厄
	Caibo    // 财帛
	Zinv     // 子女
	Fuqi     // 夫妻
	Xiongdi  // 兄弟

	Shengong // 身宫

	// 五行局
	Shui2ju // 水二局
	Mu3ju   // 木三局
	Jin4ju  // 金四局
	Tu5ju   // 土五局
	Huo6ju  // 火六局

	// 星亮度
	Miao   // 庙
	Wang   // 旺
	Dedi   // 得地
	Liyi   // 利益
	Pinghe // 平和
	Budedi // 不得地
	Xian   // 陷

	Ziwei // 紫微
	// # 紫微诸星
	Tianfu   // 天府
	Taiyang  // 太阳
	Wuqu     // 武曲
	Tianji   // 天机
	Tiantong // 天同
	Lianzhen // 廉贞
	// # 天府诸星
	Taiyin    // 太阴
	Tanlang   // 贪狼
	Jumeng    // 巨门
	Tianxiang // 天相
	Tianliang // 天梁
	Qisha     // 七杀
	Pojun     // 破军
	// # 时系诸星
	Wenchang // 文昌
	Wenqu    // 文曲
	Dijie    // 地劫
	Dikong   // 地空
	Taifu    // 台辅
	Fenggao  // 封诰

	HuoXing  // 火星
	LingXing // 铃星

	// 月系诸星
	Zuofu    // 左辅
	Youbi    // 右弼
	Tianxing // 天刑
	Tianyao  // 天姚
	Tianwu   // 天巫
	Tianyue  // 天月
	Yinsha   // 阴煞

	// 日系诸星
	Santai  // 三台
	Bazuo   // 八座
	Enguang // 恩光
	Tiangui // 天贵

	// 干系诸星表
	Lucun        // 禄存
	Qingyang     // 擎羊
	Tuoluo       // 陀罗
	Tiankui      // 天魁
	TianyueKejia // 天钺
	Tianguan     // 天官
	Tianfu4      // 天福
	Tianchu      // 天厨

	// 四化星
	HuaKe   // 化科
	HuaQuan // 化权
	HuaLu   // 化禄
	HuaJi   // 化忌

	// 支系诸星
	Tianma   // 天马
	Hongluan // 红鸾
	Tianxi   // 天喜
	Tianku   // 天哭
	Tianxu   // 天虚
	Guchen   // 孤辰
	Guasu    // 寡宿
	Jiesheng // 解神
	Longchi  // 龙池
	Fengge   // 凤阁
	Feilian  // 蜚廉
	Posui    // 破碎
	Tiankong // 天空
	Yuede    // 月德

	Tiancai  // 天才
	Tianshou // 天寿

	Jiekong // 截空
	Xunkong // 旬空

	Tianshang // 天伤
	Tianshi   // 天使

	Mingzhu // 命主
	Shenzhu // 身主

	// 五局长生十二星
	Changsheng // 长生
	Muyu       // 沐浴
	Guandai    // 冠带
	Linguan    // 临官
	Diwang     // 帝旺
	ShuaiXing  // 衰
	BingXing   // 病
	SiXing     // 死
	MuXing     // 墓
	JueXing    // 绝
	TaiXing    // 胎
	YangXing   // 养

	// 博士十二星
	Boshi        // 博士
	Lishi        // 力士
	Qinglong     // 青龙
	Xiaohao      // 小耗
	Jiangjun     // 将军
	Zoushu       // 奏书
	FeilianBoshi // 飞廉
	Xishen       // 喜神
	Bingfu       // 病符
	Dahao        // 大耗
	Fubing       // 伏兵
	Guanfu       // 官府

	// 流年将前十二星
	Jiangxing // 将星
	Panan     // 攀鞍
	Suiyi     // 岁驿
	Xi1shen   // 息神
	Huagai    // (华盖)
	Jiesha    // (劫煞)
	Zaisha    // 灾煞
	Tiansha   // 天煞
	Zhibei    // 指背
	Xianchi   // (咸池)
	Yuexi     // 月煞
	Wangshen  // 亡神

	// 流年岁前十二星
	Suijian        // 岁建
	Huiqi          // 晦气
	Sangmen        // (丧门)
	GuanSuo        // 贯索
	GuanFu         // (官符)
	XiaohaoSuiqian // (小耗)
	DahaoSuiqian   // (大耗)
	Longde         // 龙德
	Baihu          // (白虎)
	Tiande         // (天德)
	Diaoke         // (吊客)
	BingfuSuiqian  // (病符)

	Zidou // 子斗

	End
)
