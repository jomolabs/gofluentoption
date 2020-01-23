build:
	hack/build.sh

release:
	hack/release.sh

clean:
	rm -f bin/goption

.PHONY: release
