package command

import (
	tools2 "github.com/amanhigh/go-fun/common/tools"
	"os"

	"github.com/spf13/cobra"
)

var (
	localPath = "/tmp/sshfs/"
)

var sshfsCmd = &cobra.Command{
	Use:   "sfs",
	Short: "Sshfs related",
}

var sshfsMountCmd = &cobra.Command{
	Use:   "m [host] [remotePath]",
	Short: "Mount ssfs",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		tools2.SshfsMount(args[0], args[1], localPath)
		os.Chdir(localPath)
	},
}

var sshfsUnmountCmd = &cobra.Command{
	Use:   "u",
	Short: "Unmounts Mount Point",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tools2.SshfsUnmount(localPath)
	},
}

func init() {
	sshfsCmd.Flags().StringVarP(&localPath, "local", "l", localPath, "Local Path")
	sshfsCmd.AddCommand(sshfsMountCmd, sshfsUnmountCmd)
	allCmd.AddCommand(sshfsCmd)
}
