PATH := ./node_modules/.bin/:$(PATH)

all:
	browserify --transform babelify assets/src/*.js* --outfile assets/main.js
