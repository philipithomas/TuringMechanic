.PHONY: build
build:
	./build.sh
vserver:
	cd site && bundle exec jekyll serve -w
pretty:
	cd go && bash prettify.sh
