/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kylehipz/map-reduce-go/master"
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
		master.Init(addr, mWorkers, rWorkers)
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
}
