package raft

import (
	"time"
)

const (
	// 最小心跳时间
	MinHeartbeatTimeout time.Duration = 100 * time.Millisecond
)

// FollowerRun
func FollowerRun(rf *Raft) {
	after := time.After(rf.heatBeatTimeOut)
	for rf.GetRole() == Follower {
		select {
		case <-after:
			rf.SetRole(Condition)
		}
	}
}
