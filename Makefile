amf:
	go build -o bin/amf nfs/amf/cmd/main.go	
	go build -o bin/udm nfs/udm/cmd/main.go	
	go build -o bin/ausf nfs/ausf/cmd/main.go	
clean:
	rm bin/*
all: amf
