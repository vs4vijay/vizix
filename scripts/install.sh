#!/usr/bin/env bash
# vizix installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/vs4vijay/vizix/develop/scripts/install.sh | bash

PROJECT="vizix"
VERSION="1.0.0"
GITHUB_URL="https://github.com/vs4vijay/vizix"
BREW_TAP="vs4vijay/vizix/vizix"

BREW=$(command -v brew)

set -e

function copy_binary() {
  name="$1"
  if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
      mv "$name" "$HOME/.local/bin/$name"
  else
      echo "Installing $name to /usr/local/bin which is write protected"
      echo "If you'd prefer to install $name without sudo permissions, add \$HOME/.local/bin to your \$PATH and rerun the installer"
      sudo mv "$name" "/usr/local/bin/$name"
  fi
}

function install_project() {
  if [[ "$OSTYPE" == "linux-gnu" ]]; then
      set -x
      curl -fsSL "${GITHUB_URL}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_linux_x86_64.tar.gz" | tar -xzv "${PROJECT}"
      copy_binary "${PROJECT}"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
      if [[ "${BREW}" != "" ]] && [[ "${BREW_TAP}" != "" ]]; then
          set -x
          brew install "${BREW_TAP}"
      else
          set -x
          curl -fsSL "${GITHUB_URL}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_mac_x86_64.tar.gz" | tar -xzv "${PROJECT}"
          copy_binary "${PROJECT}"
      fi
  else
      set +x
      echo "${PROJECT} does not work for your platform: $OS"
      exit 1
  fi
}

function install() {
  INSTALL_PATH=$(command -v "${PROJECT}" 2>&1 || true)

  if [[ -z ${INSTALL_PATH} ]]; then
    echo "Installing ${PROJECT}"
    install_project
  else
    echo "${PROJECT} already installed, Please type '${PROJECT}' for details"
  fi
}

install
