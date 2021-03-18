using namespace System
using namespace System.IO

param(
    [string]$FilePath 
)

Clear-Host

[Console]::WriteLine("Looking at drift for terraform file: $FilePath`n")


$terraformPlan = ConvertFrom-Json $([File]::ReadAllText($FilePath))

[Console]::WriteLine("Resource type: " + $terraformPlan.resource_changes.type)
[Console]::WriteLine("Resource name: " + $terraformPlan.resource_changes.name)
[Console]::WriteLine("Resource mode: " + $terraformPlan.resource_changes.mode + "`n")

if ($terraformPlan.resource_changes.change.actions -eq [string]"no-op") {

    Write-Host "No changes needed" -ForegroundColor Yellow

} else {

    Write-Host "Changes needed:" -ForegroundColor Red

    foreach($change in $terraformPlan.resource_changes.change.actions) {
        Write-Host `t$change
    }

    Compare-Object -ReferenceObject ($terraformPlan.resource_changes.change.before) -DifferenceObject ($terraformPlan.resource_changes.change.after) -Property 'location', 'name'

    # Email can be sent to ops team showing what will be changed
}
