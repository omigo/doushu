package doushu

import "fmt"

func GetTianganYinyang(gan Element) Element {
	return p.get("十天干所属表", 0, gan-Jia)
}

func GetDizhiYingyang(zhi Element) Element {
	return p.get("十二地支所属表", 0, zhi-Zi)
}

func GetYinShou(niangan Element) Element {
	return p.get("起寅首天干表", 0, niangan-Jia)
}

// 安命宫表 按本生月及本生时,凡闰月生人，作下月论
func GetMingGong(nongliYue Element, shi Element) Element {
	return p.get("安命宫表", nongliYue-Zhengyue, shi-Zi)
}

func GetShenGong(nongliYue Element, shi Element) Element {
	return p.get("安身宫表", nongliYue-Zhengyue, shi-Zi)
}

func Get12GongsTiangan(niangan Element) []Element {
	return p.batchGet("定十二宫天干表", niangan-Jia)
}

func GetWuxingju(niangan Element, mingZhi Element) Element {
	return p.get("定五行局表", niangan-Jia, mingZhi-Zi)
}

func GetZiwei(wuxinju Element, ri Element) Element {
	return p.get("起紫微表", wuxinju-Shui2Ju, ri-Chuyi)
}

func GetZiweiStars(ziweiZhi Element) ([]Element, Element) {
	return p.batchGet("安紫微诸星表", ziweiZhi-Zi), Tianji
}

func GetTianfuStars(tianfuZhi Element) ([]Element, Element) {
	return p.batchGet("安天府诸星表", tianfuZhi-Zi), Taiyin
}

func GetShiStars(shi Element) ([]Element, Element) {
	return p.batchGet("安时系诸星表", shi-Zi), Wenchang
}

func GetHuoStar(nianZhi, shi Element) Element {
	fmt.Println(ToName(nianZhi), ToName(shi))
	return p.get("安火星表", nianZhi-Zi, shi-Zi)
}

func GetLingStar(nianZhi, shi Element) Element {
	return p.get("安铃星表", nianZhi-Zi, shi-Zi)
}

func GetYueStars(yue Element) ([]Element, Element) {
	return p.batchGet("安月系诸星表", yue-Zhengyue), Zuofu
}

func GetRiStars(ri, zuofu, youbi, wenchang, wenqu Element) ([]Element, Element) {
	fmt.Println(ToName(ri), ToName(zuofu), ToName(youbi), ToName(wenchang), ToName(wenqu))
	return []Element{
		(zuofu-Zi+ri-Chuyi)%12 + Zi,      // 三 台 从左辅上起初一，顺行，数到本日生。
		(youbi-Zi-ri-Chuyi+36)%12 + Zi,   // 八 座 从右弼上起初一，逆行，数到本日生。
		(wenchang-Zi+ri-Chuyi-1)%12 + Zi, // 恩 光 从文昌上起初一，顺行，数到本日生再退后一步
		(wenqu-Zi+ri-Chuyi-1)%12 + Zi,    // 天 贵 从文曲上起初一，顺行，数到本日生再退后一步
	}, Santai
}

func GetGanStars(niangan Element) ([]Element, Element) {
	return p.batchGet("安干系诸星表", niangan-Jia), Lucun
}

func Get4HuaStars(niangan Element) ([]Element, Element) {
	return p.batchGet("安四化诸星表", niangan-Jia), HuaKe
}

func GetBoshi12Stars(lucun Element) ([]Element, Element) {
	panic("not implemented")
}

func GetZhiStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安支系诸星表", nianzhi-Zi), Tianma
}

func GetTiancai(nianzhi Element) Element {
	return p.get("安天才表", 0, nianzhi-Zi)
}

func GetTianshou(nianzhi Element) Element {
	panic("not implemented")
}

func GetChangsheng12Stars(wuxinju, yiyang, gender Element) ([]Element, Element) {
	top := wuxinju - Shui2Ju
	if (yiyang == Yang && gender == Nan) ||
		(yiyang == Yin && gender == Nv) {
		top = top * 2
	} else {
		top = top*2 + 1
	}
	return p.batchGet("安五行长生十二星表", top), Changsheng
}

func GetJiekong(niangan Element) Element {
	return p.get("安截路空亡表(截空)", 0, niangan-Jia)
}

func GetXunkong(niangan, nianzhi Element) Element {
	return p.get("安旬中空亡表(旬空)", nianzhi-Zi, niangan-Jia)
}

func GetShangShi(mingzhi Element) ([]Element, Element) {
	return p.batchGet("安天伤、天使表", mingzhi-Zi), Tianshang
}

func GetMingzhu(mingzhi Element) Element {
	return p.get("安命主表", 0, mingzhi-Zi)
}

func GetShenZhu(nianzhi Element) Element {
	return p.get("安身主表", 0, nianzhi-Zi)
}

func GetLiuNianJiangQianStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安流年将前诸星表", nianzhi-Zi), Jiangxing
}

func GetLiuNianSuiQianStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安流年岁前诸星表", nianzhi-Zi), Suijian
}

func GetZiNianDouJunPositions(shi, yue Element) Element {
	return p.get("安子年斗君表", shi-Zi, yue-Zhengyue)
}
