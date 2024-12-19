.PHONY: download generate



generate:
	go generate

# specification for the listenbrainz api provided by rain0r
# https://github.com/rain0r/listenbrainz-openapi/tree/master
download:
	curl https://rain0r.github.io/listenbrainz-openapi/openapi.yaml > lb.yaml
