package doushu

import (
	"encoding/json"
	"testing"
)

func TestArrange(t *testing.T) {
	pan := Arrange("民国31年生", Nv, Ren, Wu, Shiyue, Eryi, Shen)
	// pan := Arrange("民国38年生", Nv, Ji, Chou, Zhengyue, Erba, Si)
	bs, _ := json.MarshalIndent(pan, "", "  ")
	t.Log(string(bs))
}
