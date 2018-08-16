BINARY = timechi
SOURCES = $(*.go)
LD_FLAGS = "-X github.com/hamidfzm/timechi-server/cmd.versionPatch=`git rev-list HEAD --count`"
BRANCH=`git rev-parse --abbrev-ref HEAD`

build:
	go build -ldflags $(LD_FLAGS) -o $(BINARY)

build-static:
	CGO_ENABLED=0 GOOS=linux go build -ldflags $(LD_FLAGS) -a -installsuffix cgo -o $(BINARY)

install:
	go install

vendor:
	glide install

clean:
	go clean
	rm -rf $(BINARY)

pull:
	git pull origin $(BRANCH)

push: pull
	git push origin $(BRANCH)

docker-image: push vendor build-static
	docker build -t hamidfzm/timechi-server:$(BRANCH) --rm .

docker: docker-image clean