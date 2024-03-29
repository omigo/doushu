package doushu

func ValueOf(name string) Element {
	for i, v := range names {
		if v == name {
			return Element(i)
		}
	}
	// log.Printf("%s not found", name)
	return Unknown
}

var names = []string{
	"", // Unknown
	//  两仪
	"阳", // Yang
	"阴", // Yin2
	// 2
	"男", // Nan
	"女", // Nv
	// 4
	"木", // Mu
	"火", // Huo
	"土", // Tu
	"金", // Jin
	"水", // Shui
	// 9
	"鼠", // ShuXiao
	"牛", // NiuXiao
	"虎", // HuXiao
	"兔", // TuXiao
	"龙", // LongXiao
	"蛇", // SheXiao
	"马", // MaXiao
	"羊", // YangXiao
	"猴", // HouXiao
	"鸡", // JiXiao
	"狗", // GouXiao
	"猪", // ZhuXiao
	// 21 十天干
	"甲", // Jia
	"乙", // Yi
	"丙", // Bing
	"丁", // Ding
	"戊", // WuGan
	"己", // Ji
	"庚", // Geng
	"辛", // Xin
	"壬", // Ren
	"癸", // Gui
	// 31 十二地支
	"子", // Zi
	"丑", // Chou
	"寅", // Yin
	"卯", // Mao
	"辰", // Chen
	"巳", // Si
	"午", // Wu
	"未", // Wei
	"申", // Shen
	"酉", // You
	"戌", // Xu
	"亥", // Hai
	// 43 农历月
	"正月", // Zhengyue
	"二月", // Eryue
	"三月", // Sanyue
	"四月", // Siyue
	"五月", // Wuyue
	"六月", // LiuYue
	"七月", // Qiyue
	"八月", // Bayue
	"九月", // Jiuyue
	"十月", // Shiyue
	"冬月", // Dongyue
	"腊月", // Layue
	// 55 农历日
	"初一",  // Chuyi
	"初二",  // Chuer
	"初三",  // Chusan
	"初四",  // Chusi
	"初五",  // Chuwu
	"初六",  // Chuliu
	"初七",  // Chuqi
	"初八",  // Chuba
	"初九",  // Chujiu
	"初十",  // Chushi
	"十一",  // Shiyi
	"十二",  // Shier
	"十三",  // Shisan
	"十四",  // Shisi
	"十五",  // Shiwu
	"十六",  // Shiliu
	"十七",  // Shiqi
	"十八",  // Shiba
	"十九",  // Shijiu
	"二十",  // Ershi
	"二十一", // Ershiyi
	"二十二", // Ershier
	"二十三", // Ershisan
	"二十四", // Ershisi
	"二十五", // Ershiwu
	"二十六", // Ershiliu
	"二十七", // Ershiqi
	"二十八", // Ershiba
	"二十九", // Ershijiu
	"三十",  // Sanshi

	// 身宫和命等十二宫
	"命宫", // MingGong
	"父母", // Fumu
	"福德", // Fude
	"田宅", // Tian
	"官禄", // Guanlu
	"交友", // Jiaoyou
	"迁移", // Qianyi
	"疾厄", // Jie
	"财帛", // Caibo
	"子女", // Zinv
	"夫妻", // Fuqi
	"兄弟", // Xiongdi

	"身宫", // ShenGong

	// 五行局
	"水二局", // Shui2Ju
	"木三局", // Mu3Ju
	"金四局", // Jin4Ju
	"土五局", // Tu5Ju
	"火六局", // Huo6Ju

	"庙",   // Miao
	"旺",   // Wang
	"得地",  // Dedi
	"利益",  // Liyi
	"平和",  // Pinghe
	"不得地", // Budedi
	"陷",   // Xian

	"紫微", // Ziwei
	"天府", // Tianfu
	"太阳", // Taiyang
	"武曲", // Wuqu
	"天机", // Tianji
	"天同", // Tiantong
	"廉贞", // Lianzhen
	// # 天府诸星
	"太阴", // Taiyin
	"贪狼", // Tanlang
	"巨门", // Jumeng
	"天相", // Tianxiang
	"天梁", // Tianliang
	"七杀", // Qisha
	"破军", // Pojun
	// # 时系诸星
	"文昌", // Wenchang
	"文曲", // Wenqu
	"地劫", // Dijie
	"地空", // Dikong
	"台辅", // Taifu
	"封诰", // Fenggao

	"火星", // HuoXing
	"铃星", // LingXing

	// 月系诸星
	"左辅", // Zuofu
	"右弼", // Youbi
	"天刑", // Tianxing
	"天姚", // Tianyao
	"天巫", // Tianwu
	"天月", // Tianyue
	"阴煞", // Yinsha

	// 日系诸星
	"三台", // Santai
	"八座", // Bazuo
	"恩光", // Enguang
	"天贵", // Tiangui

	// 干系诸星表
	"禄存", // Lucun
	"擎羊", // Qingyang
	"陀罗", // Tuoluo
	"天魁", // Tiankui
	"天钺", // Tianyue4
	"天官", // Tianguan
	"天福", // Tianfu4
	"天厨", // Tianchu

	// 四化星
	"化科", // HuaKe
	"化权", // HuaQuan
	"化禄", // HuaLu
	"化忌", // HuaJi

	// 支系诸星
	"天马", // Tianma
	"红鸾", // Hongluan
	"天喜", // Tianxi
	"天哭", // Tianku
	"天虚", // Tianxu
	"孤辰", // Guchen
	"寡宿", // Guasu
	"解神", // Jiesheng
	"龙池", // Longchi
	"凤阁", // Fengge
	"蜚廉", // Feilian
	"破碎", // Posui
	"天空", // Tiankong
	"月德", // Yuede

	"天才", // Tiancai
	"天寿", // Tianshou

	"截空", // Jiekong
	"旬空", // Xunkong

	"天伤", // Tianshang
	"天使", // Tianshi

	"命主", // Mingzhu
	"身主", // Shenzhu

	// 五行长生十二星
	"长生", // Changsheng
	"沐浴", // Muyu
	"冠带", // Guandai
	"临官", // Linguan
	"帝旺", // Diwang
	"衰",  // ShuaiXing
	"病",  // BingXing
	"死",  // SiXing
	"墓",  // MuXing
	"绝",  // JueXing
	"胎",  // TaiXing
	"养",  // YangXing

	// 博士十二星
	"博士", // Boshi
	"力士", // Lishi
	"青龙", // Qinglong
	"小耗", // Xiaohao
	"将军", // Jiangjun
	"奏书", // Zoushu
	"飞廉", // Feilian3
	"喜神", // Xishen
	"病符", // Bingfu
	"大耗", // Dahao
	"伏兵", // Fubing
	"官府", // Guanfu

	// 流年将前十二星
	"将星",   // Jiangxing
	"攀鞍",   // Panan
	"岁驿",   // Suiyi
	"息神",   // Xishen
	"(华盖)", // Huagai
	"(劫煞)", // Jiesha
	"灾煞",   // Zaisha
	"天煞",   // Tiansha
	"指背",   // Zhibei
	"(咸池)", // Xianchi
	"月煞",   // Yuexi
	"亡神",   // Wangshen

	// 流年岁前十二星
	"岁建",   // Suijian
	"晦气",   // Huiqi
	"(丧门)", // Sangmen
	"贯索",   // GuanSuo
	"(官符)", // GuanFu
	"小耗",   // Xiaohao
	"大耗",   // Dahao
	"龙德",   // Longde
	"(白虎)", // Baihu
	"(天德)", // Tiande
	"(吊客)", // Diaoke
	"病符",   // Bingfu

	"子斗", // Zidou
}
