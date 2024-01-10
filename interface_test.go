package doushu

import (
	"testing"

	"github.com/omigo/g"
)

func TestGetMingGong(t *testing.T) {
	ctx := g.WithTraceId()
	if GetMingGong(ctx, Zhengyue, Zi) != Yin {
		t.Error("GetMingGong error")
	}
}

func TestGetWuxingju(t *testing.T) {
	ctx := g.WithTraceId()
	if GetWuxingju(ctx, Jia, Zi) != Shui2Ju {
		t.Error("GetMingGong error")
	}
}

func TestGetOtherGong(t *testing.T) {
	t.Log(ValueOf("").Value())

	ctx := g.WithTraceId()
	poses, _ := GetOtherGong(ctx, Zi)
	if poses[0] != Chou {
		t.Error("GetOtherGong error")
	}
}
func TestGetZiwei(t *testing.T) {
	if GetZiwei(Mu3Ju, Chuyi) != Chen {
		t.Error("GetMingGong error")
	}
}
