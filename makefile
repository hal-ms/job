create:
	cp config.yml.template config.yml
run:
	env GO111MODULE=on go run main.go
