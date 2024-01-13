package doushu

import (
	"encoding/json"
	"testing"
)

func TestArrange(t *testing.T) {
	// pan := Arrange("民国31年生", Nv, Ren, Wu, Shiyue, Eryi, Shen)
	// pan := Arrange("民国38年生", Nv, Ji, Chou, Zhengyue, Erba, Si)
	pan := Arrange("匿名", Nv, Gui, Mao, Shier, Chusan, Si)
	bs, _ := json.MarshalIndent(pan, "", "  ")
	t.Log(string(bs))
}
