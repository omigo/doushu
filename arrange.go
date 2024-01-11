package doushu

func Arrange(name string, gender, niangan, nianzhi, yue, ri, shi Element) *MingPan {
	// 十二宫，每宫一个地支
	pan := NewMingPan(name, gender, niangan, nianzhi, yue, ri, shi)

	// 配干支
	setTiangan(pan)
	// // 定命身二宫
	setMingShen12Gong(pan)
	// 定五行局
	setWuxingju(pan)
	// 定男女阴阳
	setYinYang(pan)

	// 安星时，先建议索引，即找到星所在宫的地支
	// 最后统一把星放到宫中
	setZiwei(pan)                     // 起紫微
	setZiweiPositions(pan)            // // 安紫微诸星表
	setTianfuPositions(pan)           // // 安天府诸星表
	setShiPositions(pan)              // // 安时系诸星表
	setHuoXingPosition(pan)           // // 安火星表
	setLingXingPosition(pan)          // //安铃星表
	setYuePositions(pan)              // // 安月系诸星表
	setRiPositions(pan)               // // 安日系诸星表
	setGanPositions(pan)              // // 安干系诸星表
	setZhiPositions(pan)              // // 安支系诸星表
	setChangshengPositions(pan)       // // 安五行长生十二星表
	setJiekong(pan)                   // // 安截路空亡表(截空)
	setXunkong(pan)                   // // 安旬中空亡表(旬空)
	setTianShangTianshi(pan)          // // 安天伤、天使表
	setMingZhu(pan)                   // // 安命主表
	setShenZhu(pan)                   // // 安身主表
	setLiuNianJiangQianPositions(pan) // // 安流年将前诸星表
	setLiuNianSuiQianPositions(pan)   // // 安流年岁前诸星表
	// SetZiNianDouJunPositions( pan)     // // 安子年斗君表

	for i := 0; i < 12; i++ {
		poses := make([]Element, 0, 16)
		gong := Element(i + int(Zi))
		for star, zhi := range pan.Positions {
			if zhi == -1 {
				continue
			}
			if zhi == gong {
				poses = append(poses, Element(star))
			}
		}
		pan.Gongs[i].Stars = poses
		// pan.Gongs[].Lights =
	}
	return pan
}

func setTiangan(pan *MingPan) {
	shouTiangan := GetYinShou(pan.MingZhu.NianGan)

	for i := Zi; i <= Hai; i++ {
		pan.Gongs[(Yin+i)%12].Tiangan = (shouTiangan + i) % 12
	}
}

func setMingShen12Gong(pan *MingPan) {
	mingZhi := GetMingGong(pan.MingZhu.Yue, pan.MingZhu.Shi)
	shenZhi := GetShenGong(pan.MingZhu.Yue, pan.MingZhu.Shi)

	for gong := MingGong; gong <= Xiongdi; gong++ {
		pan.Gongs[(mingZhi-Zi+gong)%12].Gong = gong
	}
	pan.Gongs[shenZhi-Zi].Gong = ShenGong

	pan.MingGong = pan.Gongs[mingZhi-Zi]
	pan.ShenGong = pan.Gongs[shenZhi-Zi]
}

func setWuxingju(pan *MingPan) {
	pan.MingZhu.Wuxingju = GetWuxingju(pan.MingZhu.NianGan, pan.MingGong.Dizhi)
}

func setYinYang(pan *MingPan) {
	pan.MingZhu.Yinyang = GetTianganYinyang(pan.MingZhu.NianGan)
}

func setZiwei(pan *MingPan) {
	zhi := GetZiwei(pan.MingZhu.Wuxingju, pan.MingZhu.Ri)
	pan.Positions[Ziwei] = zhi
}

func setZiweiPositions(pan *MingPan) {
	poses, first := GetZiweiStars(pan.Positions[Ziwei])
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

// 安天府诸星表
func setTianfuPositions(pan *MingPan) {
	poses, first := GetTianfuStars(pan.Positions[Tianfu])
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

// 安时系诸星表
func setShiPositions(pan *MingPan) {
	poses, first := GetShiStars(pan.MingZhu.Shi)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setHuoXingPosition(pan *MingPan) {
	pan.Positions[HuoXing] = GetHuoStar(pan.MingZhu.NianZhi, pan.MingZhu.Shi)
}
func setLingXingPosition(pan *MingPan) {
	pan.Positions[LingXing] = GetLingStar(pan.MingZhu.NianZhi, pan.MingZhu.Shi)
}

// 安月系诸星表
func setYuePositions(pan *MingPan) {
	poses, first := GetYueStars(pan.MingZhu.Yue)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}
func setRiPositions(pan *MingPan) {
	poses, first := GetRiStars(pan.MingZhu.Ri,
		pan.Positions[Zuofu], pan.Positions[Youbi],
		pan.Positions[Wenchang], pan.Positions[Wenqu],
	)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setGanPositions(pan *MingPan) {
	poses, first := GetGanStars(pan.MingZhu.NianGan)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setZhiPositions(pan *MingPan) {
	poses, first := GetZhiStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setChangshengPositions(pan *MingPan) {
	poses, first := GetChangsheng12Stars(pan.MingZhu.Wuxingju, pan.MingZhu.Yinyang, pan.MingZhu.Gender)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setJiekong(pan *MingPan) {
	pan.Positions[Jiekong] = GetJiekong(pan.MingZhu.NianGan)
}
func setXunkong(pan *MingPan) {
	pan.Positions[Xunkong] = GetXunkong(pan.MingZhu.NianGan, pan.MingZhu.NianZhi)
}

func setTianShangTianshi(pan *MingPan) {
	poses, first := GetShangShi(pan.MingGong.Dizhi)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}
func setMingZhu(pan *MingPan) {
	pan.Positions[Mingzhu] = GetMingzhu(pan.MingGong.Dizhi)
}
func setShenZhu(pan *MingPan) {
	pan.Positions[Shenzhu] = GetShenZhu(pan.MingZhu.NianZhi)
}

func setLiuNianJiangQianPositions(pan *MingPan) {
	poses, first := GetLiuNianJiangQianStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}

func setLiuNianSuiQianPositions(pan *MingPan) {
	poses, first := GetLiuNianSuiQianStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[int(first)+i] = zhi
	}
}
