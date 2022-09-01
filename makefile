.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@migrate -path ./database/migrations -database "mysql://root:root@tcp(localhost:3306)/majoo" up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@migrate -path ./database/migrations -database "mysql://root:root@tcp(localhost:3306)/majoo" down 1

.PHONY: deploy
deploy: ## deploy project
	@echo "Deploy Project"
	@docker-compose up --build -d
	@echo "Deploy Project Done"

.PHONY: majoo-seed
majoo-seed: ## cake seed
	@echo "Migrate Users Seed"
	@docker cp ./database/seeds/users_seeder.sql majoo-mysql:/
	@docker exec majoo-mysql /bin/sh -c 'mysql -u root -proot majoo< /users_seeder.sql'
	@echo "Migrate Users Seed Done"
	@echo "Migrate Merchants Seed"
	@docker cp ./database/seeds/merchants_seeder.sql majoo-mysql:/
	@docker exec majoo-mysql /bin/sh -c 'mysql -u root -proot majoo< /merchants_seeder.sql'
	@echo "Migrate Merchants Seed Done"
	@echo "Migrate Outlets Seed"
	@docker cp ./database/seeds/outlets_seeder.sql majoo-mysql:/
	@docker exec majoo-mysql /bin/sh -c 'mysql -u root -proot majoo< /outlets_seeder.sql'
	@echo "Migrate Outlets Seed Done"
	@echo "Migrate Transactions Seed"
	@docker cp ./database/seeds/transactions_seeder.sql majoo-mysql:/
	@docker exec majoo-mysql /bin/sh -c 'mysql -u root -proot majoo< /transactions_seeder.sql'
	@echo "Migrate Transactions Seed Done"
	