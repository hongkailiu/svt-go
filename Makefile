.PHONY : all
all : build

.PHONY : build
build:
	go build -o build/svt github.com/hongkailiu/svt-go/main

.PHONY : clean
clean:
	rm -rf ./build/*

.PHONY : test
test:
	./script/ci/test_coverage.sh

.PHONY : coveralls
coveralls:
	./script/ci/coveralls.sh
