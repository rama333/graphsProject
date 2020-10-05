package services

import (
	"fmt"
	"graphsProject/cmd/graphs/models"
	"math/rand"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	graph, err := s.dao.Save(g)

	if err != nil {
		return nil, nil, err
	}

	for i, obj := range graph {
		if len(obj.Way) > 0 {
			for _, _ = range obj.Way {
				graph[i].EdgeStart = append(graph[i].EdgeStart, rand.Intn(30))
				rand.Seed(time.Now().UnixNano())
				graph[i].EdgeEnd = append(graph[i].EdgeEnd, rand.Intn(30))
			}
		} else {
			fmt.Println(rand.Intn(30))
			graph[i].EdgeStart = append(graph[i].EdgeStart, 0)
		}
	}

	graph[0].EdgeStart[0] = -5

	short := shortСut(graph)

	return graph, short, nil
}

func shortСut(g []models.Graph) []int {

	inf := 1 << 31
	v := 0

	m := 0
	n := 0

	graphs := []models.GraphEdge{}

	for _, obj := range g {
		n++
		if len(obj.Way) > 0 {
			for i, b := range obj.Way {
				graphs = append(graphs, models.GraphEdge{A: obj.Nm, B: b, Edge: obj.EdgeStart[i]})
				m++
				//graphs = append(graphs, models.GraphEdge{A: b, B: obj.Nm, Edge: obj.EdgeEnd[i]})
			}
		} else {
			//graphs = append(graphs, models.GraphEdge{A: obj.Nm, B: obj.Nm, Edge: obj.EdgeStart[0]})
			//m++
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

	d := []int{}
	p := []int{}

	t := n - 1

	for i := 0; i < n; i++ {
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

	fmt.Println(d)
	fmt.Println(p)

	way := []int{}

	for i := t; i != -1; i-- {
		if p[i] == -1 {
			break
		}
		way = append(way, p[i])
	}

	if len(way) == 0 {
		way = append(way, 0)
		return way
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
