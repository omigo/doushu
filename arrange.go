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
	setZiwei(pan)                // 起紫微
	setZiweiPositions(pan)       // // 安紫微诸星表
	setTianfuPositions(pan)      // // 安天府诸星表
	setShiPositions(pan)         // // 安时系诸星表
	setHuoXingPosition(pan)      // // 安火星表
	setLingXingPosition(pan)     // //安铃星表
	setYuePositions(pan)         // // 安月系诸星表
	setRiPositions(pan)          // // 安日系诸星表
	setGanPositions(pan)         // // 安干系诸星表
	setSihuaPositions(pan)       // // 安四化星表
	setBoshi12StarPositions(pan) // // 安生年博士十二星法
	setZhiPositions(pan)         // // 安支系诸星表
	setChangshengPositions(pan)  // // 安五行长生十二星表
	setTiancaiPosition(pan)      // // 安天才星表
	setTianshouPosition(pan)     // // 安天寿星表
	setJiekong(pan)              // // 安截路空亡表(截空)
	setXunkong(pan)              // // 安旬中空亡表(旬空)
	setTianShangTianshi(pan)     // // 安天伤、天使表
	setMingZhu(pan)              // // 安命主表
	setShenZhu(pan)              // // 安身主表
	setJiangqianPositions(pan)   // // 安流年将前诸星表
	setSuiqianPositions(pan)     // // 安流年岁前诸星表
	setZidouPosition(pan)        // // 安子年斗君表

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
	for i, gan := range shouTiangan.NextTo(12) {
		pan.Gongs[i].Tiangan = gan
	}
}

func setMingShen12Gong(pan *MingPan) {
	mingZhi := GetMingGong(pan.MingZhu.Yue, pan.MingZhu.Shi)
	poses, first := Get12Gongs(mingZhi)
	for i, zhi := range poses {
		pan.Gongs[zhi.Value()].Gong = first.Next(i)
		pan.Positions[first.Next(i)] = zhi
	}
	pan.MingGong = pan.Gongs[mingZhi.Value()]

	shenZhi := GetShenGong(pan.MingZhu.Yue, pan.MingZhu.Shi)
	pan.Positions[Shengong] = shenZhi
	pan.ShenGong = pan.Gongs[shenZhi.Value()]
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
		pan.Positions[first.Next(i)] = zhi
	}
}

// 安天府诸星表
func setTianfuPositions(pan *MingPan) {
	poses, first := GetTianfuStars(pan.Positions[Tianfu])
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

// 安时系诸星表
func setShiPositions(pan *MingPan) {
	poses, first := GetShiStars(pan.MingZhu.Shi)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
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
		pan.Positions[first.Next(i)] = zhi
	}
}
func setRiPositions(pan *MingPan) {
	poses, first := GetRiStars(pan.MingZhu.Ri,
		pan.Positions[Zuofu], pan.Positions[Youbi],
		pan.Positions[Wenchang], pan.Positions[Wenqu],
	)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setGanPositions(pan *MingPan) {
	poses, first := GetGanStars(pan.MingZhu.NianGan)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setSihuaPositions(pan *MingPan) {
	poses, first := GetSihuaStars(pan.MingZhu.NianGan)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setBoshi12StarPositions(pan *MingPan) {
	lucunPos := pan.Positions[Lucun]
	var direct bool
	if (pan.MingZhu.Yinyang == Yang && pan.MingZhu.Gender == Nan) ||
		(pan.MingZhu.Yinyang == YinXing && pan.MingZhu.Gender == Nv) {
		direct = true
	}
	for i, star := range Boshi.NextTo(12) {
		if direct {
			pan.Positions[star] = lucunPos.Next(i)
		} else {
			pan.Positions[star] = lucunPos.Pre(i)
		}
	}
}

func setZhiPositions(pan *MingPan) {
	poses, first := GetZhiStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setTiancaiPosition(pan *MingPan) {
	gong := GetTiancaiGong(pan.MingZhu.NianZhi)
	pan.Positions[Tiancai] = pan.Positions[gong]
}

func setTianshouPosition(pan *MingPan) {
	pan.Positions[Tianshou] = GetTianshou(pan.ShenGong.Dizhi, pan.MingZhu.NianZhi)
}

func setChangshengPositions(pan *MingPan) {
	poses, first := GetChangsheng12Stars(pan.MingZhu.Wuxingju, pan.MingZhu.Yinyang, pan.MingZhu.Gender)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
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
		pan.Positions[first.Next(i)] = zhi
	}
}

func setMingZhu(pan *MingPan) {
	pan.MingZhu.Mingzhu = GetMingzhu(pan.MingGong.Dizhi)
	pan.Positions[Mingzhu] = pan.MingZhu.Mingzhu
}

func setShenZhu(pan *MingPan) {
	pan.MingZhu.Shenzhu = GetShenZhu(pan.MingZhu.NianZhi)
	pan.Positions[Shenzhu] = pan.MingZhu.Shenzhu
}

func setJiangqianPositions(pan *MingPan) {
	poses, first := GetJiangqianStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setSuiqianPositions(pan *MingPan) {
	poses, first := GetSuiqianStars(pan.MingZhu.NianZhi)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setZidouPosition(pan *MingPan) {
	pos := GetZidou(pan.MingZhu.Shi, pan.MingZhu.Yue)
	pan.Positions[Zidou] = pos
}
