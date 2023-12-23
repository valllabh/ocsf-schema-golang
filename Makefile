
# prepare system for build
prepare:
	@echo "Preparing system for build..."

	sudo apt-get update

	# install protobuf-compiler
	sudo apt-get install -y protobuf-compiler
	export PATH="$PATH:$(go env GOPATH)/bin"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --version

	# install jq
	sudo apt-get install -y jq

	# install ocsf-tool
	make install-ocsf-tool

	# install buf
	make install-buf

	@echo "System is ready for build."

# install homebrew
install-homebrew:
	@echo "Installing homebrew..."
	curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash

install-buf: install-homebrew
	@echo "Installing buf..."
	/home/linuxbrew/.linuxbrew/bin/brew install bufbuild/buf/buf

# install ocsf-tool
install-ocsf-tool:
	@echo "Installing latest version of ocsf-tool..."
	curl -sfL https://raw.githubusercontent.com/valllabh/ocsf-tool/main/download/download.sh | bash

create-proto-files:
	@echo "Getting event list..."
	./ocsf-tool/ocsf-tool schema-class-list --output ./schema/schema-class-list.json

	@echo "Creating proto files..."
	jq -r "[.[] | to_entries[] | .value] | join(\" \")" ./schema/schema-class-list.json | xargs ./ocsf-tool/ocsf-tool generate-proto --golang-root-package / --proto-output ./proto

compile-proto: clean-output
	@echo "Compiling proto files..."
	buf generate

clean-output:
	@echo "Cleaning output directory..."
	rm -rf ./ocsf

build: create-proto-files compile-proto
	@echo "Building Finished."