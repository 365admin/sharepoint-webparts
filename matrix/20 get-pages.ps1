<#---
title: Get Matrix Pages 
connection: sharepoint
output: matrix-pages.json
---

## A file is extract from SharePoint containing the pages and the metadata needed for a Matrix web

#>
param (
    [string]$site = $env:SITEURL # "https://christianiabpos.sharepoint.com/sites/IssuerProducts"
)
$result = "$env:WORKDIR/matrix-pages.json"

Connect-PnPOnline -Url $site   -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

$listItems = Get-PnpListItem -List SitePages  

write-host "Pages in list: $($listItems.Count)"
$pages = @()
foreach ($item in $listItems) {
    $page = @{
        ID = $item.FieldValues.ID
        Title = $item.FieldValues.Title
        ValueChain = $item.FieldValues.ValueChain
        SortOrder = $item.FieldValues.SortOrder
        FileRef = $item.FieldValues.FileRef
        #URL = $item.FieldValues.
    }
    $pages += $page
   
}

$pages | ConvertTo-Json -Depth 10 | Out-File -FilePath $result -Encoding utf8NoBOM
