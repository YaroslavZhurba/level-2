package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"bufio"
	"errors"
	"path/filepath"
	"github.com/mitchellh/go-ps"
)

func processInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path argument is required")
		}
		return os.Chdir(args[1])
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(pwd)
	case "echo":
		for i := 1; i < len(args); i++ {
			fmt.Print(args[i], " ")
		}
		fmt.Println()
	case "kill":
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		return proc.Kill()
	case "ps":
		procs, err := ps.Processes()
		if err != nil {
			return err
		}
		for _, proc := range procs {
			fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
		}
	case "\\quit":
		os.Exit(0)
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}

	return nil
}

func RunShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		path, _ := filepath.Abs(".")
		path_splited := strings.Split(path, "/")
		current_directory := path_splited[len(path_splited) - 1]
		fmt.Printf("%s$ ", current_directory)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = processInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}