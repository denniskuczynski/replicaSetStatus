package main

import (
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
	rs "replicaSetStatus"
)

func FromJSON(rsStatusJSON string) *js.Object {
	var rsStatus rs.ReplicaSetStatus
	err := json.Unmarshal([]byte(rsStatusJSON), &rsStatus)
	if err != nil {
		panic(err)
	}

	return js.MakeWrapper(&rsStatus)
}

func main() {
	js.Global.Set("rs", map[string]interface{}{
		"FromJSON": FromJSON,
	})
}
