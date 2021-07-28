$WALLET_CORE_REPO="https://github.com/trustwallet/wallet-core.git"

.PHONY: wallet-core
wallet-core:
	rm -rf /tmp/wc || echo "dir is not exists"
	git clone $(WALLET_CORE_REPO) /tmp/wc
	cd /tmp/wc && sh /tmp/wc/bootstrap.sh
	cp -r /tmp/wc/build ./wallet-core
	cp -r /tmp/wc/include ./wallet-core
	rm -rf /tmp/wc
