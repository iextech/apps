package main

import (
    "fmt"
    "log"
    "context"
    "net"

    "google.golang.org/grpc"

    pb     "github.com/iextech/apps/pkg/ansibleapp"
)

type server struct {
    pb.UnimplementedAnsibleAppServiceServer
}


func (s *server) Exec(ctx context.Context, in *pb.AnsibleAppInput) (*pb.AnsibleAppOutput, error) {
    out := &pb.AnsibleAppOutput{ TaskFileContent : "Ansible File Content"}

    return out, nil
}

func (s *server) Help(ctx context.Context, in *pb.Topic) (*pb.HelpText, error) {
    topic := in.GetTopic()
    var   helptext    string
    if len(topic) == 0 {
        helptext = "General help"
    } else {
        helptext = "Help on " + topic
    }
    out := &pb.HelpText{ Text : helptext }

    return out, nil
}

func main() {
    fmt.Printf("Ansible App Server Example\n")

    port := ":50051"

    listener, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Printf("[ERROR] : %v\n", err)
        return
    }

    s := grpc.NewServer()

    pb.RegisterAnsibleAppServiceServer(s, &server{})

    fmt.Printf("server listening at %v", listener.Addr())
    if err := s.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
