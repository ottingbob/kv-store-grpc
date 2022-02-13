// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/ottingbob/kv-store-grpc/pkg/server"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "run the REST Key Value server",
	Long:  `run the REST Key Value server`,
	Run: func(cmd *cobra.Command, args []string) {
		rpcEndpoint, err := cmd.Flags().GetString("rpc-endpoint")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := server.RunREST(rpcEndpoint); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(restCmd)
	restCmd.Flags().StringP("rpc-endpoint", "r", "localhost", "RPC Endpoint the REST service should connect to")
}
