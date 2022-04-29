amf:
	go build -o bin/amf nfs/amf/cmd/main.go	
clean:
	rm bin/*
all: amf
