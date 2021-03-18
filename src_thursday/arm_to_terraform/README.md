# Managing Arm templates to Terraform

## Azure resource group

[Terraform import documentation](https://www.terraform.io/docs/cli/import/index.html)

[Terraform state documentation](https://www.terraform.io/docs/cli/commands/refresh.html)

### Installation

- Navigate to working directory
- Gather the ID of the resource you want to bring under terraform management
- A resource block must be created within main.tf to map the importing resource *ex: azurerm_resource_group.my_resource_group*
- Terraform init must be executed first
- Execute terraform import *i.e: terraform import azurerm_resource_group.my_resource_group /subscriptions/{subscription_id}/resourceGroups/{resource_group_name}*
- The manually configured resource block must match current configuration or the resource will be changed with future Terraform deployments

Example of mismatched resource block and existing resource from terraform import

---

```bash

azurerm_resource_group.my_resource_group: Refreshing state... [id=/subscriptions/{subscription_id}/resourceGroups/hms_terraform]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
-/+ destroy and then create replacement

Terraform will perform the following actions:

  # azurerm_resource_group.my_resource_group must be replaced
-/+ resource "azurerm_resource_group" "my_resource_group" {
      ~ id       = "/subscriptions/{subscription_id}/resourceGroups/hms_terraform" -> (known after apply)
      ~ location = "eastus" -> "centralus" # forces replacement
      ~ name     = "hms_terraform" -> "terraform-rg" # forces replacement
      - tags     = {} -> null

      - timeouts {}
    }

Plan: 1 to add, 0 to change, 1 to destroy.
```


### Usage

Managing configuration drift

- `terraform refresh`: This command can be used to alter the state file to match what is currently running
- `terraform plan -out="./plan.json"`: This command can be used to check the state file against what is currently in the terraform configuration files
- `terraform show -json "plan.json" >> ./readable_plan.json`: This command can be used to convert the binary plan to a human-readable json file
- Within the human-readable json file the `actions` object can be inspected to confirm weather changes would be implemented. `no-op` indicates no changes will be implementedcd