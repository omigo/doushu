package doushu

import (
	"encoding/json"
	"testing"

)

func TestArrange(t *testing.T) {
	pan := Arrange( "周延发", Nan, Ji, Si, Sanyue, Shiliu, Hai)
	bs, _:=json.MarshalIndent( pan, "", "  ")
	t.Log( string( bs))
}
