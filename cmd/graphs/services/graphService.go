package services

import (
	"fmt"
	"graphsProject/cmd/graphs/models"
)

type serviceDAO interface {
	Save([]models.Graph) ([]models.Graph, error)
}

type service struct {
	dao serviceDAO
}

func NewServiceDAO(dao serviceDAO) *service {

	return &service{dao}
}

func (s service) Save(g []models.Graph) ([]models.Graph, []int, error) {

	graph, err := s.dao.Save(g)

	if err != nil {
		return nil, nil, err
	}
	short := shortСut(graph)

	return graph, short, nil
}

func shortСut(g []models.Graph) []int {

	inf := 1 << 31
	v := 0
	t := 3

	graphs := []models.GraphEdge{}

	for _, obj := range g {
		for _, b := range obj.Way {
			graphs = append(graphs, models.GraphEdge{A: obj.Nm, B: b, Edge: 7})
			graphs = append(graphs, models.GraphEdge{A: b, B: obj.Nm, Edge: 7})
		}
	}

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	fmt.Println(g)
	fmt.Println(graphs)

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	//graphs = append(graphs, models.GraphEdge{A: 1, b: 2, edge: -5})
	//graphs = append(graphs, models.GraphEdge{A: 0, b: 2, edge: 6})
	//graphs = append(graphs, models.GraphEdge{A: 2, b: 3, edge: 4})
	//graphs = append(graphs, models.GraphEdge{A: 3, b: 1, edge: 8})

	d := []int{}
	p := []int{}
	m := len(graphs)
	n := len(graphs)

	for i := 0; i < n-1; i++ {
		d = append(d, inf)
		p = append(p, -1)
	}

	d[v] = 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			if d[graphs[j].A] < inf {

				if d[graphs[j].B] > d[graphs[j].A]+graphs[j].Edge {

					d[graphs[j].B] = min(d[graphs[j].B], d[graphs[j].A]+graphs[j].Edge)
					p[graphs[j].B] = graphs[j].A
					//fmt.Println(i, j, d[graphs[j].b], d[graphs[j].a]+graphs[j].edge)
				}
			}
		}
	}

	fmt.Println(d) //return [0 7 2 6]
	fmt.Println(p)

	way := []int{}

	for i := t; i != -1; i-- {
		if p[i] == -1 {
			break
		}
		way = append(way, p[i])

	}

	fmt.Println(way)

	return way
}

func min(a int, b int) int {

	if a < b {
		return a
	}
	return b

}
