package daos

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"graphsProject/cmd/graphs/models"
	"log"
)

type graphDAO struct{}

func NewGraphDAO() *graphDAO {

	return &graphDAO{}
}

func (g *graphDAO) Save(graps []models.Graph) ([]models.Graph, error) {

	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	grap, err := json.Marshal(graps)

	if err != nil {
		log.Fatal(err)
	}

	if graps != nil {
		if _, err := conn.Do("SET", "graph", grap); err != nil {
			log.Fatal(err)
		}
	}

	values, err := redis.String(conn.Do("GET", "graph"))
	if err != nil {
		fmt.Println("value rr")
		log.Fatal(err)

	}

	var grapsNew []models.Graph

	if err := json.Unmarshal([]byte(values), &grapsNew); err != nil {
		log.Fatal(err)
	}

	//if err := redis.ScanStruct(values, grapsNew); err != nil{
	//	log.Fatal(err)
	//}

	return grapsNew, nil
}
