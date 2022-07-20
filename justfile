set export

green := "tput setaf 2"
blue := "tput setaf 4"
cyan := "tput setaf 6"
reset := "tput sgr0"

# Make sure an executable is available
@check-installed executable:
    if ! type "$executable" > /dev/null; then \
        echo "build process requires $executable"; \
        exit 1; \
    fi

# Check for go executable and acceptable version
@check-go: (check-installed "go")
    mv="1.16.0"; \
    v=`go version | { read _ _ v _; echo ${v#go}; }`; \
    if ! go run ./cmd/semcmp "$mv" "$v" > /dev/null; then \
        echo "build requires go version > $mv but version $v is installed"; \
        exit 1; \
    fi

# Check for brotli executable
@check-brotli: (check-installed "brotli")

# Check for pnpm executable
@check-pnpm: (check-installed "pnpm")

# Check for required software
@check-all-deps: check-go check-pnpm check-brotli

# Install web dependencies
@install-web: check-pnpm
    pnpm install
    cd frontend && pnpm install
    cd docusite && pnpm install

# Build CSL WASM module
@build-csl-wasm: check-go check-brotli
    echo "==============================[    Building CSL WASM    ]=============================="
    echo "`${green}`→ Compiling to WASM`${reset}`"
    GOOS=js GOARCH=wasm go build -o dist/csl.wasm -ldflags="-s -w -X 'main.CorsProxyServer=/.netlify/functions/corsproxy'" ./cmd/cslwasm
    echo "`${cyan}`→ Compressing with brotli`${reset}`"
    brotli -f dist/csl.wasm
    # Copy the corresponding wasm_exec.js file to the frontend app and docusite projects
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./frontend/src/util
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./docusite/docs/public
    # Copy the web assembly modules so that they are accessible for the docusite
    cp ./dist/csl.wasm ./docusite/docs/public
    cp ./dist/csl.wasm.br ./docusite/docs/public

# Build the shared library
@build-shared: install-web
    cd shared && \
    pnpm build

# Build docusite
@build-docusite: install-web build-csl-wasm
    cd docusite && \
    pnpm build

# Build web frontend
@build-frontend: install-web build-csl-wasm
    cd frontend && \
    pnpm build

# Compile the CSL command line interface
@build-csl-cli: check-go
    echo "==============================[    Compiling CSL CLI    ]=============================="
    echo "`${cyan}`→ Compiling`${reset}`"
    go build -o dist/csl ./cmd/csl

# Compile the corsproxy lambda function
@build-corsproxy: check-go
    echo "==============================[   Compiling CORS Proxy  ]=============================="
    echo "`${cyan}`→ Compiling`${reset}`"
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/functions/corsproxy -ldflags="-s -w" ./cmd/corsproxy

# Build the binaries
@build-bin: check-go
    echo "==============================[   Compiling Conductorr  ]=============================="
    # Compile for Windows
    echo "`${cyan}`→ Compiling for Windows`${reset}`"
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/conductorr-windows_x64.exe -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

    # Compile for OS X
    echo "`${cyan}`→ Compiling for OS X`${reset}`"
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/conductorr-osx_amd64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/conductorr-osx_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

    # Compile for Linux
    echo "`${cyan}`→ Compiling for Linux`${reset}`"
    CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/conductorr-linux_x86 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/conductorr-linux_x64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/conductorr-linux_arm -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/conductorr-linux_arm64 -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" ./cmd/conductorr

# Build everything
@build: build-csl-wasm build-csl-cli build-corsproxy build-bin

# Run the conductorr frontend in dev mode
@run-frontend: install-web
    cd frontend && \
    pnpm dev

# Run the conductorr docusite in dev mode
@run-docusite: install-web
    cd docusite && \
    pnpm dev

# Run the conductorr shared library in dev mode
@run-shared: install-web
    cd shared && \
    pnpm dev
