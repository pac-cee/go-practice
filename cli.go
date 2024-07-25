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
