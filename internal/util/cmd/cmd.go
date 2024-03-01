package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/boywei/go-zero-check/internal/util/global"
)

func RunModel(host string, port int, uuid string) error {
	c := exec.Command("go", "run", "/home/stardust/goProjects/go-zero-check/main.go", "-host", host, "-port", strconv.Itoa(port))
	global.ModelIdMap[uuid].Cmd = c // 存储当前cmd命令到map中，方便删除模型时删除对应的进程
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
