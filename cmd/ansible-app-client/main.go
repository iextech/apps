package main

import (
    "fmt"
    "log"
    "context"
    "time"
    "flag"
    "os"
    "path/filepath"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    pb     "github.com/iextech/apps/pkg/ansibleapp"
)

func main() {
    fmt.Printf("Ansible App Client Example\n")

    hostname := flag.String("host", "localhost", "Server hostname")
    port     := flag.Int("port", 50051, "Server port to connect")
    task     := flag.String("task", "hello", "Task to request")
    flag.Parse()

    address := fmt.Sprintf("%s:%d", *hostname, *port)
    conn, err  := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("[ERROR] : %v", err)
    }
    defer conn.Close()

    client      := pb.NewAnsibleAppServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    in          := &pb.AnsibleAppInput{ Task : *task }
    out, err    := client.Exec(ctx, in)
    if err != nil {
        log.Fatalf("[ERROR] : %v", err)
    }

    switch out.GetFileType() {
        case pb.AnsibleAppOutput_TEXT:
            fmt.Println("TEXT output")
            ansible_file_content := out.GetTaskFileContent()
            fmt.Printf("%s\n", ansible_file_content)
        case pb.AnsibleAppOutput_TAR:
            fmt.Println("TAR output")
            cwd, err := os.Getwd()
            if err != nil {
                log.Fatalf("%v\n", err)
            }
            output_file_path := filepath.Join(cwd, fmt.Sprintf("%s.tar", *task))
            err = os.WriteFile(output_file_path, out.GetBinaryFileContent(), 0600)
            if err != nil {
                log.Fatalf("%v\n", err)
            }
        case pb.AnsibleAppOutput_TGZ:
            fmt.Println("TGZ output")
            cwd, err := os.Getwd()
            if err != nil {
                log.Fatalf("%v\n", err)
            }
            output_file_path := filepath.Join(cwd, fmt.Sprintf("%s.tgz", *task))
            err = os.WriteFile(output_file_path, out.GetBinaryFileContent(), 0600)
            if err != nil {
                log.Fatalf("%v\n", err)
            }
        default:
            fmt.Printf("[WARN]: Unexpected output file type")
    }

    inTopic := &pb.Topic{}
    help_g, _ := client.Help(ctx, inTopic)
    fmt.Println(help_g.GetText())

    inTopic   = &pb.Topic{ Topic : "hello" }
    help_g, _ = client.Help(ctx, inTopic)
    fmt.Println(help_g.GetText())

    return
}

