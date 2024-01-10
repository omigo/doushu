package doushu

import "bytes"

type MingPan struct {
	MingZhu *MingZhu

	Gongs []*Gong

	MingGong *Gong
	ShenGong *Gong

	Positions []element // 星所在宫的地支代码
}

func NewMingPan(name string, gender, niangan, nianzhi, yue, ri, shi element) *MingPan {
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
		gongs[i] = &Gong{
			Dizhi: i,
		}
	}

	positions := make([]element, End)
	for i := 0; i < int(End); i++ {
		positions[i] = UnknownElement(-1)
	}

	return &MingPan{
		MingZhu:   mingZhu,
		Gongs:     gongs,
		Positions: positions,
	}
}

type MingZhu struct {
	Name   string
	Gender element

	NianGan element
	NianZhi element
	Yue     element
	Ri      element
	Shi     element

	Yinyang  element
	Wuxingju element
}

type Gong struct {
	Dizhi   DizhiElement
	Tiangan TianganElement
	Gong    GongElement

	Stars  Stars
	Lights []element
}

type Stars []StarElement

func (es Stars) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, e := range es {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(`"` + e.String() + `"`)
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}
