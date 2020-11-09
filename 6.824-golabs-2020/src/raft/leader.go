package raft

import "time"

func (rf Raft) leaderRun() {

}

// leaderInit 刚成为leader之后要做的事情
func (rf Raft) leaderInit() {
	// 给所有节点发送心跳信息
	go rf.heatBeatRun()
}

// 接受心跳的RPC
func (rf *Raft) RequestHeatBeat(args *AppendEntriesArgs, reply *AppendEntriesApply) {
	reply.success = true
	reply.term = rf.currentTerm
}

func (rf *Raft) sendRequestHeatBeat(server int, args *AppendEntriesArgs, reply *AppendEntriesApply) bool {
	ok := rf.peers[server].Call("Raft.RequestHeatBeat", args, reply)
	return ok
}

// heatBeatRun 按心跳间隔时长给 follow发送心跳信息
func (rf Raft) heatBeatRun() {
	heartbeatTimeout := time.After(rf.heatBeatTimeOut)
	// 成为leader 之后先发一轮s
	sendHeatBeat(rf)
	for ; ; {
		select {
		case <-heartbeatTimeout:
			sendHeatBeat(rf)
			heartbeatTimeout = time.After(rf.heatBeatTimeOut)
		}
	}
}

func sendHeatBeat(rf Raft) {
	for k := range rf.peers {
		if k == rf.me {
			continue
		}
		var args *AppendEntriesArgs = &AppendEntriesArgs{
			term:         rf.currentTerm,
			leaderId:     rf.me,
			prevLogIndex: 0,
			prevLogTerm:  0,
			entries:      nil,
			leaderCommit: 0,
		}
		var reply *AppendEntriesApply
		go rf.sendRequestHeatBeat(k, args, reply)
	}
}

func LeaderRun(rf *Raft) {
	go rf.heatBeatRun()
}
