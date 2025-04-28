package main

import (
    "fmt"
    "log"
    "context"
    "net"
    "flag"
    "os"

    "google.golang.org/grpc"

    pb     "github.com/iextech/apps/pkg/ansibleapp"
)

var (
    g_output_type      string
)

type server struct {
    pb.UnimplementedAnsibleAppServiceServer
}


func (s *server) Exec(ctx context.Context, in *pb.AnsibleAppInput) (*pb.AnsibleAppOutput, error) {

    out := &pb.AnsibleAppOutput{ TaskFileContent : "Ansible File Content"}

    switch g_output_type {
        case "txt":
            dat, err := os.ReadFile("./hello.yaml")
            if err != nil {
                return out, err
            }
            out = &pb.AnsibleAppOutput { FileType : pb.AnsibleAppOutput_TEXT , TaskFileContent : string(dat)}
        case "tar":
            dat, err := os.ReadFile("./hello.tar")
            if err != nil {
                return out, err
            }
            out = &pb.AnsibleAppOutput { FileType : pb.AnsibleAppOutput_TAR , BinaryFileContent : dat}
        case "tgz":
            dat, err := os.ReadFile("./hello.tgz")
            if err != nil {
                return out, err
            }
            out = &pb.AnsibleAppOutput { FileType : pb.AnsibleAppOutput_TGZ , BinaryFileContent : dat}
        default:
           return out, fmt.Errorf("Unexpected output type")
    }

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
    output_type := flag.String("type", "text", "Output type - txt/tar/tgz")
    flag.Parse()

    g_output_type = *output_type

    fmt.Printf("Ansible App Server Example for output type - %s\n", *output_type)

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
