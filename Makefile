SERVICE=sicepat
migrate: 
	for f in ${SERVICE}; do \
		${MAKE} -C $${f} $@; \
	done


.PHONY: build up start down stop

up:
	docker-compose up -d --build
start:
	docker-compose start
down:
	docker-compose down
stop:
	docker-compose stop