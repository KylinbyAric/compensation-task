all: install

install: clean  build
	mkdir -p output/bin

clean:
	rm -rf output/bin
	rm -rf output/conf
	rm -rf output/control.sh
	rm -rf output/script

build: glide_install;
ifeq ($(COV),yes) 
	goc build
else
	go build
endif

glide_install:
	#glide install
	go clean -modcache

go_install:
	go build

run:
	./output/bin/compensation-task # 前台运行,方便退出
