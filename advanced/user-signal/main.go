package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func executeCommand(ctx context.Context, command string, arg string) error {
	return exec.CommandContext(ctx, command, arg).Run()
}

func setupSignalHandler(w io.Writer, cancelFunc context.CancelFunc) {
	// Signal 구조체 타입의 처리 수용력capacity이 1인 채널을 생성
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-c
		fmt.Fprintf(w, "Got signal: %v\n", s)
		cancelFunc()
	}()
}

func createContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	// WithTimeout() : 사용자가 지정한 만큼의 시간이 만료되면 취소되는 컨택스트를 생성할 때 사용
	ctx, cancel := context.WithTimeout(context.Background(), d)
	return ctx, cancel
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // CommandContext() 함수는 컨텍스트가 만료되면 커맨드로 실행시키는 외부
	// // 프로그램을 강제로 종료 force kill 수행
	// if err := exec.CommandContext(ctx, "sleep", "20").Run(); err != nil {
	// 	fmt.Fprintln(os.Stdout, err)
	// }

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <command> <argument>\n", os.Args[0])
		os.Exit(1)
	}
	command := os.Args[1]
	arg := os.Args[2]

	cmdTimeout := 30 * time.Second
	ctx, cancel := createContextWithTimeout(cmdTimeout)
	defer cancel()

	setupSignalHandler(os.Stdout, cancel)

	err := executeCommand(ctx, command, arg)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
