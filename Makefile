PATH := ./node_modules/.bin/:$(PATH)

all:
	browserify --transform babelify assets/src/*.js* --outfile assets/main.js
	go-bindata assets/
	go build

depends:
	go get -u github.com/jteeuwen/go-bindata/...
	npm i
