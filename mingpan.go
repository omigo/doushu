package doushu

import (
	"bytes"
	"fmt"
	"strconv"
)

type MingPan struct {
	MingZhu *MingZhu

	Gongs []*Gong

	Minggong *Gong `json:"-"`
	Shengong *Gong `json:"-"`

	Positions []Element `json:"-"` // 星所在宫的地支代码
	Lights    []Element `json:"-"` //
}

func (pan *MingPan) String() string {
	var buf bytes.Buffer
	buf.WriteString(pan.MingZhu.String())
	for _, gong := range pan.Gongs {
		buf.WriteString("\n")
		buf.WriteString(gong.String())
	}
	return buf.String()
}

func NewMingPan(name string, gender, niangan, nianzhi, yue, ri, shi Element) *MingPan {
	mingZhu := &MingZhu{
		Name:   name,
		Gender: gender,

		NianGan: niangan,
		NianZhi: nianzhi,
		Yue:     yue,
		Ri:      ri,
		Shi:     shi,
	}

	positions := make([]Element, End)
	for i := 0; i < int(End); i++ {
		positions[i] = Element(-1)
	}

	return &MingPan{
		MingZhu:   mingZhu,
		Positions: positions,
	}
}

type MingZhu struct {
	Name   string
	Gender Element

	NianGan Element
	NianZhi Element
	Yue     Element
	Ri      Element
	Shi     Element

	Yinyang  Element
	Wuxingju Element

	Mingzhu Element
	Shenzhu Element
	Zidou   Element
}

func (m *MingZhu) YinNanYangNv() bool {
	return (m.Yinyang == Yang && m.Gender == Nan) || (m.Yinyang == YinXing && m.Gender == Nv)
}

func (m *MingZhu) String() string {
	return fmt.Sprintf("%s %s%s年 %s%s %s时 %s%s %s\n命主:%s 身主:%s 子斗:%s",
		m.Name, m.NianGan, m.NianZhi, m.Yue, m.Ri, m.Shi, m.Yinyang, m.Gender, m.Wuxingju,
		m.Mingzhu, m.Shenzhu, m.Zidou)
}

type Gong struct {
	Tiangan, Dizhi Element

	Gong Element
	Shen bool

	JiaStars, YiStars, BingStars []*Star
	HuaStars                     []Element

	Changsheng12Stars []Element
	Boshi12Stars      []Element
	Jianqian12Stars   []Element
	Suiqian12Stars    []Element

	DaxianStart int
	Xiaoxian    int
}

func (g *Gong) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	buf.WriteString(g.Tiangan.String())
	buf.WriteString(g.Dizhi.String())
	buf.WriteString(" ")
	buf.WriteString(g.Gong.String())
	if g.Shen {
		buf.WriteString(" ")
		buf.WriteString(Shengong.String())
	}

	appendStars(buf, "甲级星", g.JiaStars)
	appendStars(buf, "乙级星", g.YiStars)
	appendStars(buf, "丙级星", g.BingStars)
	appendStars2(buf, "长生十二星", g.Changsheng12Stars)
	appendStars2(buf, "博士十二星", g.Boshi12Stars)
	appendStars2(buf, "流年将前十二星", g.Jianqian12Stars)
	appendStars2(buf, "流年岁前十二星", g.Suiqian12Stars)

	fmt.Fprintf(buf, "\n  大限: %d-%d", g.DaxianStart, g.DaxianStart+9)

	buf.WriteString("\n  小限:")
	for i := 0; i < 6; i++ {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(g.Xiaoxian + 12*i))
	}

	return buf.String()
}

func appendStars(buf *bytes.Buffer, title string, stars []*Star) {
	buf.WriteString("\n  ")
	buf.WriteString(title)
	buf.WriteString(":")
	for _, star := range stars {
		buf.WriteString(" ")
		buf.WriteString(star.Element.String())
		if star.Light != 0 {
			buf.WriteString(star.Light.String())
		}
		if star.Hua != 0 {
			buf.WriteString(star.Hua.String())
		}
	}
}

func appendStars2(buf *bytes.Buffer, title string, stars []Element) {
	buf.WriteString("\n  ")
	buf.WriteString(title)
	buf.WriteString(":")
	for _, star := range stars {
		buf.WriteString(" ")
		buf.WriteString(star.String())
	}
}

type Star struct {
	Element Element
	Level   Element
	Light   Element
	Hua     Element
}
