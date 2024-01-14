package doushu

func GetTianganYinyang(gan Element) Element {
	return p.get("十天干所属表", 0, gan.Value())
}

func GetDizhiYingyang(zhi Element) Element {
	return p.get("十二地支所属表", 0, zhi.Value())
}

func GetYinShou(niangan Element) Element {
	return p.get("起寅首天干表", 0, niangan.Value())
}

// 安命宫表 按本生月及本生时,凡闰月生人，作下月论
func GetMinggong(nongliYue Element, shi Element) Element {
	return p.get("安命宫表", nongliYue.Value(), shi.Value())
}

func GetShenGong(nongliYue Element, shi Element) Element {
	return p.get("安身宫表", nongliYue.Value(), shi.Value())
}

func Get12Gongs(mingZhi Element) ([]Element, Element) {
	es := make([]Element, 12)
	for i := 0; i < 12; i++ {
		es[i] = mingZhi.Next(i)
	}
	return es, Minggong
}

func GetWuxingju(niangan Element, mingZhi Element) Element {
	return p.get("定五行局表", niangan.Value(), mingZhi.Value())
}

func GetZiwei(wuxinju Element, ri Element) Element {
	return p.get("起紫微表", wuxinju.Value(), ri.Value())
}

func GetZiweiStars(ziweiZhi Element) ([]Element, Element) {
	return p.batchGet("安紫微诸星表", ziweiZhi.Value()), Tianfu
}

func GetTianfuStars(tianfuZhi Element) ([]Element, Element) {
	return p.batchGet("安天府诸星表", tianfuZhi.Value()), Taiyin
}

func GetShiStars(shi Element) ([]Element, Element) {
	return p.batchGet("安时系诸星表", shi.Value()), Wenchang
}

func GetHuoStar(nianZhi, shi Element) Element {
	return p.get("安火星表", nianZhi.Value(), shi.Value())
}

func GetLingStar(nianZhi, shi Element) Element {
	return p.get("安铃星表", nianZhi.Value(), shi.Value())
}

func GetYueStars(yue Element) ([]Element, Element) {
	return p.batchGet("安月系诸星表", yue.Value()), Zuofu
}

func GetRiStars(ri, zuofuPos, youbiPos, wenchangPos, wenquPos Element) ([]Element, Element) {
	return []Element{
		zuofuPos.Next(ri.Value()),        // 三 台 从左辅上起初一，顺行，数到本日生
		youbiPos.Pre(ri.Value()),         // 八 座 从右弼上起初一，逆行，数到本日生
		wenchangPos.Next(ri.Value() - 1), // 恩 光 从文昌上起初一，顺行，数到本日生再退后一步
		wenquPos.Next(ri.Value() - 1),    // 天 贵 从文曲上起初一，顺行，数到本日生再退后一步
	}, Santai
}

func GetGanStars(niangan Element) ([]Element, Element) {
	return p.batchGet("安干系诸星表", niangan.Value()), Lucun
}

func GetSihuaStars(niangan Element) ([]Element, Element) {
	return p.batchGet("安四化诸星表", niangan.Value()), HuaKe
}

func GetBoshi12Stars(lucunPos Element) ([]Element, Element) {
	return lucunPos.NextTo(Hai.Value()), Boshi
}

func GetZhiStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安支系诸星表", nianzhi.Value()), Tianma
}

func GetTiancai(nianzhi, nongliYue, shi Element) Element {
	gong := p.get("安天才星表", 0, nianzhi.Value())
	mingZhi := GetMinggong(nongliYue, shi)
	poses, _ := Get12Gongs(mingZhi)
	return poses[gong]
}

// GetTiancai  通过本生年支，得到天才星所在十二宫，可能是妻财子禄。注意，返回的不是如子丑寅卯的地支。
func GetTiancaiGong(nianzhi Element) Element {
	return p.get("安天才星表", 0, nianzhi.Value())
}

func GetTianshou(shenggongZhi, nianZhi Element) Element {
	return shenggongZhi.Next(nianZhi.Value())
}

