package replicaSetStatus

type ReplicaSetStatus struct {
	Set     string              "set"
	Members []*ReplicaSetMember "members"
}

func (status *ReplicaSetStatus) GetPrimary() *ReplicaSetMember {
	for i := 0; i < len(status.Members); i++ {
		member := status.Members[i]
		if member.StateStr == "PRIMARY" {
			return member
		}
	}

	return nil
}

func (status *ReplicaSetStatus) PopulateSyncingFrom() {
	syncMap := make(map[string]*ReplicaSetMember)

	// Build map of name to ReplicaSetMember
	for i := 0; i < len(status.Members); i++ {
		member := status.Members[i]
		syncMap[member.Name] = member
	}

	// Populate SyncingFrom for each ReplicaSetMember
	for _, member := range syncMap {
		syncingTo, ok := syncMap[member.SyncingTo]
		if ok {
			syncingTo.AddSyncingFrom(member)
		}
	}
}

type ReplicaSetMember struct {
	Name        string "name"
	StateStr    string "stateStr"
	SyncingTo   string "syncingTo"
	SyncingFrom []*ReplicaSetMember
}

func (member *ReplicaSetMember) AddSyncingFrom(syncingFrom *ReplicaSetMember) {
	if member.SyncingFrom == nil {
		member.SyncingFrom = []*ReplicaSetMember{syncingFrom}
	} else {
		member.SyncingFrom = append(member.SyncingFrom, syncingFrom)
	}
}
