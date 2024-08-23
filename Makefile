# Определяем переменные
EUREKA_DIR = /home/mr-irmag/Desktop/Irmag_Innovations/java_client/eureka_server
PRODUCT_DIR = /home/mr-irmag/Desktop/Irmag_Innovations/java_client/product_service
APIGATEWAY_DIR = /home/mr-irmag/Desktop/Irmag_Innovations/java_client/apigateway
GOLANG_PARSER_DIR = /home/mr-irmag/Desktop/Irmag_Innovations/golang_client/parseAndSendData_service

EUREKA_DOCKERFILE = $(EUREKA_DIR)/Dockerfile
PRODUCT_DOCKERFILE = $(PRODUCT_DIR)/Dockerfile
APIGATEWAY_DOCKERFILE = $(APIGATEWAY_DIR)/Dockerfile
GOLANG_PARSER_DOCKERFILE = $(GOLANG_PARSER_DIR)/Dockerfile
.PHONY: launch clean install build_docker build_all up_containers

launch: clean install build_all

clean:
	@echo "Running mvn clean for all projects..."
	mvn -f $(EUREKA_DIR)/pom.xml clean
	mvn -f $(PRODUCT_DIR)/pom.xml clean
	mvn -f $(APIGATEWAY_DIR)/pom.xml clean

install:
	@echo "Running mvn install -DskipTests for all projects..."
	mvn -f $(EUREKA_DIR)/pom.xml install -DskipTests
	mvn -f $(PRODUCT_DIR)/pom.xml install -DskipTests
	mvn -f $(APIGATEWAY_DIR)/pom.xml install -DskipTests

build_docker:
	@echo "Building Docker images..."
	docker build -t eureka_server -f $(EUREKA_DOCKERFILE) $(EUREKA_DIR)
	docker build -t product_service -f $(PRODUCT_DOCKERFILE) $(PRODUCT_DIR)
	docker build -t apigateway -f $(APIGATEWAY_DOCKERFILE) $(APIGATEWAY_DIR)


build_all: build_docker
	docker-compose build

up_containers:
	docker-compose up

build_current:
	mvn -f $(APIGATEWAY_DIR)/pom.xml clean
	mvn -f $(APIGATEWAY_DIR)/pom.xml install -DskipTests
	docker build -t apigateway -f $(APIGATEWAY_DOCKERFILE) $(APIGATEWAY_DIR)
