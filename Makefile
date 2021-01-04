index := test

.PHONY: install
install:
	go install

.PHONY: down
down:
	docker-compose down

.PHONY: up
up:
	docker-compose up -d
	sh ./elasticsearch/wait-for-it.sh ./elasticsearch/index.sh $(index) ./elasticsearch/settings.json

.PHONY: crawl
crawl:
	touch data.csv
	esinspect crawl > data.csv

.PHONY: load
load:
	esinspect load -i ${index} -p data.csv


.PHONY: search
search:
	@sh ./search/search.sh ${index} ./search/query.sample.json

.PHONY: analyze
analyze:
	@sh ./analyze/analyze.sh ${index} ${term}
