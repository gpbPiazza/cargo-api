#!/bin/bash

# Get the directory of the current script
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
# PROJECT_ROOT_DIR=$(dirname "$SCRIPT_DIR")

echo "* Generating public mock *"

# Generate public mock using mockgen
mockgen\
  -source=$SCRIPT_DIR/../interface.go \
  -destination=$SCRIPT_DIR/../mock_module.go \
  -self_package=github.com/olaisaac/backoffice-api-go/modules/charge \
  -package charge

# Replace the absolute path with the Go import path in the source comment
sed -i.bak "s|$SCRIPT_DIR/..|github.com/olaisaac/backoffice-api-go/modules/charge|g" "$SCRIPT_DIR/../mock_module.go"
rm "$SCRIPT_DIR/../mock_module.go.bak"

echo "* Generating private mocks *"

# Find all .go files in the parent directory, excluding interface.go
find $SCRIPT_DIR/.. -type f -name "*.go" ! -path "$SCRIPT_DIR/../interface.go"| while read -r file; do
  # If the file contains an interface definition
  if grep -qE '^\s*type\s+.+\s+interface' "$file"; then
    # Define the mock file path
    mock_file="$SCRIPT_DIR/../temp_mock_$(basename "$file")"
    # Generate mock for the current file
    mockgen -source="$file" -destination="$mock_file" -package "charge" -self_package=github.com/olaisaac/backoffice-api-go/modules/charge
    # Replace 'Mockinterface' with 'mockInterface' in the mock file
    perl -pe 's/Mock([a-z])/mock\u\1/g' "$mock_file" > temp && mv temp "$mock_file"
    # Replace 'Newmock' with 'newMock' in the mock file
    sed -i.bak 's/Newmock/newMock/g' "$mock_file"
    # Remove the package declaration from the mock file
    sed -i.bak '/^package charge/d' "$mock_file"
    # Remove the import block from the mock file
    sed -i.bak '/^import (/,/^)/d' "$mock_file"
    # Remove the source comment from the mock file
    sed -i.bak '/^\/\/ Source:/d' "$mock_file"
    # Remove the backup file created by sed
    rm "${mock_file}.bak"
    echo "Generated mock for: $file"
  fi
done

# Create the final mocks.go file with a package declaration and an empty import block
echo "package charge" > $SCRIPT_DIR/../mocks.go
echo "" >> $SCRIPT_DIR/../mocks.go
echo "import (" >> $SCRIPT_DIR/../mocks.go
echo ")" >> $SCRIPT_DIR/../mocks.go

# Concatenate all temporary mock files into the final mocks.go file
cat $SCRIPT_DIR/../temp_mock* >> $SCRIPT_DIR/../mocks.go

# Check if goimports is installed
PATH=$PATH:$(go env GOPATH)/bin
if ! command -v goimports &> /dev/null
then
    # If not, install goimports
    echo "goimports could not be found, installing..."
    GO111MODULE=off go get golang.org/x/tools/cmd/goimports
fi

# Format the final mocks.go file and organize the import statements
goimports -w $SCRIPT_DIR/../mocks.go

# Remove all temporary mock files
rm -rf $SCRIPT_DIR/../temp_mock*