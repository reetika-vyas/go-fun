package command

import (
	"fmt"
	tools2 "github.com/amanhigh/go-fun/common/tools"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	time int
)
var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Random Utility Commands",
	Args:  cobra.ExactArgs(1),
}

var pprofCmd = &cobra.Command{
	Use:   "pprof [Host] [Port]",
	Short: "Go Profiling with Go Torch",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		port := args[1]
		url := fmt.Sprintf("http://%v:%v/debug/pprof/profile", host, port)

		color.Blue("Profiling: %v for %v Seconds", url, time)
		tools2.RunCommandPrintError(fmt.Sprintf("go-torch -t %v -u %v && open torch.svg", time, url))
		tools2.RunCommandPrintError(fmt.Sprintf("go tool pprof -svg -output pprof.svg --seconds=%v %v && open pprof.svg", time, url))
	},
}

func init() {
	pprofCmd.Flags().IntVarP(&time, "time", "t", 30, "Profiling Time")

	utilCmd.AddCommand(pprofCmd)
	RootCmd.AddCommand(utilCmd)
}
