# Esta tarea es temporalmente manual pero se debe
# automatizar la obtencion de sus valores. ¿COMO?
DIR0 = github.com/kranfix/ejemcpp
DIR1 = factorial

INCP = ${GOINC}/${DIR0}
LIBP = ${GOPKG}/${DIR0}

CC = g++ --std=c++11
#CC = g++


CPP_FILES = $(wildcard *.cpp)
CC_FILES = $(wildcard *.cc)
OBJS = $(patsubst %.cpp, %.o, $(CPP_FILES))
BINS = $(patsubst %.cc, %, $(CC_FILES))


LIBA = libkfx_factorial.a
HEADERS = factorial.h

LIBAI =  $(patsubst %.a, ${LIBP}/%.a, $(LIBA))
HEADERSI = $(patsubst %.h, ${INCP}/%.h, $(HEADERS))
BINSI = $(patsubst %.cc, ${GOBIN}/%, $(CC_FILES))

INCPATH = -I${INCP}
LIBPATH = -L${LIBP}

# Variables para test
LIBS = -lkfx_factorial

.SUFFIXES: .a .c .cpp .h


default: install-all

build: ${LIBA}
build-bin: ${BINS}
build-all: build build-bin

install: install-inc install-pkg
install-all: install install-bin
install-bin: ${BINSI}
install-pkg: mkdir-pkg build ${LIBAI}
install-inc: mkdir-inc ${HEADERSI}

clean:
	rm -f *.a *.o 

clean-bin:
	rm -f ${BINS}

clean-all: clean clean-bin

uninstall:
	rm -rdf ${LIBP} ${INCP}

uninstall-bin:
	rm -f ${BINSI}

uninstall-all: uninstall uninstall-bin

%: %.cc 
	${CC} -I${CCINC} ${LIBPATH} ${LIBS} -o $@ $<

.cpp.o: ${HEADERS}
	${CC} -c -o $@ $<

${LIBP}/%.a: %.a
	cp -f $< $@

${INCP}/%.h: %.h
	cp -f $< $@

${GOBIN}/%: %
	cp -f $< $@

${LIBA}: ${OBJS}
	ld -r $^ -o $@

mkdir-pkg: ${LIBP}

mkdir-inc: ${INCP}

mkdir-bin: ${GOBIN}

${LIBP}:
	mkdir -p $@

${INCP}:
	mkdir -p $@

${GOBIN}:
	mkdir -p $@

verify:
	ls ${LIBP}
	ls ${INCP}
