all:
	GOPATH=${PWD} go build github.com/jkakar/otk
	GOPATH=${PWD} go get github.com/jkakar/otk

clean:
	-rm pkg tmp -rf

gofmt:
	GOPATH=${PWD} go fmt github.com/jkakar/otk

check:
	GOPATH=${PWD} go test github.com/jkakar/otk -test.v=true

install-dependencies: install-system-dependencies install-application-dependencies

install-system-dependencies:
	sudo add-apt-repository ppa:gophers/go
	sudo apt-get update
	sudo apt-get install golang-weekly

# FIXME We shouldn't have to use 'sudo' here.  Even worse, these
# commands install these dependencies system-wide, but for now we'll
# live with it.
install-application-dependencies:
	GOPATH=${PWD} sudo go get -u launchpad.net/gocheck
