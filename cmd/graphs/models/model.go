package models

type Graph struct {
	Nm  int   `json:"nm"`
	Way []int `json:"go"`
	X   int   `json:"x"`
	Y   int   `json:"y"`
	Chk int   `json:"chk"`
}

type GraphEdge struct {
	A    int `json:"a"`
	B    int `json:"b"`
	Edge int `json:"edge"`
}
