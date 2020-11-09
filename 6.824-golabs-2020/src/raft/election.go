package raft

import (
	"time"
)

const (
	// 最小选举超时时间
	MinElectionTimeout = 150 * time.Millisecond
)
