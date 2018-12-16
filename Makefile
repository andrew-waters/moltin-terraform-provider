build:
	go build -o terraform-provider-moltin

apply:
	export TF_LOG=ERROR && \
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform apply

destroy:
	export TF_LOG=ERROR && \
	go build -o terraform-provider-moltin && \
	terraform init && \
	terraform destroy