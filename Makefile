WALLET_CORE_REPO = https://github.com/trustwallet/wallet-core.git
PACKAGE = crypto-api
PROTO_SRC_DIR = ./infrastructure/proto
PROTO_DST_DIR = ./domain

.PHONY: tools
tools:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go


.PHONY: wallet-core
wallet-core:
	@mkdir "wallet-core" || echo "dir already exists"
	@rm -rf /tmp/wc || echo "dir is not exists"
	git clone $(WALLET_CORE_REPO) /tmp/wc
	cd /tmp/wc \
 		&& /tmp/wc/tools/install-dependencies \
 		&& tools/generate-files \
		&& cmake -H. -Bbuild -DCMAKE_BUILD_TYPE=Debug \
		&& make -Cbuild -j12

	cp -r /tmp/wc/build ./wallet-core
	cp -r /tmp/wc/include ./wallet-core
	rm -rf /tmp/wc

.PHONY: proto
proto: tools
	# TODO: fix the Common.proto issue, imported by bitcoin
	protoc -I=$(PROTO_SRC_DIR) --go_out=$(PROTO_DST_DIR) $(PROTO_SRC_DIR)/*

.PHONY: sign_test
sign_test:
	echo "running HTTP tests"
	sh ./tests.sh
	echo "==================="
	echo "running JSON RPC tests"
	go run testing/main.go

.PHONY: build
build:
	if ! [ -d "./wallet-core" ]; then make wallet-core; fi
	make proto
	go build ./cmd/crypto_api/main.go
