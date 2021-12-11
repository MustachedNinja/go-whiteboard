package handler

import (
	"fmt"
	"log"
	"net/http"
	"whiteboard/state"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MarshalProto(protoObj proto.Message, asJson bool) ([]byte, error) {
	if asJson {
		jsonBytes, err := protojson.MarshalOptions{}.Marshal(protoObj)
		return jsonBytes, err
	} else {
		return proto.Marshal(protoObj)
	}
}

func WriteResponse(w http.ResponseWriter, data []byte, asJson bool) {
	if asJson {
		fmt.Fprint(w, string(data))
	} else {
		fmt.Fprint(w, data)
	}
}

func GetUpdates(w http.ResponseWriter, req *http.Request) {
	update := state.StateUpdate{
		Version: "abra",
		Data1:   1234,
	}

	out, err := MarshalProto(&update, true)
	if err != nil {
		log.Fatalf("Marshaling error: ", err)
	}
	WriteResponse(w, out, true)
}
