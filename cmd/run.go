package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// flags
var timeout int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a command, with the specified timeout",
	Long: `Runs a command, setting a timeout. If the command takes longer
to run than the specified timeout, then the command is killed.
By default, the timeout is 15 seconds`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

		runCmd := exec.CommandContext(ctx, args[0], args[1:]...)
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		if err := runCmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if ctx.Err() != nil {
					fmt.Println("Process timed out")
					os.Exit(42)
				}
				exitCode := exitError.Sys().(syscall.WaitStatus)
				os.Exit(exitCode.ExitStatus())
			} else {
				log.Printf("Could not run process: %v", err)
				os.Exit(17)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&timeout, "timeout", "t", 15, "The timeout, in seconds")
}
