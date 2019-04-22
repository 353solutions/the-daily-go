package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"nlp/pb"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func freePort() (int, error) {
	conn, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	conn.Close()
	return conn.Addr().(*net.TCPAddr).Port, nil
}

func runServer(require *require.Assertions) (*exec.Cmd, int) {
	exeFile := fmt.Sprintf("%s/nlpd-test", os.TempDir())
	cmd := exec.Command("go", "build", "-o", exeFile, ".")
	cmd.Stderr = os.Stderr
	require.NoError(cmd.Run(), "build")
	port, err := freePort()
	require.NoError(err, "get free port")
	cmd = exec.Command(exeFile)
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("NLP_GRPC_ADDR=:%d", port))
	err = cmd.Start()
	require.NoError(err, "run server")
	return cmd, port
}

func waitForServer(client pb.NLPClient) error {
	start := time.Now()
	req := &pb.PingRequest{}
	timeout := 10 * time.Second
	var err error
	for time.Now().Sub(start) < timeout {
		_, err = client.Ping(context.Background(), req)
		if err == nil {
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}

	return fmt.Errorf("server not ready after %s (%s)", timeout, err)
}

func TestGRPC(t *testing.T) {
	require := require.New(t)
	cmd, port := runServer(require)
	defer cmd.Process.Kill()

	addr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(err, "gRPC dial")
	client := pb.NewNLPClient(conn)
	err = waitForServer(client)
	require.NoError(err, "wait for server")

	req := &pb.TokenizeRequest{
		Text: "Make the zero value useful.",
	}

	resp, err := client.Tokenize(context.Background(), req)
	require.NoError(err, "call tokenize")
	expected := []string{"make", "the", "zero", "value", "useful"}
	require.Equal(expected, resp.Tokens, "tokens")

}
