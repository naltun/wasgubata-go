NAME=wgb
LOGFILE=.wasgubata.log

all:
	go build -o ${NAME}
	sudo mv ${NAME} /usr/local/bin
	touch ${HOME}/${LOGFILE}

clean:
	sudo rm /usr/local/bin/${NAME}
	rm ${HOME}/${LOGFILE}