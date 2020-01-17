/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	zsh "github.com/rsteube/cobra-zsh-gen"
	"github.com/spf13/cobra"
	"os"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:     "completion",
	Short:   "Generates bash/zsh completion scripts",
	Long:    `Output shell completion code for the specified shell (bash or zsh).`,
	Example: completionExample(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			complet := args[0]
			if complet == "bash" {
				rootCmd.GenBashCompletion(os.Stdout)
			} else if complet == "zsh" {
				zsh.Wrap(rootCmd).GenZshCompletion(os.Stdout)
			} else {
				Warning.Println("Parameter error! Please input bash or zsh")
			}
		} else {
			Warning.Println("Please input bash or zsh.")
		}

	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.SetArgs([]string{""})
}

func completionExample() string {
	return `
# bash
kubecm completion bash > ~/.kube/kubecm.bash.inc
printf "
# kubecm shell completion
source '$HOME/.kube/kubecm.bash.inc'
" >> $HOME/.bash_profile
source $HOME/.bash_profile

# add to $HOME/.zshrc
source <(kubecm completion zsh)
# or
kubecm completion zsh > "${fpath[1]}/_kubecm"
`
}
