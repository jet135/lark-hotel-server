package consts

var (
	BusinessFlagYes int8 = 1
	BusinessFlagNo  int8 = 0
)

type RoomType int8

const (
	None RoomType = iota
	SingleBed
	DoubleBed
	JuniorSuite
	SmallMahjongSuite
	LargeSuite
)

// 用于判断房间号属于什么房型
var (
	SingleBedArray         = [...]string{"401"}
	DoubleBedArray         = [...]string{"402"}
	JuniorSuiteArray       = [...]string{"403"}
	SmallMahjongSuiteArray = [...]string{"405"}
	LargeSuiteArray        = [...]string{"406", "407"}
)
