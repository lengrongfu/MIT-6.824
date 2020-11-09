package raft

const (
	Leader    Role = 1
	Condition Role = 2
	Follower  Role = 3
)

type Role int

func (r Role) String() string {
	switch r {
	case Leader:
		return "Leader"
	case Follower:
		return "Follower"
	case Condition:
		return "Condition"
	default:
		return "Unknown"
	}
}
