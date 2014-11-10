.PHONY: build
build:
	./build.sh
vserver:
	cd site && bundle exec jekyll serve -w
