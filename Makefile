all: demo.piqi.pb piqi.doc.piqi.proto

.PHONY: all clean

piqi.piqi:
	piqi cc -o $@

piqi.doc.piqi: piqi.doc piqi.piqi
	piqi cc -o $@ $<

piqi.doc.piqi.proto: piqi.doc.piqi
	piqi to-proto $<

doc.piqi.pb: piqi.doc.piqi
	piqi cc -t pb -o $@ $<

demo.piqi.pb: doc.piqi.pb demo.piqi 
	piqi compile --self-spec $< demo.piqi -t pb -o $@


clean:
	rm -f piqi.piqi piqi.doc.piqi doc.piqi.pb demo.piqi.pb piqi.doc.piqi.proto
