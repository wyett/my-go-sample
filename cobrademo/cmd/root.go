//
//Author: your name
//Date: 2021-10-25 11:22:51
//LastEditTime: 2021-10-25 17:55:57
//LastEditors: Please set LastEditors
//Description: In User Settings Edit
//FilePath: \my-go-sample\cobrademo\cmd\root.go
//
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(rootCmd)
}
