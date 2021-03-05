package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"kn-plugin-diag/pkg/command/diagnose"
	"kn-plugin-diag/pkg/utils"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"

)


func main() {

	p := &utils.ConnectionConfig{}
	p.Initialize()

	rootCmd := &cobra.Command{
		Use:   "kn-diag",
		Short: "A CLI of to help with Diagnose Knative Resources",
		Long:  `A CLI of to help with Diagnose Knative Resources.`,
	}
	rootCmd.AddCommand(diagnose.NewServiceCmd(p))
	rootCmd.InitDefaultHelpCmd()
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed : ", err)
		os.Exit(1)
	}
	
}
