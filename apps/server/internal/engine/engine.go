package engine

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/Shobhit-Nagpal/chess/apps/server/internal/utils"
)

type Engine struct {
	name      string
	path      string
	isRunning bool
	stdin     io.WriteCloser
	stdout    io.ReadCloser
	scanner   *bufio.Scanner
}

func New() *Engine {
	return &Engine{
		name: "stockfish",
		path: utils.GetStockfishPath(),
	}
}

func (e *Engine) Spawn() error {
	cmd := exec.Command(e.path)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
		return err
	}
	e.stdin = stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return err
	}
	e.stdout = stdout

	e.scanner = bufio.NewScanner(e.stdout)

	err = cmd.Start()

	if err != nil {
		log.Println(err)
		return err
	}

	e.isRunning = true

	return nil
}

func (e *Engine) SendCommand(command string) error {
	_, err := fmt.Fprint(e.stdin, command)
	return err
}

func (e *Engine) IsRunning() bool {
	return e.isRunning
}

func (e *Engine) ReadResponse() string {
	if e.scanner.Scan() {
		return e.scanner.Text()
	}

	err := e.scanner.Err()
	if err != nil {
		log.Println(err)
	}

	return ""
}
