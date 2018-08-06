package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// root 默认，不带参数程序运行
var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "test is simple program for cobra",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("root command run")
	},
}

// 定义子命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Args:  cobra.NoArgs, // 没有参数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verion: v1.0.0")
	},
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echo message",
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("echo message: name:", name, " id:", id)
	},
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "show time now",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("now: ", time.Now(), name)
	},
}

var (
	msg string
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print message",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("print prefix =======")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print: ", msg)
	},
}

var (
	name string
	id   int
)

func init() {
	//rootCmd.Flags().StringVarP(&name, "name", "n", "", "name set")
	//rootCmd.Flags().IntVarP(&id, "id", "i", 0, "identity")

	// 顶级命令
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(printCmd)

	echoCmd.Flags().StringVarP(&name, "name", "n", "", "name set")
	echoCmd.Flags().IntVarP(&id, "id", "i", 0, "identity")
	// 子命令
	echoCmd.AddCommand(timeCmd)

	printCmd.Flags().StringVarP(&msg, "msg", "m", "", "message of print")

	//timeCmd.Flags().StringVarP()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("exec:", err)
		os.Exit(1)
	}
}
