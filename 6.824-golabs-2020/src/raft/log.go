package raft



type AppendEntriesArgs struct {
	term         int
	leaderId     int            // leader id
	prevLogIndex int            //紧邻新日志条目之前的那个日志条目的索引
	prevLogTerm  int            // 紧邻新日志条目之前的那个日志条目的任期
	entries      [] interface{} // 需要被保存的日志条目（被当做心跳使用是 则日志条目内容为空；为了提高效率可能一次性发送多个）
	leaderCommit int            //领导者的已知已提交的最高的日志条目的索引
}

type AppendEntriesApply struct {
	term    int
	success bool // 如果跟随者所含有的条目和preLogIndex以及preLogTerm匹配上了,则为true
}


