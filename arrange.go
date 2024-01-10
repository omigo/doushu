package doushu

import (
	"context"
	"fmt"
)

func Arrange(ctx context.Context, name string, gender, niangan, nianzhi, yue, ri, shi element) *MingPan {
	// 十二宫，每宫一个地支
	pan := NewMingPan(name, gender, niangan, nianzhi, yue, ri, shi)

	// 配干支
	setTiangan(pan, niangan)
	// // 定命身二宫
	setMingShen12Gong(ctx, pan, yue, shi)
	// 定五行局
	setWuxingju(ctx, pan)
	// 定男女阴阳
	setYinYang(pan)

	// 安星时，先建议索引，即找到星所在宫的地支
	// 最后统一把星放到宫中
	setZiwei(pan)                          // 起紫微
	setZiweiPositions(ctx, pan)            // // 安紫微诸星表
	setTianfuPositions(ctx, pan)           // // 安天府诸星表
	setShiPositions(ctx, pan)              // // 安时系诸星表
	setHuoXingPosition(ctx, pan)           // // 安火星表
	setLingXingPosition(ctx, pan)          // //安铃星表
	setMonthPositions(ctx, pan)            // // 安月系诸星表
	setDayPositions(ctx, pan)              // // 安日系诸星表
	setGanPositions(ctx, pan)              // // 安干系诸星表
	setZhiPositions(ctx, pan)              // // 安支系诸星表
	setChangshengPositions(ctx, pan)       // // 安五行长生十二星表
	setJiekong(pan)                        // // 安截路空亡表(截空)
	setXunkong(pan)                        // // 安旬中空亡表(旬空)
	setTianShangTianshi(ctx, pan)          // // 安天伤、天使表
	setMingZhu(pan)                        // // 安命主表
	setShenZhu(pan)                        // // 安身主表
	setLiuNianJiangQianPositions(ctx, pan) // // 安流年将前诸星表
	setLiuNianSuiQianPositions(ctx, pan)   // // 安流年岁前诸星表
	// SetZiNianDouJunPositions(ctx, pan)     // // 安子年斗君表

	for i := Zi; i <= Hai; i++ {
		poses := make([]StarElement, 0, 16)
		for star, zhi := range pan.Positions {
			if zhi.Value() == -1 {
				continue
			}
			if zhi == i {
				poses = append(poses, StarElement(star))
			}
		}
		pan.Gongs[i].Stars = poses
		// pan.Gongs[].Lights =
	}
	return pan
}

func setTiangan(pan *MingPan, minggan element) {
	shouTiangan := GetYinShou(minggan)

	for i := Zi; i <= Hai; i++ {
		pan.Gongs[Yin.Add(i)%12].Tiangan = TianganElement(shouTiangan.Add(i) % 12)
	}
}

func setMingShen12Gong(ctx context.Context, pan *MingPan, yue, shi element) {
	ming := GetMingGong(ctx, yue, shi)
	shen := GetShenGong(yue, shi)

	pan.MingGong = pan.Gongs[ming.Value()]
	pan.ShenGong = pan.Gongs[shen.Value()]

	// TODO 不需要查表
	others, first := GetOtherGong(ctx, ming)
	for gong, pos := range others {
		pan.Gongs[pos.Add(first)%12].Gong = GongElement(gong + 1)
	}
}

func setWuxingju(ctx context.Context, pan *MingPan) {
	pan.MingZhu.Wuxingju = GetWuxingju(ctx, pan.MingZhu.NianGan, pan.MingGong.Dizhi)
}

func setYinYang(pan *MingPan) {
	pan.MingZhu.Yinyang = GetTianganYinyang(pan.MingZhu.NianGan)
}

func setZiwei(pan *MingPan) {
	ziweiZhi := GetZiwei(pan.MingZhu.Wuxingju, pan.MingZhu.Ri)
	fmt.Println(pan.MingZhu.Wuxingju.Value(), pan.MingZhu.Ri.Value(), ziweiZhi.Value())
	pan.Positions[Ziwei] = ziweiZhi
}

func setZiweiPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetZiweiStars(ctx, pan.Positions[Ziwei])
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

// 安天府诸星表
func setTianfuPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetTianfuStars(ctx, pan.Positions[Tianfu])
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

// 安时系诸星表
func setShiPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetShiStars(ctx, pan.MingZhu.Shi)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setHuoXingPosition(ctx context.Context, pan *MingPan) {
	pan.Positions[HuoXing] = GetHuoStar(pan.MingZhu.NianZhi, pan.MingZhu.Shi)
}
func setLingXingPosition(ctx context.Context, pan *MingPan) {
	pan.Positions[LingXing] = GetLingStar(pan.MingZhu.NianZhi, pan.MingZhu.Shi)
}

// 安月系诸星表
func setMonthPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetYueStars(ctx, pan.MingZhu.Yue)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}
func setDayPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetRiStars(ctx, pan.MingZhu.Ri,
		pan.Positions[Zuofu], pan.Positions[Youbi],
		pan.Positions[Wenchang], pan.Positions[Wenqu],
	)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setGanPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetGanStars(ctx, pan.MingZhu.NianGan)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setZhiPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetZhiStars(ctx, pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setChangshengPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetChangsheng12Stars(ctx, pan.MingZhu.Wuxingju, pan.MingZhu.Yinyang, pan.MingZhu.Gender)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setJiekong(pan *MingPan) {
	pan.Positions[Jiekong] = GetJiekong(pan.MingZhu.NianGan)
}
func setXunkong(pan *MingPan) {
	pan.Positions[Xunkong] = GetXunkong(pan.MingZhu.NianGan, pan.MingZhu.NianZhi)
}

func setTianShangTianshi(ctx context.Context, pan *MingPan) {
	poses, first := GetShangShi(ctx, pan.MingGong.Dizhi)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}
func setMingZhu(pan *MingPan) {
	pan.Positions[Mingzhu] = GetMingzhu(pan.MingGong.Dizhi)
}
func setShenZhu(pan *MingPan) {
	pan.Positions[Shenzhu] = GetShenZhu(pan.MingZhu.NianZhi)
}

func setLiuNianJiangQianPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetLiuNianJiangQianStars(ctx, pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}

func setLiuNianSuiQianPositions(ctx context.Context, pan *MingPan) {
	poses, first := GetLiuNianSuiQianStars(ctx, pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Value()+i] = zhi
	}
}
