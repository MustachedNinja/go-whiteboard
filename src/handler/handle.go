package handler

import (
	"fmt"
	"log"
	"net/http"
	"whiteboard/state"

	"google.golang.org/protobuf/proto"
)

func GetUpdates(w http.ResponseWriter, req *http.Request) {
	update := state.StateUpdate{
		Version: "abra",
		Data1:   1234,
	}
	out, err := proto.Marshal(&update)
	if err != nil {
		log.Fatalf("Marshaling error: ", err)
	}
	fmt.Fprint(w, out)
}
