New-Item -ItemType Directory -Force -Path dist
foreach($platform in (go tool dist list)) {
    $os, $arch = $platform -split '/'
    $env:GOOS = $os
    $env:GOARCH = $arch
    go build -o "dist/app-$os-$arch$(if ($os -eq 'windows') {'.exe'})"
}