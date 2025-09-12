# Developer's Foundry Bootcamp

## DAY 1 - DevOps Principles and Setting up of local environment.

- DevOps is like a service provider that ensures the clients are satisfied.
- DevOps is not just about tools. Linux is important.
- VirtualBox, orbstack.
- **Assignment:** Setup **Virtual Box**, Ubuntu server, Docker desktop, Setup **localstack** (Cloud Emulator).
- **VM SSH** - `ssh -i hammedkey Hammed@172.190.74.173`
- **hammedkey.pub**
``` ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIGgI/1DZ4/LLx52MgqhV99u91p1tIEBU2k9b60wkyPMK hammed@Hammed```
- **hammedkey**
	```
	-----BEGIN OPENSSH PRIVATE KEY-----
	b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	QyNTUxOQAAACBoCP9Q2ePyy8edjIKoVffbvdadbSBAVNpPW+tMJMjzCgAAAJCTlMJhk5TC
	YQAAAAtzc2gtZWQyNTUxOQAAACBoCP9Q2ePyy8edjIKoVffbvdadbSBAVNpPW+tMJMjzCg
	AAAECFdQMGjcfmp1A1MD013Bh7Wyen4DzcxQqirxyiwVhDxmgI/1DZ4/LLx52MgqhV99u9
	1p1tIEBU2k9b60wkyPMKAAAADWhhbW1lZEBIYW1tZWQ=
	-----END OPENSSH PRIVATE KEY-----
	```

## DAY 2

## DAY 3 - Introduction to Docker
**Tools**
 1. ctop
 2. lazydocker
 3. vim
- Docker Compose example
	```
	services:
		app:
			build:
			ports:
				- “8080:8080”
			Depends_on:
				- db
	```
## DAY 4 - Introduction to CI/CD - Continuous Delivery/Deployment
- In continuous integration workflow, we need a tool like GIT for version control.
- [nektos/act: Run your GitHub Actions locally 🚀](https://github.com/nektos/act)
- 

## DAY 5 - Infrastructure as Code IaC
- Terraform
- Pulumi
- Cloud Formation

In Terraform, we use command like "**terraform plan**" to connect to the cloud platform (e.g. AWS), check resource availability, check the stage, "**terraform apply**" to create resource, "**terraform destroy**" to delete your infra.

**Terraform state and state file** - tfstate - To store infra state files. It is preferably in an object form, for example AWS S3 bucket. We can **lock** the statefile from multiple updates by keeping keys in the Dynamo DB or the new method to add lock keys to terraform itself.
