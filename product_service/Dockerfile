FROM openjdk:17-oracle
LABEL authors="mrirmag"
ARG WAR_FILE=target/*.war
COPY ${WAR_FILE} product_service.war
ENTRYPOINT ["java", "-Dspring.data.mongodb.uri=mongodb://mongodb_java:27017/irmag_innovations", "-Djava.security.egd=file:/dev/./urandom","-jar","/product_service.war"]