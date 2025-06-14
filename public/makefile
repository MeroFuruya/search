NODE_MODULES := node_modules

build:
	if [ ! -d $(NODE_MODULES) ]; then npm ci; fi
	npm run build

.PHONY: build