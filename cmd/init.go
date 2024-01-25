/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addr     string
	mWorkers int
	rWorkers int
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize master node",
	Long:  `Initialize a master node that controls the map/reduce workers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(addr, mWorkers, rWorkers)
		fmt.Println("This is kyle")
	},
}

func init() {
	initCmd.Flags().
		StringVarP(&addr, "addr", "a", "localhost:8000", "Address where the master node listens to")
	initCmd.Flags().
		IntVarP(&mWorkers, "map-workers", "m", 1, "Number of map workers to be used")
	initCmd.Flags().
		IntVarP(&rWorkers, "reduce-workers", "r", 1, "Number of reduce workers to be used")
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
