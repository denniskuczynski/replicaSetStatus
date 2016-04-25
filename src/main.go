package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	rs "replicaSetStatus"
)

func main() {
	session, err := mgo.Dial("localhost:3000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	result := rs.ReplicaSetStatus{}
	if err := session.DB("admin").Run("replSetGetStatus", &result); err != nil {
		panic(err)
	} else {
		result.PopulateSyncingFrom()
		primary := result.GetPrimary()
		if primary == nil {
			fmt.Println("No Primary")
		}
		fmt.Printf("Primary = %v\n", primary)
	}
}
