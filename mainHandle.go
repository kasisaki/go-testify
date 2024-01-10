package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("count missing"))
		if err != nil {
			log.Print("Error writing response")
		}
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("wrong count value"))
		if err != nil {
			log.Print("Error writing response")
		}
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("wrong city value"))
		if err != nil {
			log.Print("Error writing response")
		}
		return
	}

	if count > len(cafe) || count == 0 {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(answer))
	if err != nil {
		log.Print("Error writing response")
	}
}
