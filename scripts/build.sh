#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Determine the arch/os combos we're building for
XC_OS=$(go env GOOS)
XC_ARCH=$(go env GOARCH)

# Check flags for cross-compiling
if [[ -n "${FULL_BUILD}" ]]; then
    XC_ARCH="386 amd64 arm"
    XC_OS="linux darwin windows freebsd openbsd solaris"
    XC_EXCLUDE_OSARCH="!darwin/arm !darwin/386"
fi

# link staically
GO_TAGS="";
if [[ -n "${STATIC_LINK}" ]]; then
    GO_TAGS="${GO_TAGS} static"
fi

# Instruct gox to build statically linked binaries
export CGO_ENABLED=1

# Set module download mode to readonly to not implicitly update go.mod
export GOFLAGS="-mod=readonly"

# # In release mode we don't want debug information in the binary
if [[ -n "${PESCA_RELEASE}" ]]; then
    LD_FLAGS="-s -w"
else 
    # In dev mode, LD_FLAGS to be appended during development compilations
    LD_FLAGS="-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY} $LD_FLAGS"
fi

# Delete the old dir
echo "==> Removing old directory..."
rm -rf bin/*
rm -rf pkg/*
mkdir -p bin/

# install gox
if ! which gox > /dev/null; then
    echo "==> Installing gox..."
    go get github.com/mitchellh/gox
fi

# Ensure all remote modules are downloaded and cached before build so that
# the concurrent builds launched by gox won't race to redundantly download them.
go mod download

# Create GOPATH/bin if it's doesn't exists
if [ ! -d $MAIN_GOPATH/bin ]; then
    echo "==> Creating GOPATH/bin directory..."
    mkdir -p $MAIN_GOPATH/bin
fi

# Build!
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -ldflags "${LD_FLAGS}" \
    -tags "${GO_TAGS}" \
    -output "bin/{{.OS}}_{{.Arch}}/go-bf" \
    ./cmd/go-bf

# Move all the compiled things to the $GOPATH/bin
GOPATH=${GOPATH:-$(go env GOPATH)}
case $(uname) in
    CYGWIN*)
        GOPATH="$(cygpath $GOPATH)"
        ;;
esac
OLDIFS=$IFS
IFS=: MAIN_GOPATH=($GOPATH)
IFS=$OLDIFS

# Done!
echo
echo "==> Results:"
ls -hlR bin/
