package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

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

func (s *tasksServer) Save(ctx context.Context, n *entity.Task) (*entity.TaskId, error) {
	log.Printf("Recieved a task to save: %v", n.Url)

	id := s.store.CreateTask(n.Method, n.Url, n.Headers)

	go func(taskId int) {
		fmt.Println("Run crawler")
		s.store.UpdateTask(taskId, storage.TaskProgress, "")
		task, _ := s.store.GetTask(taskId)

		req, err := http.NewRequest(task.Method, task.Url, nil)
		if err != nil {
			s.store.UpdateTask(taskId, storage.TaskError, fmt.Sprintf("%v", err))
			return
		}

		for header, val := range task.Headers {
			req.Header.Add(header, val)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			s.store.UpdateTask(taskId, storage.TaskError, fmt.Sprintf("%v", err))
			return
		}

		s.store.UpdateTask(taskId, storage.TaskDone, fmt.Sprintf("%v", res))
	}(id)

	return &entity.TaskId{Id: int32(id)}, nil
}

func (s *tasksServer) Check(ctx context.Context, search *entity.TaskId) (*entity.Status, error) {
	log.Printf("Recieved a tasl to save: %v", search.Id)
	n, err := s.store.GetTask(int(search.Id))

	if err != nil {
		return &entity.Status{}, err
	}

	return &entity.Status{Status: n.Status}, nil
}
