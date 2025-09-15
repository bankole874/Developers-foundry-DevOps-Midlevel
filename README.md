# Developer's Foundry Bootcamp

- [StackEdit](https://stackedit.io/app#)
- [Aviatrix-ACE Multicloud Network Associate Course - Google Docs](https://docs.google.com/document/d/1J14t4qmMbh_Ey1PJFTx3ZqvJwd6EqCIx4rNjlbEoxrI/edit?tab=t.yszmcjeh44me)

## DAY 1.1 - DevOps Principles and Setting up of local environment.

- DevOps is like a service provider that ensures the clients are satisfied.
- DevOps is not just about tools. Linux is important.
- VirtualBox, orbstack.
- **Assignment:** Setup **Virtual Box**, Ubuntu server, Docker desktop, Setup **localstack** (Cloud Emulator).

## DAY 1.2 - Introduction to Linux
- Learnt about **commands** in Linux.
![No alternative text description for this image](https://media.licdn.com/dms/image/v2/D4D22AQFP69cIcN2jVQ/feedshare-shrink_800/B4DZYalXoTG8Ao-/0/1744202726930?e=1760572800&v=beta&t=vvB_wbRQojivvVzwqnJeAZc3TQpmSVo0qIElVW2vgCM)

## DAY 1.3 - Introduction to Docker
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
## DAY 1.4 - Introduction to CI/CD - Continuous Delivery/Deployment
- In continuous integration workflow, we need a tool like GIT for version control.
- [nektos/act: Run your GitHub Actions locally 🚀](https://github.com/nektos/act)

## DAY 1.5 - Infrastructure as Code IaC
- Terraform
- Pulumi
- Cloud Formation

In Terraform, we use command like "**terraform plan**" to connect to the cloud platform (e.g. AWS), check resource availability, check the stage, "**terraform apply**" to create resource, "**terraform destroy**" to delete your infra.

**Terraform state and state file** - tfstate - To store infra state files. It is preferably in an object form, for example AWS S3 bucket. We can **lock** the statefile from multiple updates by keeping keys in the Dynamo DB or the new method to add lock keys to terraform itself.
