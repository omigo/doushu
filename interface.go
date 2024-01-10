package doushu

import (
	"context"
	"fmt"
)

func GetTianganYinyang(gan element) element {
	return p.get("十天干所属表", DizhiElement(0), gan)
}

func GetDizhiYingyang(zhi element) element {
	return p.get("十二地支所属表", DizhiElement(0), zhi)
}

func GetYinShou(niangan element) element {
	return p.get("起寅首天干表", DizhiElement(0), niangan)
}

// 安命宫表 按本生月及本生时,凡闰月生人，作下月论
func GetMingGong(ctx context.Context, nongliYue, nongliShi element) element {
	return p.get("安命宫表", nongliYue, nongliShi)
}

func GetShenGong(nongliYue, nongliShi element) element {
	return p.get("安身宫表", nongliYue, nongliShi)
}

func GetOtherGong(ctx context.Context, mingZhi element) ([]element, element) {
	return p.batchGet(ctx, "定十二宫表", mingZhi), Fumu
}

func Get12GongsTiangan(ctx context.Context, niangan element) []element {
	return p.batchGet(ctx, "定十二宫天干表", niangan)
}

func GetWuxingju(ctx context.Context, niangan, mingZhi element) element {
	return p.get("定五行局表", niangan, mingZhi)
}

func GetZiwei(wuxinju, ri element) element {
	return p.get("起紫微表", wuxinju, ri)
}

func GetZiweiStars(ctx context.Context, ziweiZhi element) ([]element, element) {
	return p.batchGet(ctx, "安紫微诸星表", ziweiZhi), Tianji
}

func GetTianfuStars(ctx context.Context, tianfu element) ([]element, element) {
	return p.batchGet(ctx, "安天府诸星表", tianfu), Taiyin
}

func GetShiStars(ctx context.Context, shi element) ([]element, element) {
	return p.batchGet(ctx, "安时系诸星表", shi), Wenchang
}

func GetHuoStar(nianZhi, shi element) element {
	fmt.Println(nianZhi, shi, p.get("安火星表", nianZhi, shi))
	return p.get("安火星表", nianZhi, shi)
}

func GetLingStar(nianZhi, shi element) element {
	return p.get("安铃星表", nianZhi, shi)
}

func GetYueStars(ctx context.Context, yue element) ([]element, element) {
	return p.batchGet(ctx, "安月系诸星表", yue), Zuofu
}

func GetRiStars(ctx context.Context, ri, zuofu, youbi, wenchang, wenqu element) ([]element, element) {
	return []element{
		DizhiElement(zuofu.Add(ri) % 12),          // 三 台 从左辅上起初一，顺行，数到本日生。
		DizhiElement((youbi.Sub(ri) + 36) % 12),   // 八 座 从右弼上起初一，逆行，数到本日生。
		DizhiElement((wenchang.Add(ri) - 1) % 12), // 恩 光 从文昌上起初一，顺行，数到本日生再退后一步
		DizhiElement((wenqu.Add(ri) - 1) % 12),    // 天 贵 从文曲上起初一，顺行，数到本日生再退后一步
	}, Santai
}

func GetGanStars(ctx context.Context, niangan element) ([]element, element) {
	return p.batchGet(ctx, "安干系诸星表", niangan), Lucun
}

func Get4HuaStars(ctx context.Context, niangan element) ([]element, element) {
	return p.batchGet(ctx, "安四化诸星表", niangan), HuaKe
}

func GetBoshi12Stars(lucun element) ([]element, element) {
	panic("not implemented")
}

func GetZhiStars(ctx context.Context, nianzhi element) ([]element, element) {
	return p.batchGet(ctx, "安支系诸星表", nianzhi), Tianma
}

func GetTiancai(nianzhi element) element {
	return p.get("安天才表", DizhiElement(0), nianzhi)
}

func GetTianshou(nianzhi element) element {
	panic("not implemented")
}

func GetChangsheng12Stars(ctx context.Context, wuxinju, yiyang, gender element) ([]element, element) {
	top := wuxinju.Value()

	if (yiyang == Yang && gender == Nan) ||
		(yiyang == Yin && gender == Nv) {
		top = top * 2
	} else {
		top = top*2 + 1
	}
	return p.batchGet(ctx, "安五行长生十二星表", DizhiElement(top)), Changsheng
}

func GetJiekong(niangan element) element {
	return p.get("安截路空亡表(截空)", DizhiElement(0), niangan)
}

func GetXunkong(niangan, nianzhi element) element {
	return p.get("安旬中空亡表(旬空)", nianzhi, niangan)
}

func GetShangShi(ctx context.Context, mingzhi element) ([]element, element) {
	return p.batchGet(ctx, "安天伤、天使表", mingzhi), Tianshang
}

func GetMingzhu(mingzhi element) element {
	return p.get("安命主表", DizhiElement(0), mingzhi)
}

func GetShenZhu(nianzhi element) element {
	return p.get("安身主表", DizhiElement(0), nianzhi)
}

func GetLiuNianJiangQianStars(ctx context.Context, nianzhi element) ([]element, element) {
	return p.batchGet(ctx, "安流年将前诸星表", nianzhi), Jiangxing
}

func GetLiuNianSuiQianStars(ctx context.Context, nianzhi element) ([]element, element) {
	return p.batchGet(ctx, "安流年岁前诸星表", nianzhi), Suijian
}

func GetZiNianDouJunPositions(shi, yue element) element {
	return p.get("安子年斗君表", shi, yue)
}
