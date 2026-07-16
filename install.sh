#!/bin/sh
# shellcheck shell=dash
# Install akavecli — the Akave SDK CLI
#
# Usage:
#   curl -sSL https://raw.githubusercontent.com/akave-ai/akavesdk/main/install.sh | bash
#
# Overrides (env vars):
#   AKAVECLI_VERSION=v0.4.4   install a specific version
#   INSTALL_DIR=/usr/local/bin install location (default: /usr/local/bin)
#   GITHUB_ORG=akave-ai        GitHub org (override for forks)

set -eu

GITHUB_ORG=${GITHUB_ORG:-"akave-ai"}
GITHUB_REPO=${GITHUB_REPO:-"akavesdk"}
BINARY_NAME="akavecli"
INSTALL_DIR=${INSTALL_DIR:-"/usr/local/bin"}
HTTP_CLI=${HTTP_CLI:-"curl"}
version=${AKAVECLI_VERSION:-""}

get_latest_release() {
    local release_url="https://api.github.com/repos/${GITHUB_ORG}/${GITHUB_REPO}/releases/latest"

    if [ "$HTTP_CLI" = "curl" ]; then
        version=$(curl -s "$release_url" | grep '"tag_name"' | grep -Eo '"v[^"]+"' | head -1 | tr -d '"')
    else
        version=$(wget -q -O - "$release_url" | grep '"tag_name"' | grep -Eo '"v[^"]+"' | head -1 | tr -d '"')
    fi

    if [ -z "$version" ]; then
        echo "Error: could not determine latest release." >&2
        echo "Set AKAVECLI_VERSION to install a specific version, e.g.:" >&2
        echo "  AKAVECLI_VERSION=v0.4.4 curl -sSL ... | bash" >&2
        exit 1
    fi

    echo "Latest release: $version"
}

detect_platform() {
    ARCH=$(uname -m)
    case "$ARCH" in
        x86_64|amd64)  ARCH="amd64" ;;
        aarch64|arm64) ARCH="arm64" ;;
        *)
            echo "Unsupported architecture: $ARCH" >&2
            exit 1
            ;;
    esac

    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    case "$OS" in
        linux|darwin) ;;
        *)
            echo "Unsupported OS: $OS" >&2
            echo "For Windows, download from: https://github.com/${GITHUB_ORG}/${GITHUB_REPO}/releases/latest" >&2
            exit 1
            ;;
    esac

    export ARCH OS
}

install_akavecli() {
    if [ -z "$version" ]; then
        get_latest_release
    fi

    local asset="${BINARY_NAME}-${OS}-${ARCH}"
    local url="https://github.com/${GITHUB_ORG}/${GITHUB_REPO}/releases/download/${version}/${asset}"
    local tmp
    tmp=$(mktemp)

    echo "Downloading ${BINARY_NAME} ${version} (${OS}/${ARCH})..."

    if [ "$HTTP_CLI" = "curl" ]; then
        curl -fsSL -o "$tmp" "$url" || { echo "Download failed: $url" >&2; rm -f "$tmp"; exit 1; }
    else
        wget -q -O "$tmp" "$url" || { echo "Download failed: $url" >&2; rm -f "$tmp"; exit 1; }
    fi

    chmod +x "$tmp"

    if [ -w "$INSTALL_DIR" ]; then
        mv "$tmp" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        echo "Requires elevated permissions to install to ${INSTALL_DIR}..."
        sudo mv "$tmp" "${INSTALL_DIR}/${BINARY_NAME}"
    fi

    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        echo ""
        echo "akavecli installed to ${INSTALL_DIR}/${BINARY_NAME}"
        echo "Run 'akavecli --help' to get started."
        echo "Docs: https://docs.akave.ai"
    else
        echo "Installed to ${INSTALL_DIR}/${BINARY_NAME}. Ensure ${INSTALL_DIR} is in your PATH."
    fi
}

detect_platform
install_akavecli
