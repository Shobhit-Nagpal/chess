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

func ParseUciCommand(command string) {
	if command == "" {
		return
	}

	if isUciCommand(command) {
		parseCommand(command)
	}
}

func isUciCommand(command string) bool {
	commandParts := strings.Split(command, " ")
	baseCommand := commandParts[0]

	switch baseCommand {
	case Id.ToString(), UciOk.ToString(), ReadyOk.ToString(), Option.ToString(), Stockfish.ToString(), Info.ToString():
		return true
	default:
		return false
	}
}

func parseCommand(command string) {
	baseCommand := strings.Split(command, " ")[0]

	switch baseCommand {
	case Stockfish.ToString(), Option.ToString():
		return
	case Id.ToString():
		parseIdCommand(command)
	case UciOk.ToString():
		parseUciOkCommand()
	case ReadyOk.ToString():
		parseReadyOkCommand()
	case BestMove.ToString():
		parseBestMoveCommand(command)
	default:
		return
	}
}

func parseUciOkCommand() {
	fmt.Println("[server]: received", UciOk.ToString())
}

func parseReadyOkCommand() {
	fmt.Println("[server]: received", ReadyOk.ToString())
}

func parseIdCommand(idCmd string) {
	parts := strings.Split(idCmd, " ")
	param := parts[1]
	paramValue := ""

	switch param {
	case "name", "author":
		paramValue = strings.Join(parts[2:], " ")
	default:
		return
	}

	fmt.Println("[server]: received id", param, "-", paramValue)
}

func parseBestMoveCommand(bestmoveCommand string) {
	move := strings.Split(bestmoveCommand, " ")[1]
	fmt.Println("[server]:  best move -", move)
}
