package doushu

import (
	"testing"
)

func TestGetMingGong(t *testing.T) {
	if GetMinggong(Zhengyue, Zi) != Yin {
		t.Error("GetMingGong error")
	}
}

func TestGetWuxingju(t *testing.T) {
	if GetWuxingju(Jia, Zi) != Shui2ju {
		t.Error("GetMingGong error")
	}
}

func TestGetZiwei(t *testing.T) {
	if GetZiwei(Mu3ju, Chuyi) != Chen {
		t.Error("GetMingGong error")
	}
}

func TestGetZiweiStars(t *testing.T) {
	poses, star := GetZiweiStars(You)
	t.Log(star.String(), poses[0].String())
	if poses[0] != Shen {
		t.Fatal("GetZiweiStars error")
	}
}

func TestGetHuoStar(t *testing.T) {
	if GetHuoStar(Si, Hai) != Yin {
		t.Fatal(GetHuoStar(Si, Hai))
	}
}

func TestGetChangshengStars(t *testing.T) {
	poses, star := GetChangsheng12Stars(Jin4ju, Yang, Nv)
	t.Log(star.String(), poses[0].String())
	if poses[0] != Si {
		t.Fatal("GetZiweiStars error")
	}
}

func TestGetXiaoxianStart(t *testing.T) {
	if GetXiaoxianStart(Shen) != Xu {
		t.Fatal(GetXiaoxianStart(Shen))
	}
}
