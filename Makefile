amf:
	go build -o bin/amf nfs/amf/cmd/main.go	
udm:
	go build -o bin/udm nfs/udm/cmd/main.go	
ausf:
	go build -o bin/ausf nfs/ausf/cmd/main.go	
smf:
	go build -o bin/smf nfs/smf/cmd/main.go	
pcf:
	go build -o bin/pcf nfs/pcf/cmd/main.go	
udr:
	go build -o bin/udr nfs/udr/cmd/main.go	

test_compile:
	go build -o bin/test test_compiling/main.go	

clean:
	rm bin/*

.DEFAULT_GOAL := all
all: amf udm ausf smf pcf udr
