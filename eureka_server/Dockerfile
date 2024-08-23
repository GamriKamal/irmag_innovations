FROM openjdk:21-jdk
LABEL authors="mrirmag"
ARG JAR_FILE=target/*.jar
COPY ${JAR_FILE} eureka_server.jar
ENTRYPOINT ["java","-jar","/eureka_server.jar"]