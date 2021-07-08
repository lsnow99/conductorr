#!/bin/sh
# Build script for Conductorr
# Author Logan Snow
# Last updated 07/07/2021

# Helper functions
mustHaveInstalled() {
    if ! builtin type -P "$1" &> /dev/null
    then
        echo "build process requires $1"
        exit 1
    fi
}

# Check for build requirements
mustHaveInstalled "yarn"
mustHaveInstalled "go"
mustHaveInstalled "brotli"

mv="1.16.0"
v=`go version | { read _ _ v _; echo ${v#go}; }`

if ! go run ./cmd/semcmp $mv $v &> /dev/null
then
    echo "build requires go version > $mv but version $v is installed"
    exit 1
fi

# Build CSL
bash build_csl.sh

# Build web frontend
yarn --cwd ./web install
yarn --cwd ./web build

# Compile for Windows
GOOS=windows GOARCH=386 go build -o bin/conductorr-windows_x86.exe -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
GOOS=windows GOARCH=amd64 go build -o bin/conductorr-windows_x64.exe -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

# Compile for OS X
GOOS=darwin GOARCH=amd64 go build -o bin/conductorr-osx_amd64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
GOOS=darwin GOARCH=arm64 go build -o bin/conductorr-osx_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

# Compile for Linux
GOOS=linux GOARCH=386 go build -o bin/conductorr-linux_x86 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
GOOS=linux GOARCH=amd64 go build -o bin/conductorr-linux_x64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
GOOS=linux GOARCH=arm go build -o bin/conductorr-linux_arm -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
GOOS=linux GOARCH=arm64 go build -o bin/conductorr-linux_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
