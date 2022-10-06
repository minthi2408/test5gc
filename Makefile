amf:
	go build -o bin/amf nfs/amf/cmd/main.go	
udm:
	go build -o bin/udm nfs/udm/cmd/main.go	
ausf:
	go build -o bin/ausf nfs/ausf/cmd/main.go	
smf:
	go build -o bin/smf nfs/smf/cmd/main.go	
test_compile:
	go build -o bin/test test_compiling/main.go	

clean:
	rm bin/*
all: amf udm ausf smf
