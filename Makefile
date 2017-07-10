.PHONY : all
all : build

.PHONY : build
build:
	go build -o build/svt github.com/hongkailiu/svt-go

.PHONY : clean
clean:
	rm -rf ./build/*

.PHONY : test
test:
	./script/ci/test_coverage.sh

.PHONY : coveralls
coveralls:
	./script/ci/coveralls.sh

.PHONY : package
package: build
	./script/ci/package.sh

.PHONY : release
release: package
	./script/ci/release.sh

.PHONY : godep_save
godep_save:
	godep save ./...