#!/usr/bin/env bash
# Installs the papercuts CLI from the latest GitHub release.
#
#   curl -fsSL https://raw.githubusercontent.com/jergason/papercuts/main/install.sh | sh
#
# Override the install directory with INSTALL_DIR (default: $HOME/.local/bin).

set -euo pipefail

REPO="jergason/papercuts"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

os() {
  case "$(uname -s)" in
    Linux) echo "linux" ;;
    Darwin) echo "darwin" ;;
    *) echo "unsupported OS: $(uname -s)" >&2; exit 1 ;;
  esac
}

arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    *) echo "unsupported architecture: $(uname -m)" >&2; exit 1 ;;
  esac
}

OS="$(os)"
ARCH="$(arch)"

echo "Fetching latest papercuts release..." >&2
LATEST_TAG="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')"
if [ -z "${LATEST_TAG}" ]; then
  echo "could not determine latest release tag" >&2
  exit 1
fi

ARCHIVE="papercuts_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ARCHIVE}"
CHECKSUMS_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/checksums.txt"

WORKDIR="$(mktemp -d)"
trap 'rm -rf "${WORKDIR}"' EXIT

echo "Downloading ${URL}..." >&2
curl -fsSL "${URL}" -o "${WORKDIR}/${ARCHIVE}"
curl -fsSL "${CHECKSUMS_URL}" -o "${WORKDIR}/checksums.txt"

echo "Verifying checksum..." >&2
EXPECTED="$(grep " ${ARCHIVE}\$" "${WORKDIR}/checksums.txt" | awk '{print $1}')"
if [ -z "${EXPECTED}" ]; then
  echo "no checksum entry found for ${ARCHIVE}" >&2
  exit 1
fi
if command -v sha256sum >/dev/null 2>&1; then
  ACTUAL="$(sha256sum "${WORKDIR}/${ARCHIVE}" | awk '{print $1}')"
else
  ACTUAL="$(shasum -a 256 "${WORKDIR}/${ARCHIVE}" | awk '{print $1}')"
fi
if [ "${EXPECTED}" != "${ACTUAL}" ]; then
  echo "checksum mismatch: expected ${EXPECTED}, got ${ACTUAL}" >&2
  exit 1
fi

tar -xzf "${WORKDIR}/${ARCHIVE}" -C "${WORKDIR}"

mkdir -p "${INSTALL_DIR}"
mv "${WORKDIR}/papercuts" "${INSTALL_DIR}/papercuts"
chmod +x "${INSTALL_DIR}/papercuts"

echo "Installed papercuts to ${INSTALL_DIR}/papercuts" >&2
if ! command -v papercuts >/dev/null 2>&1; then
  echo "Note: ${INSTALL_DIR} is not on your PATH. Add it, e.g.:" >&2
  echo "  export PATH=\"${INSTALL_DIR}:\$PATH\"" >&2
fi
