package consensus

import (
	"time"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

/*
RequestVote from the RAFT spec:
Receiver implementation:
1. Reply false if term < currentTerm (§5.1)
2. If votedFor is null or candidateId, and candidate’s log is at
least as up-to-date as receiver’s log, grant vote (§5.2, §5.4)
*/
func (s *server) RequestVote(ctx context.Context, in *RequestVoteReq) (*RequestVoteRes, error) {
	log.Infof("RequestVote: %+v", in)
	mutex.Lock()
	currentTerm := s.state.currentTerm
	s.leaderAliveTs = time.Now()
	mutex.Unlock()
	if in.CurrentTerm < currentTerm {
		log.Infof("RequestVote:in.CurrentTerm:%d < currentTerm:%d", in.CurrentTerm, currentTerm)
		return &RequestVoteRes{
			VoteGranted: false,
			CurrentTerm: currentTerm,
		}, nil
	} else {
		mutex.Lock()
		s.state.currentTerm = in.CurrentTerm
		s.state.votedFor = ""
		if s.state.state == candidate {
			s.state.state = follower
		}

		mutex.Unlock()
	}

	mutex.Lock()
	defer mutex.Unlock()
	if s.state.votedFor == "" || s.state.votedFor == in.CandidateId {
		s.state.votedFor = in.CandidateId
		log.Infof("REQUEST VOTE service: %s voted for: %s", s.id, in.CandidateId)
		return &RequestVoteRes{
			VoteGranted: true,
			CurrentTerm: in.CurrentTerm,
		}, nil

	}
	return &RequestVoteRes{
		VoteGranted: false,
		CurrentTerm: s.state.currentTerm,
	}, nil
}

/*
AppendEntries from the RAFT spec:
Receiver implementation:
1. Reply false if term < currentTerm (§5.1)
2. Reply false if log doesn’t contain an entry at prevLogIndex
whose term matches prevLogTerm (§5.3)
3. If an existing entry conflicts with a new one (same index
but different terms), delete the existing entry and all that
follow it (§5.3)
4. Append any new entries not already in the log
5. If leaderCommit > commitIndex, set commitIndex =
min(leaderCommit, index of last new entry)
*/
func (s *server) AppendEntries(ctx context.Context, in *AppendEntriesReq) (*AppendEntriesRes, error) {
	log.Infof("AppendEntries: %+v", in)
	mutex.Lock()
	currentTerm := s.state.currentTerm
	s.leaderAliveTs = time.Now()
	mutex.Unlock()

	appendResult := &AppendEntriesRes{
		CurrentTerm: currentTerm,
		Success:     false,
	}
	if in.CurrentTerm < currentTerm {
		return appendResult, nil
	}

	mutex.Lock()
	l := int64(len(s.state.log)) - 1
	mutex.Unlock()

	if l < in.PrevLogIndex {
		return appendResult, nil
	} else if l == in.PrevLogIndex && currentTerm != in.PrevLogTerm {
		mutex.Lock()
		s.state.log = s.state.log[:l-1]
		mutex.Unlock()
	}
	mutex.Lock()
	s.state.log = append(s.state.log, in.Entries...)

	l = int64(len(s.state.log)) - 1
	// If leaderCommit > commitIndex, set commitIndex =
	// min(leaderCommit, index of last new entry)
	if in.LeaderCommit > s.state.commitIndex {
		var min int64
		if in.LeaderCommit < l {
			min = in.LeaderCommit
		} else {
			min = l
		}
		s.state.commitIndex = min
	}
	if s.state.commitIndex > s.state.lastApplied {
		s.state.lastApplied = l
	}

	mutex.Lock()
	return &AppendEntriesRes{}, nil
}
