NASM 	:= nasm -f macho test.asm
LD		:= ld -macosx_version_min 10.7.0 -o test test.o
RMO		:= rm -f test.o
RMT		:= rm -f test
RUN		:= ./test

build:
	@${NASM}
	@${LD}
	@${RMO}

run:
	@${RUN}

clean:
	@${RMO}
	@${RMT}