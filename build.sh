#!/bin/bash
# Build script for Conductorr
# Author Logan Snow
# Last updated 08/03/2021

# Set script to fail if any sub commands fail
set -e

# Terminal color definitions
green=`tput setaf 2`
blue=`tput setaf 4`
cyan=`tput setaf 6`
reset=`tput sgr0`

# Parse command
buildcsl=0
buildweb=0
buildbin=0

if [ "$1" == "all" ] || [ "$1" == "" ]; then
    buildcsl=1
    buildweb=1
    buildbin=1
fi

if [ "$1" == "csl" ]; then
    buildcsl=1
fi

if [ "$1" == "web" ]; then
    buildweb=1
fi

if [ "$1" == "bin" ]; then
    buildbin=1
fi

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

# Build CSL WASM module
if [ $buildcsl == 1 ]; then
    echo "==============================[   Building CSL Module   ]=============================="
    echo "${cyan}→ Compiling to WASM${reset}"
    GOOS=js GOARCH=wasm go build -o dist/csl.wasm ./cmd/csl
    echo "${cyan}→ Compressing with brotli${reset}"
    brotli -f dist/csl.wasm
    # Copy the corresponding wasm_exec.js file to the web app and docusite projects
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./web/src/util
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./docusite/docs/.vuepress
    # Copy the web assembly modules so that they are accessible for the docusite
    cp ./dist/csl.wasm ./docusite/docs/.vuepress/public
    cp ./dist/csl.wasm.br ./docusite/docs/.vuepress/public
fi

# Build web frontend
if [ $buildweb == 1 ]; then
    echo "==============================[    Building Frontend    ]=============================="
    echo "${cyan}→ Installing frontend dependencies${reset}"
    yarn --cwd ./web install
    echo "${cyan}→ Building frontend for distribution${reset}"
    yarn --cwd ./web build
fi

# Build the full binaries
if [ $buildbin == 1 ]; then
    echo "==============================[ Compiling Full Binaries ]=============================="
    # Compile for Windows
    echo "${cyan}→ Compiling for Windows${reset}"
    GOOS=windows GOARCH=386 go build -o bin/conductorr-windows_x86.exe -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    GOOS=windows GOARCH=amd64 go build -o bin/conductorr-windows_x64.exe -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

    # Compile for OS X
    echo "${cyan}→ Compiling for OS X${reset}"
    GOOS=darwin GOARCH=amd64 go build -o bin/conductorr-osx_amd64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    GOOS=darwin GOARCH=arm64 go build -o bin/conductorr-osx_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

    # Compile for Linux
    echo "${cyan}→ Compiling for Linux${reset}"
    GOOS=linux GOARCH=386 go build -o bin/conductorr-linux_x86 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    GOOS=linux GOARCH=amd64 go build -o bin/conductorr-linux_x64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    GOOS=linux GOARCH=arm go build -o bin/conductorr-linux_arm -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    GOOS=linux GOARCH=arm64 go build -o bin/conductorr-linux_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

    echo "${green}✓ Build succeeded${reset}"
fi
