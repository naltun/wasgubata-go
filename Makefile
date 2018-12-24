NAME=wgb

all:
	go build -o ${NAME}
	sudo mv ${NAME} /usr/local/bin

clean:
	rm ./wgb