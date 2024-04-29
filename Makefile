run : 
	go run main.go rest

docs : 
	swag init -g internal/delivery/http/rest.go --parseDependency true --parseInternal