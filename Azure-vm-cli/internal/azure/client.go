package azure

import (
    "context"
    "fmt"
    "os"

    "github.com/Azure/azure-sdk-for-go/sdk/azcore"
    "github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

type Client struct {
    subscriptionID string
    vmClient       *armcompute.VirtualMachinesClient
    ruleClient     *armnetwork.SecurityRulesClient
    creds          azcore.TokenCredential
}

// ------------------------------------------------------------
//  New Azure Client
// ------------------------------------------------------------
func NewClient(ctx context.Context, subscriptionID string) (*Client, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        return nil, fmt.Errorf("failed to load Azure credentials: %w", err)
    }

    if subscriptionID == "" {
        subscriptionID = getSubscriptionFromEnvOrConfig()
    }

    vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create VM client: %w", err)
    }

    ruleClient, err := armnetwork.NewSecurityRulesClient(subscriptionID, cred, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create NSG rule client: %w", err)
    }

    return &Client{
        subscriptionID: subscriptionID,
        vmClient:       vmClient,
        ruleClient:     ruleClient,
        creds:          cred,
    }, nil
}

// ------------------------------------------------------------
//  Start VM  (updated for modern SDK)
// ------------------------------------------------------------
func (c *Client) StartVM(ctx context.Context, resourceGroup, vmName string) error {
    poller, err := c.vmClient.BeginStart(ctx, resourceGroup, vmName, nil)
    if err != nil {
        return fmt.Errorf("begin start VM failed: %w", err)
    }

    _, err = poller.PollUntilDone(ctx, nil)
    if err != nil {
        return fmt.Errorf("polling VM start failed: %w", err)
    }

    return nil
}

// ------------------------------------------------------------
//  Add NSG Rule (updated for SDK v1.0.0)
// ------------------------------------------------------------
func (c *Client) AddNSGRule(
    ctx context.Context,
    resourceGroup, nsgName, ruleName, protocol string,
    port int32,
    priority int32,
) (*armnetwork.SecurityRule, error) {

    rule := armnetwork.SecurityRule{
        Properties: &armnetwork.SecurityRulePropertiesFormat{
            Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocol(protocol)),
            SourcePortRange:          to.Ptr("*"),
            DestinationPortRange:     to.Ptr(fmt.Sprintf("%d", port)),
            SourceAddressPrefix:      to.Ptr("*"),
            DestinationAddressPrefix: to.Ptr("*"),
            Access:                   to.Ptr(armnetwork.SecurityRuleAccessAllow),
            Priority:                 to.Ptr(priority),
            Direction:                to.Ptr(armnetwork.SecurityRuleDirectionInbound),
        },
    }

    poller, err := c.ruleClient.BeginCreateOrUpdate(
        ctx,
        resourceGroup,
        nsgName,
        ruleName,
        rule,
        nil,
    )
    if err != nil {
        return nil, fmt.Errorf("begin NSG rule create failed: %w", err)
    }

    resp, err := poller.PollUntilDone(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("polling NSG rule creation failed: %w", err)
    }

    return &resp.SecurityRule, nil
}

// ------------------------------------------------------------
//  Subscription from env
// ------------------------------------------------------------
func getSubscriptionFromEnvOrConfig() string {
    if v := os.Getenv("AZURE_SUBSCRIPTION_ID"); v != "" {
        return v
    }
    return ""
}
