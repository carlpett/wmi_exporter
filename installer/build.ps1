[CmdletBinding()]
Param (
  [Parameter(Mandatory=$true)]
  [String] $Version,
  [Parameter(Mandatory=$false)]
  [ValidateSet("amd64","386")]
  [String] $Arch = "amd64"
)
$ErrorActionPreference = "Stop"

if ($PSVersionTable.PSVersion.Major -lt 5) {
  Write-Error "Powershell version 5 required"
  exit 1
}

$wc = New-Object System.Net.WebClient
function Get-FileIfNotExists {
  Param (
    $Url,
    $Destination
  )
  if(-not (Test-Path $Destination)) {
    Write-Verbose "Downloading $Url"
    $wc.DownloadFile($Url, $Destination)
  }
  else {
    Write-Verbose "${Destination} already exists. Skipping."
  }
}

$sourceDir = mkdir -Force Source
mkdir -Force Work,Output | Out-Null

Write-Verbose "Downloading files"
# Somewhat obscure url, points to WiX 3.10 binary release
Write-Verbose "Downloading WiX..."
Get-FileIfNotExists "http://download-codeplex.sec.s-msft.com/Download/Release?ProjectName=wix&DownloadId=1504735&FileTime=130906491728530000&Build=21031" "$sourceDir\wix-binaries.zip"
mkdir -Force WiX | Out-Null
Expand-Archive -Path "${sourceDir}\wix-binaries.zip" -DestinationPath WiX -Force

Write-Verbose "Downloading wmi_exporter..."
Get-FileIfNotExists "https://github.com/martinlindhe/wmi_exporter/releases/v${Version}_${Arch}.zip" "${sourceDir}\wmi_exporter-${Version}_${Arch}.zip"

Write-Verbose "Unpacking"
Expand-Archive "${sourceDir}\wmi_exporter-${Version}_${Arch}.zip" -DestinationPath Work -Force

Write-Verbose "Creating wmi_exporter-${Version}-${Arch}.msi"
$wixArch = @{"amd64"="x64"; "386"="x86"}[$Arch]
$WixOpts = ""
Invoke-Expression "WiX\candle.exe -nologo -arch $wixArch $WixOpts -out Work\wmi_exporter.wixobj -dVersion=`"$Version`" wmi_exporter.wxs"
Invoke-Expression "WiX\light.exe -nologo -spdb $WixOpts -out `"Output\wmi_exporter-${Version}-${Arch}.msi`" Work\wmi_exporter.wixobj"

Write-Verbose "Done!"