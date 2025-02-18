package main

import (
    "fmt"
    "log"
    "context"
    "time"
    "flag"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    pb     "github.com/iextech/apps/pkg/ansibleapp"
)

func main() {
    fmt.Printf("Ansible App Client Example\n")

    hostname := flag.String("host", "localhost", "Server hostname")
    port     := flag.Int("port", 50051, "Server port to connect")
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

    in          := &pb.AnsibleAppInput{ Task : "hello" }
    out, err    := client.Exec(ctx, in)
    if err != nil {
        log.Fatalf("[ERROR] : %v", err)
    }

    ansible_file_content := out.GetTaskFileContent()
    fmt.Printf("%s\n", ansible_file_content)

    inTopic := &pb.Topic{}
    help_g, _ := client.Help(ctx, inTopic)
    fmt.Println(help_g.GetText())

    inTopic   = &pb.Topic{ Topic : "hello" }
    help_g, _ = client.Help(ctx, inTopic)
    fmt.Println(help_g.GetText())

    return
}

