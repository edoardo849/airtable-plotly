
BINARY=api

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/edoardo849/airtable-ideas/api.Version=${VERSION} -X github.com/edordo849/airtable-ideas/api.BuildTime=${BUILD_TIME}"

all: clean
	go build ${LDFLAGS} -o ${BINARY} ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
