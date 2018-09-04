package main

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jaimemartinez88/consensus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// grpcServerKeepalive is a duration that if the server doesn't see any activity after this time, it pings the client to see if the transport is still alive.
// https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#connection-idle-timeout
// we use AWS Network Load Balancer that has fixed, unmodifiable 350 sec idle timeout.
// We are enabling keep alive to prevent load balancer to reset the connection
// https://translate.google.com.au/translate?hl=en&sl=ja&u=https://hori-ryota.com/blog/failed-to-grpc-with-nlb/&prev=search
// Setting 150 sec, which is less than half.
const grpcServerKeepalive = 150 * time.Second

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	serverID := os.Getenv("SERVICE_NAME")
	serverAddr := os.Getenv("SERVICE_ADDR")
	clusterHosts := os.Getenv("CLUSTER_HOSTS")
	electionTimeout := os.Getenv("ELECTION_TIMEOUT")

	grpcServer, err := createServer(serverID, serverAddr, clusterHosts, electionTimeout)
	if err != nil {
		log.Fatalf("failed to start grpc service: %s", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info("shutting down server...")
		grpcServer.GracefulStop()
		log.Println("server gracefully stopped. bye")
	}()

	log.Printf("server ready! %s (pid: %d)", serverID, os.Getpid())
	serveGRPC(grpcServer, serverAddr)

}

func createServer(id, addr, clusterPorts, electionTimeout string) (*grpc.Server, error) {
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{Time: grpcServerKeepalive}),
	)

	eToutMsec, err := strconv.Atoi(electionTimeout)
	if err != nil {
		return nil, err
	}

	termDuration := time.Second * 1
	eTimeout := time.Duration(eToutMsec) * time.Millisecond

	ports := strings.Split(clusterPorts, ",")

	server := consensus.New(id, termDuration, eTimeout)

	consensus.RegisterRaftServiceServer(s, server)
	// wait for services to register
	time.Sleep(2 * time.Second)
	server.NewDefaultCluster(ports)
	go server.ElectionWatch()
	go server.CurrentTerm()
	go server.StayAlive()

	return s, nil
}

func serveGRPC(s *grpc.Server, grpcAddr string) {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infof("grpc   :%s listening...", grpcAddr)
	err = s.Serve(lis)
	if err != nil {
		log.Warn("gRPC terminating Serve() %s", err)
	}
}
