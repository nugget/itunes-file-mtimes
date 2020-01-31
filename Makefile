macos:
	env GOOS=darwin GOARCH=amd64 go build -o dist/itunes-file-mtimes-macos

windows:
	env GOOS=windows GOARCH=amd64 go build -o dist/itunes-file-mtimes-win64

linux:
	env GOOS=linux GOARCH=amd64 go build -o dist/itunes-file-mtimes-linux

dist: linux macos windows
