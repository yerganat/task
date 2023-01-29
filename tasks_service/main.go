package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"testtask/entity"
	"testtask/storage"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// parse arguments from the command line
	// this lets use define the port for the server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instanciate the server
	s := grpc.NewServer()

	// Register server method (actions the server il do)
	store := storage.New()
	entity.RegisterTasksServer(s, &tasksServer{store: store})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Implement the entity service (entity.tasksServer interface)
type tasksServer struct {
	entity.UnimplementedTasksServer
	store *storage.TaskStore
}

func (s *tasksServer) Save(ctx context.Context, n *entity.Task) (*entity.Status, error) {
	log.Printf("Recieved a task to save: %v", n.Url)

	id := s.store.CreateTask(n.Method, n.Url, n.Headers)

	return &entity.Status{Status: string(id)}, nil
}

func (s *tasksServer) Check(ctx context.Context, search *entity.TaskCheck) (*entity.Status, error) {
	log.Printf("Recieved a tasl to save: %v", search.Id)
	n, err := s.store.GetTask(int(search.Id))

	if err != nil {
		return &entity.Status{}, err
	}

	return &entity.Status{Status: n.Status}, nil
}
