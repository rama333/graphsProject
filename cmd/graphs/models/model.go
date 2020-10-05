package models

type Graph struct {
	Nm        int   `json:"nm"`
	Way       []int `json:"go"`
	X         int   `json:"x"`
	Y         int   `json:"y"`
	Chk       int   `json:"chk"`
	EdgeStart []int `json:"edge_start"`
	EdgeEnd   []int `json:"edge_end"`
}

type GraphEdge struct {
	A    int `json:"a"`
	B    int `json:"b"`
	Edge int `json:"edge"`
}
