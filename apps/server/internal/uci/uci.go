package uci

import (
	"fmt"
	"strings"
)

type UciCommand struct {
	command string
}

func NewCommand(command string) *UciCommand {
	return &UciCommand{
		command: command,
	}
}

func (uc *UciCommand) EnsureNewLine() string {
	if strings.HasSuffix(uc.command, "\n") {
		return uc.command
	}

	return uc.command + "\n"
}

func (uc *UciCommand) ToString() string {
	return uc.command
}

func (uc *UciCommand) WithParam(key, value string) *UciCommand {
	uc.command += " " + key + " " + value
	return uc
}

func (uc *UciCommand) WithFlag(flag string) *UciCommand {
	uc.command += " " + flag
	return uc
}

func ParseUciCommand(command string) string {
	if command == "" {
		return command
	}

	commandParts := strings.Split(command, " ")

	if !isUciCommand(commandParts[0]) {
		return ""
	}

	return command
}

func isUciCommand(command string) bool {
	switch command {
	case Id.ToString(), UciOk.ToString(), ReadyOk.ToString(), Option.ToString(), Stockfish.ToString():
		fmt.Println(command)
		return true
	default:
		fmt.Println("NOT A UCI COMMAND")
		return false
	}
}
