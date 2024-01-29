build: clean
	go build -o bin/mr

init:
	./bin/mr init /home/kylehipz/

join:
	./bin/mr join

clean:
	rm -rf bin/*
