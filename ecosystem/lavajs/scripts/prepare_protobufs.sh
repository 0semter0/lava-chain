#!/bin/bash

function prepare() {
    echo "🔹make sure to run 'go mod tidy' from the lava repo before trying to run this file"

    file_path="../../go.mod"
    expected_lines=(
        "github.com/gogo/googleapis v1.4.1 // indirect"
        "github.com/cosmos/cosmos-sdk v0.47.3"
        "github.com/cosmos/gogoproto v1.4.10"
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


    if [[ -z "$GOPATH" ]]; then
        echo "Error: GOPATH is not set. Set the GOPATH environment variable to your Go workspace directory." >&2
        exit 1
    fi

    if [[ ! -d "$GOPATH" ]]; then
        echo "Error: The directory specified in GOPATH ('$GOPATH') does not exist." >&2
        exit 1
    fi

    specific_dir="$GOPATH/pkg/mod/github.com/cosmos/cosmos-sdk@v0.47.3"

    if [[ ! -d "$specific_dir" ]]; then
        echo "Error: The cosmos-sdk directory ('$specific_dir') does not exist under '$GOPATH/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    gogodir="$GOPATH/pkg/mod/github.com/cosmos/gogoproto@v1.4.10"

    if [[ ! -d "$gogodir" ]]; then
        echo "Error: The gogodir directory ('$gogodir') does not exist under '$GOPATH/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    googledir="$GOPATH/pkg/mod/github.com/gogo/googleapis@v1.4.1"

    if [[ ! -d "$googledir" ]]; then
        echo "Error: The googledir directory ('$googledir') does not exist under '$GOPATH/pkg/mod'." >&2
        echo "make sure you ran 'go mod tidy' in the lava main repo"
        exit 1
    fi

    sudo rm -rf ./proto/cosmos; cp -r $specific_dir/proto/cosmos ./proto
    sudo rm -rf ./proto/amino; cp -r $specific_dir/proto/amino ./proto
    sudo rm -rf ./proto/tendermint; cp -r $specific_dir/proto/tendermint ./proto
    sudo rm -rf ./proto/gogoproto; cp -r $gogodir/gogoproto ./proto
    sudo rm -rf ./proto/google; cp -r $gogodir/protobuf/google ./proto
    sudo mkdir ./proto/google/api
    sudo cp -r $googledir/google/api/annotations.proto ./proto/google/api/.
    sudo cp -r $googledir/google/api/http.proto ./proto/google/api/.

    sudo chown -R $(whoami):$(whoami) ./proto
}