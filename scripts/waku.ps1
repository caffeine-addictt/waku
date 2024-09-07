if (-not (Get-Command "docker" -ErrorAction SilentlyContinue)) {
  Write-Host "Error: Docker is not installed. Please install Docker and try again."
  exit 1
}

param (
  [string]$registry = "dockerhub",
  [Parameter(ValueFromRemainingArguments=$true)] [string[]]$args
)

# Select the image based on the registry preference
$image = "caffeinec/waku:latest"
if ($registry -eq "ghcr") {
  $image = "ghcr.io/caffeine-addictt/waku:latest"
}

docker run -v "$(Get-Location):/app" $image $args
