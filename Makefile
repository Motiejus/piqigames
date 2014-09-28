.PHONY: defn all

all: ./piqigames
	./piqigames < definition/demo.piqi.pb

piqigames: docgen.go defn
	go build

defn:
	$(MAKE) $(MAKEARGS) -C definition
