$url = "https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-win64.zip"
$destinationPath = "./protoc-26.1-win64.zip"
$extractPath = "./protoc"

if (-Not (Test-Path -Path $extractPath)) {
    Invoke-WebRequest -Uri $url -OutFile $destinationPath
    New-Item -ItemType Directory -Path $extractPath
    Expand-Archive -Path $destinationPath -DestinationPath $extractPath -Force
}
Write-Output "complete: $extractPath"

$sourceFile = $extractPath + "/bin/protoc.exe"
$gobin = [System.Environment]::GetEnvironmentVariable("GOBIN", [System.EnvironmentVariableTarget]::User)
if ([string]::IsNullOrEmpty($gobin)){
    $gopath = [System.Environment]::GetEnvironmentVariable("GOPATH", [System.EnvironmentVariableTarget]::User)
    if (-not $gopath) {
        Write-Error "GOPATH no set"
        exit 1
    }
    $gobin = Join-Path -Path $gopath -ChildPath "bin"
    [System.Environment]::SetEnvironmentVariable("GOBIN", $gobin, [System.EnvironmentVariableTarget]::User)
}
Write-Output "gobin: $gobin"
Copy-Item -Path $sourceFile -Destination $gobin

Write-Output "clean up"
Remove-Item -Path $extractPath -Recurse -Force
Remove-Item -Path $destinationPath -Force
