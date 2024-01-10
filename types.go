package doushu

type element interface {
	Value() int
	String() string
	Add(element) int
	Sub(element) int
}

type UnknownElement int
type YingyangElement int
type NannvElement int
type WuxingElement int
type ShenxiaoElement int
type TianganElement int
type DizhiElement int
type NongliYueElement int
type NongliRiElement int
type GongElement int
type WuxingjuElement int
type StarElement int

func (a UnknownElement) String() string    { return "unknown" }
func (a UnknownElement) Value() int        { return int(a) }
func (a UnknownElement) Add(e element) int { return a.Value() + e.Value() }
func (a UnknownElement) Sub(e element) int { return a.Value() - e.Value() }

func (a YingyangElement) String() string    { return toString(a) }
func (a YingyangElement) Value() int        { return int(a) }
func (a YingyangElement) Add(e element) int { return a.Value() + e.Value() }
func (a YingyangElement) Sub(e element) int { return a.Value() - e.Value() }

func (a NannvElement) String() string    { return toString(a) }
func (a NannvElement) Value() int        { return int(a) }
func (a NannvElement) Add(e element) int { return a.Value() + e.Value() }
func (a NannvElement) Sub(e element) int { return a.Value() - e.Value() }

func (a WuxingElement) String() string    { return toString(a) }
func (a WuxingElement) Value() int        { return int(a) }
func (a WuxingElement) Add(e element) int { return a.Value() + e.Value() }
func (a WuxingElement) Sub(e element) int { return a.Value() - e.Value() }

func (a ShenxiaoElement) String() string    { return toString(a) }
func (a ShenxiaoElement) Value() int        { return int(a) }
func (a ShenxiaoElement) Add(e element) int { return a.Value() + e.Value() }
func (a ShenxiaoElement) Sub(e element) int { return a.Value() - e.Value() }

func (a TianganElement) String() string    { return toString(a) }
func (a TianganElement) Value() int        { return int(a) }
func (a TianganElement) Add(e element) int { return a.Value() + e.Value() }
func (a TianganElement) Sub(e element) int { return a.Value() - e.Value() }

func (a DizhiElement) String() string    { return toString(a) }
func (a DizhiElement) Value() int        { return int(a) }
func (a DizhiElement) Add(e element) int { return a.Value() + e.Value() }
func (a DizhiElement) Sub(e element) int { return a.Value() - e.Value() }

func (a NongliYueElement) String() string    { return toString(a) }
func (a NongliYueElement) Value() int        { return int(a) }
func (a NongliYueElement) Add(e element) int { return a.Value() + e.Value() }
func (a NongliYueElement) Sub(e element) int { return a.Value() - e.Value() }

func (a NongliRiElement) String() string    { return toString(a) }
func (a NongliRiElement) Value() int        { return int(a) }
func (a NongliRiElement) Add(e element) int { return a.Value() + e.Value() }
func (a NongliRiElement) Sub(e element) int { return a.Value() - e.Value() }

func (a GongElement) String() string    { return toString(a) }
func (a GongElement) Value() int        { return int(a) }
func (a GongElement) Add(e element) int { return a.Value() + e.Value() }
func (a GongElement) Sub(e element) int { return a.Value() - e.Value() }

func (a WuxingjuElement) String() string    { return toString(a) }
func (a WuxingjuElement) Value() int        { return int(a) }
func (a WuxingjuElement) Add(e element) int { return a.Value() + e.Value() }
func (a WuxingjuElement) Sub(e element) int { return a.Value() - e.Value() }

func (a StarElement) String() string    { return toString(a) }
func (a StarElement) Value() int        { return int(a) }
func (a StarElement) Add(e element) int { return a.Value() + e.Value() }
func (a StarElement) Sub(e element) int { return a.Value() - e.Value() }
