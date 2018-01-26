# Makefile

PROGRAM = fftabholder

DESTDIR ?= /usr/local

RM = rm -f
INSTALL_DATA = install
INSTALL_PROGRAM = install

.PHONY: all
all: ${PROGRAM}


${PROGRAM}: main.go icons.go
	go build .

icons.go: mkicons.sh *.png
	sh mkicons.sh *.png > icons.go

.PHONY: install
install: all
	${INSTALL_PROGRAM} -g wheel -m 755 -o root "${PROGRAM}" "${DESTDIR}/sbin/${PROGRAM}"
	${INSTALL_DATA} -g wheel -m 755 -o root "${PROGRAM}.sh" "${DESTDIR}/etc/rc.d/${PROGRAM}"


.PHONY: deinstall
deinstall:
	${RM} "${DESTDIR}/sbin/${PROGRAM}"
	${RM} "${DESTDIR}/etc/rc.d/${PROGRAM}"


.PHONY: clean
clean:
	${RM} ${PROGRAM}
	${RM} *~
