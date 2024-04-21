all: setup compile_c

run_c: setup compile_c
	./tmp/main

setup:
	mkdir -p tmp

compile_c:
	gcc -Wall -lm `pkg-config --cflags --libs liburing` -o tmp/main c/main.c