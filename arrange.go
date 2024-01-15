package doushu

import (
	"time"

	"github.com/liujiawm/gocalendar"
)

func build(pan *MingPan) *MingPan {
	// 地支
	pan.Gongs = make([]*Gong, 12)
	for i := Zi; i <= Hai; i++ {
		pan.Gongs[i.Value()] = &Gong{Dizhi: i}
	}
	// 天干
	for i := Jia; i <= Gui; i++ {
		pos := pan.Positions[i]
		pan.Gongs[pos.Value()].Tiangan = i
	}
	pan.Gongs[Zi.Value()].Tiangan = pan.Gongs[Yin.Value()].Tiangan
	pan.Gongs[Chou.Value()].Tiangan = pan.Gongs[Mao.Value()].Tiangan
	// 身、命十二宫
	pan.Shengong = pan.Gongs[pan.Positions[Shengong].Value()]
	pan.Shengong.Shen = true
	pan.Minggong = pan.Gongs[pan.Positions[Minggong].Value()]
	for i := Minggong; i <= Xiongdi; i++ {
		pos := pan.Positions[i]
		pan.Gongs[pos.Value()].Gong = i
	}
	// 甲乙丙星入宫
	sihuaStars, first := GetSihuaStars(pan.MingZhu.NianGan)
	for star := Ziwei; star < Mingzhu; star++ {
		if star == HuaKe || star == HuaQuan || star == HuaLu || star == HuaJi {
			continue
		}

		pos := pan.Positions[star]
		s := &Star{
			Element: star,
			Level:   GetStarLevel(star),
			Light:   GetStarLight(star, pos),
		}
		for i, huaStar := range sihuaStars {
			if star == huaStar {
				s.Hua = first.Next(i)
			}
		}
		gong := pan.Gongs[pos.Value()]

		switch s.Level {
		case Jia:
			gong.JiaStars = append(gong.JiaStars, s)
		case Yi:
			gong.YiStars = append(gong.YiStars, s)
		case Bing:
			gong.BingStars = append(gong.BingStars, s)
		default:
		}
	}

	// 长生十二星入宫
	for star := Changsheng; star <= YangXing; star++ {
		pos := pan.Positions[star]
		gong := pan.Gongs[pos.Value()]
		gong.Changsheng12Star = star
	}
	// 博士十二星入宫
	for star := Boshi; star <= Guanfu; star++ {
		pos := pan.Positions[star]
		gong := pan.Gongs[pos.Value()]
		gong.Boshi12Star = star
	}
	// 流年将前十二星入宫
	for star := Jiangxing; star <= Wangshen; star++ {
		pos := pan.Positions[star]
		gong := pan.Gongs[pos.Value()]
		gong.Jianqian12Star = star
	}
	// 流年岁前十二星入宫
	for star := Suijian; star <= BingfuSuiqian; star++ {
		pos := pan.Positions[star]
		gong := pan.Gongs[pos.Value()]
		gong.Suiqian12Star = star
	}

	// 大限
	starts := GetDaxianStarts(pan.Positions[Minggong], pan.MingZhu.Wuxingju, pan.MingZhu.Yinyang, pan.MingZhu.Gender)
	for i, start := range starts {
		pan.Gongs[i].DaxianStart = start
	}
	// 小限
	firsts := GetXiaoxianFirsts(pan.MingZhu.NianZhi, pan.MingZhu.Gender)
	for i, first := range firsts {
		pan.Gongs[i].Xiaoxian = first
	}

	pan.MingZhu.Mingzhu = pan.Positions[Mingzhu]
	pan.MingZhu.Shenzhu = pan.Positions[Shenzhu]
	pan.MingZhu.Zidou = pan.Positions[Zidou]

	return pan
}

var c = gocalendar.NewCalendar(gocalendar.CalendarConfig{
	NightZiHour: true,
})

func ToNongli(timeOfBirth time.Time) (niangan, nianzhi, yue, ri, shi Element) {
	ld := c.GregorianToLunar(timeOfBirth.Year(), int(timeOfBirth.Month()), timeOfBirth.Day())
	niangan, nianzhi = Jia.Next(ld.YearGZ.HSI), Zi.Next(ld.YearGZ.EBI)
	yue, ri = Zhengyue.Next(ld.Month-1), Chuyi.Next(ld.Day-1)
	ct := c.ChineseSexagenaryCycle(timeOfBirth)
	shi = Zi.Next(ct.Hour.EBI)
	return niangan, nianzhi, yue, ri, shi
}

func Arrange(name string, gender, niangan, nianzhi, yue, ri, shi Element) *MingPan {
	// 初始化索引
	pan := NewMingPan(name, gender, niangan, nianzhi, yue, ri, shi)

	setTiangan(pan)        // 冠盖天支
	setMingShen12Gong(pan) // 定命等十二宫和身宫
	setWuxingju(pan)       // 定命宫五行局
	setYinYang(pan)        // 定男女阴阳

	setZiwei(pan)                // 起紫微
	setZiweiPositions(pan)       // 安紫微诸星表
	setTianfuPositions(pan)      // 安天府诸星表
	setShiPositions(pan)         // 安时系诸星表
	setHuoXingPosition(pan)      // 安火星表
	setLingXingPosition(pan)     // 安铃星表
	setYuePositions(pan)         // 安月系诸星表
	setRiPositions(pan)          // 安日系诸星表
	setGanPositions(pan)         // 安干系诸星表
	setSihuaPositions(pan)       // 安四化星表
	setBoshi12StarPositions(pan) // 安生年博士十二星法
	setZhiPositions(pan)         // 安支系诸星表
	setChangshengPositions(pan)  // 安五局长生十二星表
	setTiancaiPosition(pan)      // 安天才星表
	setTianshouPosition(pan)     // 安天寿星表
	setJiekong(pan)              // 安截路空亡表(截空)
	setXunkong(pan)              // 安旬中空亡表(旬空)
	setTianShangTianshi(pan)     // 安天伤、天使表
	setMingZhu(pan)              // 安命主表
	setShenZhu(pan)              // 安身主表
	setJiangqianPositions(pan)   // 安流年将前诸星表
	setSuiqianPositions(pan)     // 安流年岁前诸星表
	setZidouPosition(pan)        // 安子年斗君表

	return build(pan)
}

func setTiangan(pan *MingPan) {
	shouTiangan := GetYinShou(pan.MingZhu.NianGan)
	for i, gan := range shouTiangan.NextTo(10) {
		pan.Positions[gan] = Yin.Next(i)
	}
}

func setMingShen12Gong(pan *MingPan) {
	mingZhi := GetMinggong(pan.MingZhu.Yue, pan.MingZhu.Shi)
	poses, first := Get12Gongs(mingZhi)
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}

	shenZhi := GetShenGong(pan.MingZhu.Yue, pan.MingZhu.Shi)
	pan.Positions[Shengong] = shenZhi
}

func setWuxingju(pan *MingPan) {
	pan.MingZhu.Wuxingju = GetWuxingju(pan.MingZhu.NianGan, pan.Positions[Minggong])
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
	pan.Positions[Tianshou] = GetTianshou(pan.Positions[Shengong], pan.MingZhu.NianZhi)
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
	poses, first := GetShangShi(pan.Positions[Minggong])
	for i, zhi := range poses {
		pan.Positions[first.Next(i)] = zhi
	}
}

func setMingZhu(pan *MingPan) {
	pan.MingZhu.Mingzhu = GetMingzhu(pan.Positions[Minggong])
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
