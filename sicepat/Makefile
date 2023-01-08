.PHONY: migrate
migrate:
	for f in $(shell ls migrations); do \
		docker exec -i db-sicepat mysql -uroot -ptest_pass < migrations/$${f}; \
	done
