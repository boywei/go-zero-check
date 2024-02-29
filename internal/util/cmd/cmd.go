package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
)

func RunModel(host string, port int) error {
	c := exec.Command("go", "run", "/home/stardust/goProjects/go-zero-check/main.go", "-host", host, "-port", strconv.Itoa(port))
	stdout, _ := c.StdoutPipe()
	reader := bufio.NewReader(stdout)

	go func() {
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			fmt.Println(string(line))
		}
	}()

	c.Start()
	c.Wait()
	return nil
}
