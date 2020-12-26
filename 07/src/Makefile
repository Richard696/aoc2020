INPUTS = sample.txt sample_02.txt #input.txt
# test: App $(INPUTS)
# 	./App $(INPUTS)

# .PHONY: test
main:
	go build -o App App.go
	./App $(INPUTS)

clean:	
	rm App