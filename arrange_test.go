package doushu

import (
	"testing"

	"github.com/omigo/g"
)

func TestArrange(t *testing.T) {
	ctx := g.WithTraceId()
	ctx = g.WithLevel(ctx, g.Ltrace)
	pan := Arrange(ctx, "周延发", Nan, Ji, Si, Sanyue, Shiliu, Hai)
	g.Debug(g.WithTraceId(), pan)
}
