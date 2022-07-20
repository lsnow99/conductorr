set export

@check-installed executable:
    if ! type "$executable" > /dev/null; then \
        echo "build process requires $executable"; \
        exit 1; \
    fi

@install-web:
    pnpm install

@build-shared: install-web
    cd shared
    pnpm build

@build-docusite: install-web build-shared
    cd docusite
    pnpm build

@build-frontend: install-web build-shared
    cd frontend
    pnpm build

@colors:
    if tty -s; then; \
        green=`tput setaf 2`; \
        blue=`tput setaf 4`; \
        cyan=`tput setaf 6`; \
        reset=`tput sgr0`; \
    fi

@build-wasm: colors
    echo "==============================[    Building CSL WASM    ]=============================="
    echo "${cyan}→ Compiling to WASM${reset}"
    GOOS=js GOARCH=wasm go build -o dist/csl.wasm -ldflags="-s -w -X 'main.CorsProxyServer=/.netlify/functions/corsproxy'" ./cmd/cslwasm
    echo "${cyan}→ Compressing with brotli${reset}"
    brotli -f dist/csl.wasm
    # Copy the corresponding wasm_exec.js file to the frontend app and docusite projects
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./frontend/src/util
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./docusite/docs/js
    # Copy the web assembly modules so that they are accessible for the docusite
    cp ./dist/csl.wasm ./docusite/docs/js
    cp ./dist/csl.wasm.br ./docusite/docs/js