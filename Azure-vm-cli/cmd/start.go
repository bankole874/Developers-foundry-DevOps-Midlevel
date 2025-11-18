package cmd


import (
"context"
"fmt"
"time"


"github.com/spf13/cobra"


"github.com/example/azure-vm-cli/internal/azure"
"github.com/sirupsen/logrus"
)


var (
vmName string
)


var startCmd = &cobra.Command{
Use: "start",
Short: "Start a VM",
Long: "Start a VM by resource group and name",
RunE: func(cmd *cobra.Command, args []string) error {
logrus.Infof("Starting VM %s in rg %s (sub %s)", vmName, rgName, subID)
client, err := azure.NewClient(cmd.Context(), subID)
if err != nil {
return err
}
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()


if err := client.StartVM(ctx, rgName, vmName); err != nil {
    return fmt.Errorf("failed to start VM: %w", err)
}

logrus.Infof("VM %s started", vmName)

return nil
},
}


func init() {
startCmd.Flags().StringVarP(&vmName, "vm", "v", "", "VM name (required)")
startCmd.Flags().StringVarP(&rgName, "resource-group", "g", "", "resource group name (required)")
startCmd.Flags().StringVarP(&subID, "subscription", "s", "", "azure subscription id (optional, can be set via AZURE_SUBSCRIPTION_ID env or config)")
startCmd.MarkFlagRequired("vm")
startCmd.MarkFlagRequired("resource-group")
}
