.PHONY: defn

definition/demo.html: piqigames
	./piqigames -in=definition/demo.piqi.pb -out=definition/demo.html

piqigames: docgen.go defn
	go build

defn:
	$(MAKE) $(MAKEARGS) -C definition
