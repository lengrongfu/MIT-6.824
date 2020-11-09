package raft

// ConditionRun
func ConditionRun(rf *Raft) {
	DPrintf("peer %d is condition role,term is %d", rf.me, rf.GetCurrentTerm()+1)

	// term + 1
	rf.SetCurrentTerm(rf.GetCurrentTerm() + 1)

	// ask vote
	var voteC = make(chan *RequestVoteReply, len(rf.peers))
	askVote := func(peer int) {
		args := &RequestVoteArgs{
			term:        rf.GetCurrentTerm(),
			candidateId: rf.getMe(),
		}
		var reply *RequestVoteReply
		b := rf.sendRequestVote(peer, args, reply)
		if !b {
			reply.voteGranted = false
		}
		voteC <- reply
	}
	for index := range rf.peers {
		if index == rf.getMe() {
			voteSelf(rf, voteC)
			continue
		}
		go askVote(index)
	}

	hopeVoteNum := needVoteNum(rf)
	infactVoteNum := 0
	for rf.GetRole() == Condition {
		select {
		case reply := <-voteC:
			if reply.term > rf.GetCurrentTerm() {
				rf.SetRole(Follower)
				return
			}
			if reply.voteGranted {
				infactVoteNum++
			}
			if infactVoteNum > hopeVoteNum {
				rf.SetRole(Leader)
			}
		}
	}

}

func voteSelf(rf *Raft, voteC chan *RequestVoteReply) {
	voteC <- &RequestVoteReply{
		term:        rf.GetCurrentTerm(),
		voteGranted: true,
	}
}

func needVoteNum(rf *Raft) int {
	return len(rf.peers)/2 + 1
}
