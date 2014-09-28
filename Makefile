.PHONY: defn

piqigames: docgen.go defn
	go build

defn:
	$(MAKE) $(MAKEARGS) -C definition
