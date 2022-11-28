package types

type JobRef struct {
	Ref string
	Idx int
}

//node used in simple algorithms
type JobNode struct {
	Id int

	IOTime  float64

	Name string

	Inputs []JobRef
	InIds []int

	OutIds []int

	CpuTime float64
	Depth	int
}