func GetChangsheng12Stars(wuxinju, yiyang, gender Element) ([]Element, Element) {
	top := wuxinju.Value()
	if isYangNanYinNv(yiyang, gender) {
		top = top * 2
	} else {
		top = top*2 + 1
	}
	return p.batchGet("安五局长生十二星表", top), Changsheng
}

func GetJiekong(niangan Element) Element {
	return p.get("安截路空亡表(截空)", 0, niangan.Value())
}

func GetXunkong(niangan, nianzhi Element) Element {
	return p.get("安旬中空亡表(旬空)", nianzhi.Value(), niangan.Value())
}

func GetShangShi(mingzhi Element) ([]Element, Element) {
	return p.batchGet("安天伤、天使表", mingzhi.Value()), Tianshang
}

func GetMingzhu(mingzhi Element) Element {
	return p.get("安命主表", 0, mingzhi.Value())
}

func GetShenZhu(nianzhi Element) Element {
	return p.get("安身主表", 0, nianzhi.Value())
}

func GetJiangqianStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安流年将前诸星表", nianzhi.Value()), Jiangxing
}

func GetSuiqianStars(nianzhi Element) ([]Element, Element) {
	return p.batchGet("安流年岁前诸星表", nianzhi.Value()), Suijian
}

func GetZidou(shi, yue Element) Element {
	return p.get("安子年斗君表", shi.Value(), yue.Value())
}

func GetStarLight(star, zhi Element) Element {
	if star < Ziwei || star > Mingzhu {
		return Unknown
	}
	if zhi < Zi || zhi > Hai {
		return Unknown
	}
	if star == Zuofu || star == Youbi || star == Tiankui || star == Tianyue || star == Lucun {
		return Miao
	}
	return p.get("诸星在十二宫庙旺利陷表", star.Value(), zhi.Value())
}

func GetStarLevel(star Element) Element {
	return p.get("诸星级行分化一览表", star.Value(), 0)
}

// GetDaxianStarts 通过命宫，阴阳，性别，得到大限起始年龄所在的地址。返回 [寅,子,丑,...]，表示命在寅宫，即第一个10年在寅，第二个10年在子。
func GetDaxians(mingZhi, yinyang, gender Element) []Element {
	xians := make([]Element, 12)
	if isYangNanYinNv(yinyang, gender) {
		for i := 0; i < 12; i++ {
			xians = append(xians, mingZhi.Next(i))
		}
	} else {
		for i := 0; i < 12; i++ {
			xians = append(xians, mingZhi.Pre(i))
		}
	}
	return xians
}

// GetDaxianStarts 通过命宫，五行局，阴阳，性别，得到大限起始年龄。返回 [23,13,3,..]，表示命在寅宫，子宫大限起始23岁，丑宫大限起始13岁。
func GetDaxianStarts(mingZhi, wuxinju Element, yinyang, gender Element) []int {
	starts := make([]int, 12)
	xians := GetDaxians(mingZhi, yinyang, gender)
	for i, zhi := range xians {
		starts[zhi.Value()] = i*10 + wuxinju.Value() + 2
	}
	return starts
}

func isYangNanYinNv(yinyang, gender Element) bool {
	return (yinyang == Yang && gender == Nan) || (yinyang == YinXing && gender == Nv)
}

func GetXiaoxianStart(nianzhi Element) Element {
	return Xu.Pre(3 * nianzhi.Value())
}

// GetXiaoxians
func GetXiaoxian(nianzhi, gender Element) []Element {
	start := GetXiaoxianStart(nianzhi)
	xians := make([]Element, 12)
	if gender == Nan {
		for i := 0; i < 12; i++ {
			xians = append(xians, start.Next(i))
		}
	} else {
		for i := 0; i < 12; i++ {
			xians = append(xians, start.Pre(i))
		}
	}

	return xians
}

// GetXiaoxianFirsts 通过本生年支、性别，得到小限起始年龄。返回 [3,4,5,..]，表示1岁在戌宫，3岁在子宫。
func GetXiaoxianFirsts(nianzhi, gender Element) []int {
	firsts := make([]int, 12)
	xians := GetXiaoxian(nianzhi, gender)
	for i, zhi := range xians {
		firsts[zhi.Value()] = i + 1
	}
	return firsts
}
