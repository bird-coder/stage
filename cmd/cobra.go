/*
 * @Author: yujiajie
 * @Date: 2024-02-28 15:48:29
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-13 09:01:04
 * @FilePath: /stage/cmd/cobra.go
 * @Description:
 */
package cmd

import (
	"errors"
	"fmt"
	"os"
	"stage/cmd/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "stage",
	Short:        "stage",
	SilenceUsage: true,
	Long:         "stage",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Println("欢迎使用stage")
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
