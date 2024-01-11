package doushu

import (
	"testing"
)

func TestGetMingGong(t *testing.T) {
	if GetMingGong(Zhengyue, Zi) != Yin {
		t.Error("GetMingGong error")
	}
}

func TestGetWuxingju(t *testing.T) {
	if GetWuxingju(Jia, Zi) != Shui2Ju {
		t.Error("GetMingGong error")
	}
}

func TestGetZiwei(t *testing.T) {
	if GetZiwei(Mu3Ju, Chuyi) != Chen {
		t.Error("GetMingGong error")
	}
}

func TestGetZiweiStars(t *testing.T) {
	poses, star := GetZiweiStars(You)
	t.Log(ToName(star), ToName(poses[0]))
	if poses[0] != Shen {
		t.Fatal("GetZiweiStars error")
	}
}

func TestGetHuoStar(t *testing.T) {
	if GetHuoStar(Si, Hai) != Yin {
		t.Fatal(ToName(GetHuoStar(Si, Hai)))
	}
}
