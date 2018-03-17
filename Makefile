
build: 
	docker build ./ -t bandc

run: 
	docker run -it \
		-v '/home/mao/my/projects/pet/bandc/:/app/' \
		-p 4000:4000 \
		bandc


