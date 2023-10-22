package demo

var (
	// Global ! 大写开头的外部可访问, 但尽量不要用
	Global   = "外部可访问"
	internal = "外部不可访问"
)

const (
	External         = "外部可访问"
	internalV1       = "外部不可访问"
	abc        uint8 = 123
)

const (
	Status0 = iota
	Status1 = iota
	Status2 = iota
	Abc     = iota
)

const (
	MyStatus0 = (iota + 1) << 10
	MyStatus1 = (iota + 1) << 10
)
