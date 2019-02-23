// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	rportserver "github.com/snwfdhmp/rport-server"
	"github.com/spf13/cobra"
)

var (
	serverPort  int
	storageFile string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start an rport server instance",
	Long:  `start an rport server instance listening to the given port, using the given storage file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := rportserver.NewServer().Start(serverPort, storageFile); err != nil {
			fmt.Printf("fatal: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().IntVarP(&serverPort, "port", "p", 9123, "server port")
	startCmd.Flags().StringVarP(&storageFile, "storage", "s", "./.rport-server.data.json", "storage path")
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
