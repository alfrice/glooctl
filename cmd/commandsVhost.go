package cmd

import (
	common "github.com/solo-io/glooctl/platform/executor"
	"github.com/spf13/cobra"
)

var vhost *VHost

// vhostCmd represents the vhost command
var vhostCmd = &cobra.Command{
	Use:   "vhost",
	Short: "create virtual host",
	Long:  `Create Virtual Host`,
	Run: func(cmd *cobra.Command, args []string) {
		InteractiveModeVhost("create", vhost)
		common.GetExecutor("vhost", GetGlobalFlags().Namespace).RunCreate(GetGlobalFlags(), GetVhostParams())
	},
}

var vhostDelCmd = &cobra.Command{
	Use:   "vhost",
	Short: "delete virtual host",
	Long:  `Delete Virtual Host`,
	Run: func(cmd *cobra.Command, args []string) {
		InteractiveModeVhost("delete", vhost)
		common.GetExecutor("vhost", GetGlobalFlags().Namespace).RunDelete(GetGlobalFlags(), GetVhostParams())
	},
}

var vhostDescribeCmd = &cobra.Command{
	Use:   "vhost",
	Short: "describe virtual host",
	Long:  `Describe Virtual Host`,
	Run: func(cmd *cobra.Command, args []string) {
		InteractiveModeVhost("describe", vhost)
		common.GetExecutor("vhost", GetGlobalFlags().Namespace).RunDescribe(GetGlobalFlags(), GetVhostParams())
	},
}

var vhostGetCmd = &cobra.Command{
	Use:   "vhost",
	Short: "get virtual host",
	Long:  `Get Virtual Host`,
	Run: func(cmd *cobra.Command, args []string) {
		InteractiveModeVhost("get", vhost)
		common.GetExecutor("vhost", GetGlobalFlags().Namespace).RunGet(GetGlobalFlags(), GetVhostParams())
	},
}

var vhostUpdateCmd = &cobra.Command{
	Use:   "vhost",
	Short: "update virtual host",
	Long:  `Update Virtual Host`,
	Run: func(cmd *cobra.Command, args []string) {
		InteractiveModeVhost("update", vhost)
		common.GetExecutor("vhost", GetGlobalFlags().Namespace).RunUpdate(GetGlobalFlags(), GetVhostParams())
	},
}

func init() {
	vhost = NewVHost()
	createCmd.AddCommand(vhostCmd)
	deleteCmd.AddCommand(vhostDelCmd)
	describeCmd.AddCommand(vhostDescribeCmd)
	getCmd.AddCommand(vhostGetCmd)
	updateCmd.AddCommand(vhostUpdateCmd)
}
