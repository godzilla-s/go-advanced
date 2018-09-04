package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Command struct {
	//out bytes.Buffer
	output string
}

func (c *Command) Exec(s string) error {
	log.Println("exce: ", s)
	cmd := exec.Command("/bin/bash", "-c", s)
	//cmd.Stdout = &c.out
	//return cmd.Run()
	// out, err := cmd.Output() //运行命令，并返回标准输出
	// if err != nil {
	// 	log.Println("command exec error:", err)
	// 	return err
	// }
	//log.Println("output:", string(out))

	out, err := cmd.CombinedOutput() //运行命令，并返回标准输出和标准错误
	if err != nil {
		fmt.Println(err)
	}

	c.output = string(out)
	return nil
}

func (c Command) Output() string {
	return c.output //c.out.String()
}

func main() {
	var cmd Command

	cmd.Exec("peer channel join ./mychannel.block")
	fmt.Println(cmd.Output())
}
