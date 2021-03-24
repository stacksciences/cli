REV != git rev-parse --short HEAD 2> /dev/null || echo 'unknown'
SHA1 != git rev-parse HEAD 2> /dev/null || echo 'unknow'
BUILD_BRANCH != git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown'
BUILD_DATE != date +%Y%m%d-%H:%M:%S

GOBUILDFLAGS=-ldflags "-s -w -X stctl/cmd.Version=$(SHA1) \
								-X stctl/cmd.GitTag=$(REV) \
	 							-X stctl/cmd.BuildDate=$(BUILD_DATE)"

all:
	go build $(GOBUILDFLAGS) -o stctl

clean:
	rm -rf stctl
