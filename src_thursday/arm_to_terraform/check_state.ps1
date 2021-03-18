using namespace System
using namespace System.IO

param(
    [string]$FilePath 
)

Clear-Host

[Console]::WriteLine("Looking at drift for terraform file: $FilePath`n")


$terraformPlan = ConvertFrom-Json $([File]::ReadAllText($FilePath))

foreach($resource in $terraformPlan.resource_changes) {

    [Console]::WriteLine("Resource type: " + $resource.type)
    [Console]::WriteLine("Resource name: " + $resource.name)
    [Console]::WriteLine("Resource mode: " + $resource.mode + "`n")

    if ($resource.change.actions -eq [string]"no-op") {

        Write-Host "No changes needed" -ForegroundColor Yellow

    } else {

        Write-Host "Changes needed:" -ForegroundColor Red

        foreach($change in $resource.change.actions) {
            Write-Host `t$change
        }

        if ($resource.change.before) {
            Compare-Object -ReferenceObject ($resource.change.before) -DifferenceObject ($resource.change.after) -Property 'location', 'name'
        } else {
            Write-Host $resource.change.after | ConvertTo-Json -AsArray
        }

        # Email can be sent to ops team showing what will be changed
    }
}