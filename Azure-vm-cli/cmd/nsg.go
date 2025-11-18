package cmd

import (
    "context"
    "fmt"
    "time"

    "github.com/example/azure-vm-cli/internal/azure"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

// Flags
var (
    nsgName   string
    ruleName  string
    vmType    string
    priority  int32
	nsgCmd = &cobra.Command{
	Use:   "nsg",
	Short: "Manage Network Security Group (NSG) rules",
}
)

// NSG add command
var nsgAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Add an NSG rule to allow RDP or SSH",
    RunE: func(cmd *cobra.Command, args []string) error {
        if vmType != "linux" && vmType != "windows" {
            return fmt.Errorf("vm-type must be 'linux' or 'windows'")
        }

        port := int32(22)
        protocol := "Tcp"
        if vmType == "windows" {
            port = 3389
        }

        logrus.Infof("Adding NSG rule %s to NSG %s on port %d", ruleName, nsgName, port)

        client, err := azure.NewClient(cmd.Context(), subID)
        if err != nil {
            return err
        }

        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
        defer cancel()

        if priority == 0 {
            priority = 1001
        }

        if _, err := client.AddNSGRule(ctx, rgName, nsgName, ruleName, protocol, port, priority); err != nil {
            return fmt.Errorf("failed to add NSG rule: %w", err)
        }

        logrus.Infof("NSG rule %s added", ruleName)
        return nil
    },
}

func init() {
    nsgCmd.AddCommand(nsgAddCmd)
    nsgAddCmd.Flags().StringVarP(&nsgName, "nsg", "n", "", "NSG name (required)")
    nsgAddCmd.Flags().StringVarP(&ruleName, "rule", "r", "allow-ssh-rdp", "rule name")
    nsgAddCmd.Flags().StringVarP(&vmType, "vm-type", "t", "linux", "vm type: linux or windows")
    nsgAddCmd.Flags().Int32VarP(&priority, "priority", "p", 0, "rule priority (100-4096)")
    nsgAddCmd.Flags().StringVarP(&rgName, "resource-group", "g", "", "resource group name (required)")
    nsgAddCmd.Flags().StringVarP(&subID, "subscription", "s", "", "azure subscription id (optional)")

    nsgAddCmd.MarkFlagRequired("nsg")
    nsgAddCmd.MarkFlagRequired("resource-group")
}
