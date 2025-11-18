package cmd

import (
"fmt"
"os"


"github.com/spf13/cobra"
"github.com/spf13/viper"


"github.com/example/azure-vm-cli/pkg/logging"
)

var (
    rgName string
    subID  string
)

var cfgFile string
var rootCmd = &cobra.Command{
Use: "azvmctl",
Short: "azvmctl — manage Azure VMs & NSG rules",
Long: "A small production-ready CLI to authenticate to Azure, start VMs and add NSG rules for Linux and Windows VMs.",
}


func Execute() {
if err := rootCmd.Execute(); err != nil {
fmt.Println(err)
os.Exit(1)
}
}


func init() {
cobra.OnInitialize(initConfig)


rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/azvmctl/config.yaml)")


rootCmd.PersistentFlags().StringP("log-level", "l", "info", "log level: debug, info, warn, error")
viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))


// subcommands
rootCmd.AddCommand(startCmd)
rootCmd.AddCommand(nsgCmd)
}


func initConfig() {
if cfgFile != "" {
viper.SetConfigFile(cfgFile)
} else {
viper.AddConfigPath("$HOME/.config/azvmctl")
viper.SetConfigName("config")
viper.SetConfigType("yaml")
}


viper.AutomaticEnv()


if err := viper.ReadInConfig(); err != nil {
// not fatal — config can be created on demand
fmt.Printf("Using default config values (no config found): %v\n", err)
}


logging.Setup(viper.GetString("log.level"))
}
