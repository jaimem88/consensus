package consensus

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

const (
	_ int64 = iota
	leader
	follower
	candidate
)

type state struct {
	state       int64
	currentTerm int64
	votedFor    string
	log         []string
	commitIndex int64
	lastApplied int64
}
type server struct {
	id              string
	state           *state
	termDuration    time.Duration
	electionTimeout time.Duration
	cluster         []RaftServiceClient
	leaderAliveTs   time.Time
}

var (
	mutex = &sync.Mutex{}
)

// NewDefaultCluster tries to connect to all specified ports
func (s *server) NewDefaultCluster(hosts []string) {
	clients := []RaftServiceClient{}
	for _, host := range hosts {
		log.Infof("dialing grpc service %s", host)
		conn, err := grpc.Dial(host,
			grpc.WithInsecure(),
		)
		if err != nil {
			log.Errorf("failed to dial user grpc : %s %s", host, err)
			continue
		}
		grpcClient := NewRaftServiceClient(conn)

		clients = append(clients, grpcClient)
	}
	s.cluster = clients
}

func New(id string, termDuration, electionTimeout time.Duration) *server {
	return &server{
		id:           id,
		termDuration: termDuration,
		state: &state{
			state:       follower,
			currentTerm: 0,
			votedFor:    "",
			log:         []string{},
		},
		electionTimeout: electionTimeout,
	}
}

// ElectionWatch keeps an eye in the election timeout
// to start a new election process when it hasn't received keep alives from the leader
func (s *server) ElectionWatch() {
	for {
		time.Sleep(s.electionTimeout)
		log.Infof("election watch with state %d", s.state.state)
		mutex.Lock()
		leaderTs := s.leaderAliveTs
		mutex.Unlock()

		if time.Since(leaderTs) > s.electionTimeout && s.state.state != leader {
			go s.initiateVote()
		}
	}
}

// AutoIncrements it's current term every s.termDuration
func (s *server) CurrentTerm() {

	for {
		time.Sleep(s.termDuration)
		mutex.Lock()
		s.state.currentTerm++
		mutex.Unlock()
	}

}

func (s *server) initiateVote() {

	log.Infof("iniitiated vote for node id: %s", s.id)
	mutex.Lock()
	s.state.state = candidate
	s.state.currentTerm++
	currentTerm := s.state.currentTerm
	s.leaderAliveTs = time.Now()
	cluster := s.cluster
	n := len(cluster) + 1 // include yourself in the cluster
	mutex.Unlock()
	if s.state.state != candidate {
		return
	}
	votes := 0
	for _, client := range cluster {

		voteResponse, err := client.RequestVote(context.Background(),
			&RequestVoteReq{
				CandidateId:  s.id,
				CurrentTerm:  currentTerm,
				LastLogIndex: s.state.commitIndex,
				LastLogTerm:  currentTerm - 1,
			})
		if err != nil {
			continue

		}
		if voteResponse.VoteGranted {
			votes++
		}

	}
	log.Infof("got votes: %d, nodes in cluster: %d", votes, n)
	if votes > n/2 {
		mutex.Lock()
		s.state.state = leader
		mutex.Unlock()
		go s.StayAlive()
		log.Infof("finished vote for node id: %s state: %d", s.id, s.state.state)
	}

	log.Infof("finished vote for node id: %s state: %d", s.id, s.state.state)

}

// StayAlive send this message to all clients
func (s *server) StayAlive() {
	if s.state.state != leader {
		return
	}
	for {
		time.Sleep(s.electionTimeout / time.Duration(len(s.cluster)))
		log.Infof("sending stay alive from: %s", s.id)
		mutex.Lock()
		currentTerm := s.state.currentTerm
		cluster := s.cluster
		mutex.Unlock()
		for _, client := range cluster {
			go func(client RaftServiceClient) {
				_, err := client.AppendEntries(context.Background(), &AppendEntriesReq{
					CurrentTerm: currentTerm,
					LeaderId:    s.id,
				})
				if err != nil {
					return
				}

			}(client)

		}
	}
}
