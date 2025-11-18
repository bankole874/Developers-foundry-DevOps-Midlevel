# azure-vm-cli


A production-ready CLI to automate starting VMs and adding NSG rules in Azure, written in Go with Cobra.


## Features
- Authenticate using `DefaultAzureCredential` (supports Azure CLI credentials, environment variables, managed identities)
- Start an Azure VM (with polling)
- Add NSG rule to allow SSH (22) for Linux VMs or RDP (3389) for Windows VMs
- Configurable via `~/.config/azvmctl/config.yaml` or flags
- Structured logging (logrus)


## Quickstart
1. Install Go 1.20+
2. Set up Azure CLI or environment credentials, e.g. `az login`
3. Build


```bash
GO111MODULE=on go build -o azvmctl ./...
```


4. Start VM


```bash
./azvmctl start --vm my-linux-vm -g my-rg -s <subscription-id>
```


5. Add NSG rule for Linux


```bash
./azvmctl nsg add --nsg my-nsg -g my-rg --vm-type linux --rule allow-ssh
```


## Notes & next steps
- You may prefer to set `AZURE_SUBSCRIPTION_ID` in your environment rather than passing `-s` each call.
- The project intentionally uses the Azure SDK for Go (armcompute & armnetwork) and `azidentity.NewDefaultAzureCredential()`.
- For production: fill in `getenv` and configuration reading, add unit tests, add CI pipeline and static analysis, and improve error handling / retries.
