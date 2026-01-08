
## Goal
Create a four-node Kubernetes cluster using KinD and K3s.

## Results
| Cluster Type | Status | Notes |
| :--- | :--- | :--- |
| **K3s (K3d)** | Success | Created 4-node cluster (1 Server, 3 Agents). |
| **KinD** | Failed | Failed due to environment limitations (cgroup/systemd issues). |

## K3s Cluster (K3d)
Successfully created a 4-node cluster using `k3d`.

## Verification
```bash 
kubectl get nodes --context k3d-k3d-4node
```

#### Output:
<img width="1110" height="192" alt="image" src="https://github.com/user-attachments/assets/f080b67a-b203-40c0-8915-1f37243902ee" />


## Steps to Reproduce (k3s)
#### Install K3d:
```bash
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | sudo bash
```

#### Create Cluster:
```bash
k3d cluster create k3d-4node --servers 1 --agents 3
```

## Steps to Reproduce (kinD)

#### Install Kind:
```bash
# For AMD64
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64

# For ARM64
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-arm64

chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

#### Create Config (kind-config.yaml):
```bash
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
- role: worker
```

#### Create Cluster:
```kind create cluster --config kind-config.yaml --name kind-4node```

## Issue Encountered
The ```kind create cluster``` command failed with: ```ERROR: failed to create cluster: could not find a log line that matches "Reached target .*Multi-User System.*|detected cgroup v1"```

This indicates an issue with running systemd/cgroups within the containerized environment. K3d was able to bypass this limitation.

    
