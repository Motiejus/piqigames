all: demo.piq

.PHONY: all clean

piqi.piqi:
	piqi cc -o $@

piqi.doc.piqi: piqi.doc piqi.piqi
	piqi cc -o $@ $<

doc.piqi.pb: piqi.doc.piqi
	piqi cc -t pb -o $@ $<

demo.piq: doc.piqi.pb demo.piqi 
	piqi compile --self-spec $< demo.piqi -o $@

clean:
	rm -f piqi.piqi piqi.doc.piqi doc.piqi.pb demo.piq
