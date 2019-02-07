release:
	mkdir -p dist 
	go build -o ./dist/debug-aws
	cp ./dist/debug-aws /usr/local/bin/debug-aws