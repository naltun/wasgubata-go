NAME=wgb

all:
	go build -o ${NAME}
	sudo mv ${NAME} /usr/local/bin
	touch ${HOME}/.wasgubata.log

clean:
	sudo rm /usr/local/bin/${NAME}
	rm ${HOME}/.wasgubata.log