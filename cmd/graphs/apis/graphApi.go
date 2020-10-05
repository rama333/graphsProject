package apis

import (
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"graphsProject/cmd/graphs/daos"
	"graphsProject/cmd/graphs/models"
	"graphsProject/cmd/graphs/services"
)

func GraphApi() (*socketio.Server, error) {

	service := services.NewServiceDAO(daos.NewGraphDAO())

	server, _ := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("gr")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/graph", "data", func(s socketio.Conn, data string) string {
		s.SetContext(data)

		fmt.Println(data)

		//if data == "" {
		//	fmt.Println("data nil ok", data)
		//	response, short,  _ := service.Save(nil)
		//
		//	d, _ := json.Marshal(response)
		//
		//	fmt.Println("API", string(d))
		//
		//	server.BroadcastToRoom("", "gr", "event:name", string(d))
		//
		//	return string(d)
		//}

		var graphs []models.Graph
		err := json.Unmarshal([]byte(data), &graphs)

		if err != nil {
			fmt.Println(err)
		}

		response, short, err := service.Save(graphs)

		d, err := json.Marshal(response)
		sh, _ := json.Marshal(short)
		fmt.Println("API", string(d))

		server.BroadcastToRoom("", "gr", "event:graphs", string(d))

		server.BroadcastToRoom("", "gr", "event:short", string(sh))

		return string(d)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server, nil
}
