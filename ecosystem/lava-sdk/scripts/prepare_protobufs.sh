#!/bin/bash

function prepare() {
    echo "🔹make sure to run 'go mod tidy' from the lava repo before trying to run this file"
    
    use_sudo=$1
    if [ "$use_sudo" = true ]; then
        SUDO=sudo
    else
        SUDO=''
    fi

    file_path="../../go.mod"
    expected_lines=(
        "github.com/gogo/googleapis v1.4.1 // indirect"
        "github.com/cosmos/cosmos-sdk v0.47.3"
        "github.com/cosmos/gogoproto v1.4.10"
        "github.com/cosmos/cosmos-proto v1.0.0-beta.2"
    )

    missing_lines=()

    for line in "${expected_lines[@]}"; do
        if ! grep -qF "$line" "$file_path"; then
            missing_lines+=("$line")
        fi
    done

    if [[ ${#missing_lines[@]} -eq 0 ]]; then
        echo "✨ All expected lines are present in the $file_path file."
    else
        echo "Some or all expected lines are missing in the $file_path file."
        echo "Missing lines:"
        for missing_line in "${missing_lines[@]}"; do
            echo "🔹$missing_line"
        done
        exit 1
    fi

    gopath=$(go env GOPATH)
    if [[ -z "$gopath" ]]; then
        echo "Error: GOPATH is not set. setting it to ~/go" >&2
    fi

    if [[ ! -d "$gopath" ]]; then
        echo "Error: The directory specified in GOPATH ('$gopath') does not exist." >&2
        exit 1
    fi

    specific_dir="$gopath/pkg/mod/github.com/lavanet/cosmos-sdk@v0.47.7-0.20231211141641-2a9ea55b724d"

    if [[ ! -d "$specific_dir" ]]; then
        echo "Error: The cosmos-sdk directory ('$specific_dir') does not exist under '$gopath/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    gogodir="$gopath/pkg/mod/github.com/cosmos/gogoproto@v1.4.10"

    if [[ ! -d "$gogodir" ]]; then
        echo "Error: The gogodir directory ('$gogodir') does not exist under '$gopath/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    googledir="$gopath/pkg/mod/github.com/gogo/googleapis@v1.4.1"

    if [[ ! -d "$googledir" ]]; then
        echo "Error: The googledir directory ('$googledir') does not exist under '$gopath/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    cosmosprotosdir="$gopath/pkg/mod/github.com/cosmos/cosmos-proto@v1.0.0-beta.2"

    if [[ ! -d "$cosmosprotosdir" ]]; then
        echo "Error: The cosmosprotosdir directory ('$cosmosprotosdir') does not exist under '$gopath/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    $SUDO rm -rf ./proto

    mkdir -p proto/cosmos/cosmos-sdk/google/api

    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/cosmos; cp -r $specific_dir/proto/cosmos ./proto/cosmos/cosmos-sdk
    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/amino; cp -r $specific_dir/proto/amino ./proto/cosmos/cosmos-sdk
    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/tendermint; cp -r $specific_dir/proto/tendermint ./proto/cosmos/cosmos-sdk
    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/gogoproto; cp -r $gogodir/gogoproto ./proto/cosmos/cosmos-sdk
    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/google; cp -r $gogodir/protobuf/google ./proto/cosmos/cosmos-sdk
    $SUDO rm -rf ./proto/cosmos/cosmos-sdk/cosmos_proto; cp -r $cosmosprotosdir/proto/cosmos_proto ./proto/cosmos/cosmos-sdk
    $SUDO cp -r $googledir/google/api ./proto/cosmos/cosmos-sdk/google

    cp -r ../../proto/lavanet ./proto

    group=$(groups $(whoami) | cut -d' ' -f1)
    $SUDO chown -R $(whoami):$group ./proto
}