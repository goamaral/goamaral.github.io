all: setup compile_c

run_c: setup compile_c
	./tmp/main

setup:
	mkdir -p tmp

compile_c:
	gcc -Wall -o tmp/main c/main.c -lm `pkg-config --cflags --libs liburing`

run_go:
	go run go/main.go