build:
	go build -o terraform-provider-moltin

plan:
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform plan

apply:
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform apply

destroy:
	export TF_LOG=ERROR && \
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform destroy

refresh:
	export TF_LOG=ERROR && \
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform refresh
