package raft

import (
	"fmt"
	"sync"
	"time"
)

type LogEntry struct {
	Command any
	Term int
}

type CMState int 

const (
	Follower CMState = iota
	Candidate 
	Leader
	Dead
)

func (state CMState) String() string {
	switch state {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic(fmt.Sprintf("unknown CMState value: %d", state))		
	}
}

// ConsensusModule (CM) a single node of Raft consensus 
type ConsensusModule struct {
	mu sync.Mutex // Will be for all fields below

	id int // The server ID of this CM

	peerIds []int // The ID of our peers in the cluster
	
	server *Server // The server containing this CM. Will be used for RPC calls

	// Persistent states - for all servers 
	currentTerm int
	votedFor int // Candidate ID or -1 if not there 
	log []LogEntry // first index is 0 

	// Volatile states - for all servers
	state CMState
	electionResetEvent time.Time 

}






