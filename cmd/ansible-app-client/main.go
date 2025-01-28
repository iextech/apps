package main

import (
    "fmt"
    "log"
    "context"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    pb     "github.com/iextech/apps/pkg/ansibleapp"
)

func main() {
    fmt.Printf("Ansible App Client Example\n")

    address := "localhost:50051"
    //tls_opts   := []grpc.DialOption { insecure.NewCredentials() }
    //grpc.WithTransportCredentials(insecure.NewCredentials())
    conn, err  := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("[ERROR] : %v", err)
    }
    defer conn.Close()

    client      := pb.NewAnsibleAppServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    in          := &pb.AnsibleAppInput{}
    out, err    := client.Exec(ctx, in)
    if err != nil {
        log.Fatalf("[ERROR] : %v", err)
    }

    ansible_file_content := out.GetTaskFileContent()
    fmt.Printf("%s\n", ansible_file_content)

    return
}

