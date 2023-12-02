package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

/*
	데드라인을 사용한 사용자 입력
	프로그램이 사용자에게 입력하도록 요청하고, 사용자는 5초 내로 입력을
	마치고 엔터를 눌려야 하며, 그렇지 못한 경우에는 그냥 기본값을 사용하는 상황을 가정
*/

var totalDuration time.Duration = 5

func getName(r io.Reader, w io.Writer) (string, error) {
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done"
	fmt.Fprintln(w, msg)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You entered an empty name")
	}
	return name, nil
}

func getNameContext(ctx context.Context) (string, error) {
	var err error
	name := "Default Name"
	c := make(chan error, 1)

	go func() {
		name, err = getName(os.Stdin, os.Stdout)
		c <- err
	}()

	select {
	case <-ctx.Done():
		return name, ctx.Err()
	case err := <-c:
		return name, err
	}
}

func main() {
	allowedDuration := totalDuration * time.Second

	// 부모 Context 객체, 콘텍스트가 만료되는 시간를 정의하는 time.Duration 객체
	// 별도로 사용할 부모 컨텍스트가 없기 때문에 context.Background()를 사용하여
	// 공백의 콘텍스트를 생성
	ctx, cancel := context.WithTimeout(context.Background(), allowedDuration)
	defer cancel() // 콘텍스트를 취소하는 함수
	// main함수가 반환되기 직전에 반드시 호출되도록 하여 콘텍스트를 해제

	name, err := getNameContext(ctx)

	// context.DeadlineExceeded인 경우 사용자에게 따로 에러를 출력하지 않고
	// 이름만 출력하며, 그 외의 에러가 발생한 경우에는 에러를 출력하고 0 외의
	// 종료 코드를 반환하고 프로그램을 종료
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, name)
}
