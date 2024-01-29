package test

type Sequence struct {
	Data []float64 `json:"data"`
}

type Result struct {
	MaxNum float64 `json:"max_num"`
	MinNum float64 `json:"min_num"`
}
