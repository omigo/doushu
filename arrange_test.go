package doushu

import (
	"testing"
	"time"
)

func TestArrange(t *testing.T) {
	// pan := Arrange("民国31年生", Nv, Ren, Wu, Shiyue, Eryi, Shen)
	// pan := Arrange("民国38年生", Nv, Ji, Chou, Yiyue, Erba, Si)
	pan := Arrange("匿名", Nv, Gui, Mao, Shier, Chusan, Si)
	t.Log("\n" + pan.String())
}

func TestArrangeByModemTime(t *testing.T) {
	tob, err := time.ParseInLocation("2006-01-02 15:04", "2024-01-13 10:27", time.Local)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tob)

	niangan, nianzhi, yue, ri, shi := ToNongli(tob)
	pan := Arrange("匿名", Nv, niangan, nianzhi, yue, ri, shi)
	t.Log("\n" + pan.String())
}
