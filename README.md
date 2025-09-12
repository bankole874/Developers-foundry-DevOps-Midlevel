# Developer's Foundry Bootcamp

## DAY 1 - DevOps Principles and Setting up of local environment.

- DevOps is like a service provider that ensures the clients are satisfied.
- DevOps is not just about tools. Linux is important.
- VirtualBox, orbstack.
- **Assignment:** Setup **Virtual Box**, Ubuntu server, Docker desktop, Setup **localstack** (Cloud Emulator).
- [Aviatrix-ACE Multicloud Network Associate Course - Google Docs](https://docs.google.com/document/d/1J14t4qmMbh_Ey1PJFTx3ZqvJwd6EqCIx4rNjlbEoxrI/edit?tab=t.yszmcjeh44me)

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
