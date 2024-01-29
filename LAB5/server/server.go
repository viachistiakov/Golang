package main

import (
	"encoding/json"
	"flag"
	utils "lab5"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8023", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("Последовательность действительных чисел: %s", message)

		var seq utils.Sequence
		err = json.Unmarshal(message, &seq)
		if err != nil {
			log.Fatalln(err)
		}
		var res utils.Result
		var max_res float64 = math.MinInt64
		var min_res float64 = math.MaxFloat64
		for i := 0; i < len(seq.Data); i++ {
			if seq.Data[i] > max_res {
				max_res = seq.Data[i]
			}
			if seq.Data[i] < min_res {
				min_res = seq.Data[i]
			}
		}
		res.MaxNum = max_res
		res.MinNum = min_res

		message, err = json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
