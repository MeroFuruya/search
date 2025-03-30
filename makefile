GO_BUILD_FLAGS :=
GO_BUILD_STANDALONE_FLAGS := $(GO_BUILD_FLAGS) -tags=standalone

build-public:
	make -C public build

build:
	go build $(GO_BUILD_FLAGS) -o=build/search .

build-standalone: build-public
	go build $(GO_BUILD_STANDALONE_FLAGS) -o=build/search-standalone .

docker:
	docker build \
	-f Dockerfile \
	-t search:latest \
	.

docker-standalone: build-public
	docker build \
	-f Dockerfile.standalone \
	-t search:latest-standalone \
	.

standalone: build-standalone
.PHONY: build
