TAGS = 

ifeq ($(TAG),)
	TAGS=v0.0.1
else
	TAGS=${TAG}
endif

.PHONY: clean binary docker archive

default: clean binary

image: clean binary docker

binary:
	echo "Building go-logstash-pusher binary......"
	@CGO_ENABLED=0 \
		go build -v -ldflags "-X go-logstash/config.VERSION=${TAGS} -s -w" \
		-o go-logstash-pusher ./main.go

docker:
	cp go-logstash-pusher images/go-logstash-pusher
	@docker build \
		--no-cache=true \
		-f images/go-logstash-pusher/Dockerfile \
		-t utilities/go-logstash-pusher:${TAGS} \
		images/go-logstash-pusher
	rm -rf images/go-logstash-pusher/go-logstash-pusher go-logstash-pusher

archive:
	docker save `docker images -a --format {{.Repository}}:{{.Tag}} | grep -E "utilities/(go-logstash-pusher)"` | gzip > go-logstash-pusher-${TAGS}.tgz

export: image archive

clean:
	rm -rf go-logstash-pusher *.tgz
	docker images -a --format {{.Repository}}:{{.Tag}} | grep -E "utilities/(go-logstash-pusher)" | xargs -i docker rmi {} >/dev/null 2>&1
	docker images -a --format {{.Repository}}:{{.ID}} | grep "none" | awk -F":" '{print $$2}' | xargs -i docker rmi {} >/dev/null 2>&1 || true
