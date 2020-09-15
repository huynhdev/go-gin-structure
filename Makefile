createdb:
	docker-compose exec -T mysql-development mysql -uroot -phelloworld -e 'CREATE DATABASE blogapi'
dropdb: 
	docker-compose exec -T mysql-development mysql -uroot -phelloworld -e 'DROP DATABASE blogapi'

.PHONY: createdb dropdb
