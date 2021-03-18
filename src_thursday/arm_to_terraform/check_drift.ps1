using namespace System.IO

terraform refresh
terraform plan -out="./plan.json" | Out-File -FilePath "./PlanOutput.txt"
terraform show -json "plan.json" >> "./readable_plan.json"

./check_state.ps1 -FilePath "./readable_plan.json"
