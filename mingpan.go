package doushu

import (
	"bytes"
	"fmt"
)

type MingPan struct {
	MingZhu *MingZhu

	Gongs []*Gong

	MingGong *Gong `json:"-"`
	ShenGong *Gong `json:"-"`

	Positions []Element `json:"-"` // 星所在宫的地支代码
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

	gongs := make([]*Gong, 12)
	for i := Zi; i <= Hai; i++ {
		gongs[i-Zi] = &Gong{
			Dizhi: i,
		}
	}

	positions := make([]Element, End)
	for i := 0; i < int(End); i++ {
		positions[i] = Element(-1)
	}

	return &MingPan{
		MingZhu:   mingZhu,
		Gongs:     gongs,
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
}

func (m *MingZhu) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s %s%s年 %s%s %s时 %s%s %s"`, m.Name,
		m.NianGan, m.NianZhi, m.Yue, m.Ri, m.Shi, m.Yinyang, m.Gender, m.Wuxingju,
	)), nil
}

type Gong struct {
	Dizhi   Element
	Tiangan Element
	Gong    Element

	Stars  []Element
	Lights []Element
}

func (m *Gong) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	for i, star := range m.Stars {
		if i == 0 {
			continue
		}
		buf.WriteString(star.String())
		buf.WriteString(" ")
	}
	return []byte(fmt.Sprintf(`"%s%s %s: %s"`,
		m.Tiangan, m.Dizhi, m.Gong, buf.String(),
	)), nil
}
