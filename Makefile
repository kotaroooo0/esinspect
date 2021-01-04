index := test

.PHONY: up
up:
	docker-compose down
	docker-compose up -d
	sh ./elasticsearch/wait-for-it.sh ./elasticsearch/index.sh $(index) ./elasticsearch/settings.json

.PHONY: crawl
crawl:
	touch data.csv
	go run crawler/main.go > data.csv

.PHONY: load
load:
	go run dataloader/main.go -i ${index} -f data.csv


.PHONY: search
search:
	sh ./search/search.sh ${index} ./search/query.sample.json



