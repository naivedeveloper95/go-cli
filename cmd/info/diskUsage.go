/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Prints the disk usage information of the current directory",
	Long:  `Prints the disk usage information of the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultDir := "."
		if dir := viper.GetViper().GetString("cmd.info.diskUsage.defaultDir"); dir != "" {
			defaultDir = dir
		}
		usage := du.NewDiskUsage(defaultDir).Usage()
		fmt.Printf("Space used by %v by the directory %s\n", usage, defaultDir)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
