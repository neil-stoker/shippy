#!/make

SSH_PRIVATE_KEY=`cat ~/.ssh/id_rsa`

x:
	echo $(SSH_PRIVATE_KEY)
	echo ok

build:

	docker build --build-arg SSH_PRIVATE_KEY="$(SSH_PRIVATE_KEY)" -t shippy-cli-consignment .

run:
	docker run shippy-cli-consignment

