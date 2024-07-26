package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Execute     func(args []string) error
}

type CLI struct {
	commands    map[string]Command
	history     []string
	running     bool
	workingDir  string
	maxHistory  int
	promptStyle string
}

func NewCLI() *CLI {
	cli := &CLI{
		commands:    make(map[string]Command),
		history:     make([]string, 0),
		running:     true,
		workingDir:  ".",
		maxHistory:  100,
		promptStyle: "default",
	}
	cli.registerCommands()
	return cli
}

func (c *CLI) registerCommands() {
	c.commands = map[string]Command{
		"help": {
			Name:        "help",
			Description: "Show available commands",
			Usage:       "help [command]",
			Execute:     c.cmdHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the CLI",
			Usage:       "exit",
			Execute:     c.cmdExit,
		},
		"clear": {
			Name:        "clear",
			Description: "Clear the screen",
			Usage:       "clear",
			Execute:     c.cmdClear,
		},
		"time": {
			Name:        "time",
			Description: "Show current time and date",
			Usage:       "time",
			Execute:     c.cmdTime,
		},
		"history": {
			Name:        "history",
			Description: "Show command history",
			Usage:       "history [count]",
			Execute:     c.cmdHistory,
		},
		"cd": {
			Name:        "cd",
			Description: "Change working directory",
			Usage:       "cd <directory>",
			Execute:     c.cmdCd,
		},
		"pwd": {
			Name:        "pwd",
			Description: "Print working directory",
			Usage:       "pwd",
			Execute:     c.cmdPwd,
		},
		"ls": {
			Name:        "ls",
			Description: "List directory contents",
			Usage:       "ls [directory]",
			Execute:     c.cmdLs,
		},
		"prompt": {
			Name:        "prompt",
