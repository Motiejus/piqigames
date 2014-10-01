.PHONY: defn all

all: index.html

index.html: ./piqigames
	./piqigames -in=definition/demo.piqi.pb -out=index.html

piqigames: docgen.go defn
	go build

defn:
	$(MAKE) $(MAKEARGS) -C definition
