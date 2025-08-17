param(
    [switch]$Uninstall
)

$Repo = "Alfakynz/PackWize"
$InstallDir = "$env:USERPROFILE\AppData\Local\Packwize\bin"
$BinPath = "$InstallDir\packwize.exe"

function Update-UserPath($NewPath) {
    [System.Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
}

if ($Uninstall) {
    Write-Host "Uninstalling Packwize..."
    if (Test-Path $BinPath) {
        Remove-Item $BinPath -Force
        Write-Host "Removed $BinPath"
    } else {
        Write-Host "Packwize is not installed in $InstallDir"
    }

    # Remove from PATH if present
    $envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
    if ($envPath -like "*$InstallDir*") {
        $newPath = ($envPath -split ";" | Where-Object {$_ -ne $InstallDir}) -join ";"
        Update-UserPath $newPath
        Write-Host "Removed $InstallDir from PATH"
    }

    exit 0
}

# Detect latest version from GitHub API
Write-Host "Fetching latest version..."
$LatestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
$Version = $LatestRelease.tag_name

if (-not $Version) {
    Write-Host "Could not fetch latest version from GitHub"
    exit 1
}

$Filename = "packwize-windows.zip"
$Url = "https://github.com/$Repo/releases/download/$Version/$Filename"
$Tmp = "$env:TEMP\packwize.zip"
$ExtractDir = "$env:TEMP\packwize"

Write-Host "Downloading $Url..."
Invoke-WebRequest -Uri $Url -OutFile $Tmp -UseBasicParsing

Write-Host "Extracting..."
if (Test-Path $ExtractDir) { Remove-Item $ExtractDir -Recurse -Force }
Expand-Archive -Path $Tmp -DestinationPath $ExtractDir -Force

Write-Host "Installing to $InstallDir"
if (!(Test-Path $InstallDir)) { New-Item -ItemType Directory -Path $InstallDir | Out-Null }
Move-Item -Force "$ExtractDir\packwize.exe" $BinPath

# Add to PATH for current user if not already present
$envPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($envPath -notlike "*$InstallDir*") {
    Write-Host "Adding $InstallDir to PATH..."
    [System.Environment]::SetEnvironmentVariable(
        "Path",
        $envPath + ";" + $InstallDir,
        [System.EnvironmentVariableTarget]::User
    )
}

Write-Host "Installation complete!"
Write-Host "Open a new terminal and run: packwize --help"
