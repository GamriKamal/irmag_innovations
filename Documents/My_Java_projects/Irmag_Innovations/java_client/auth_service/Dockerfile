FROM openjdk:17-oracle
ARG JAR_FILE=target/*.war
COPY ./target/online_store-3.3.1.war app.war
#EXPOSE 8089
ENTRYPOINT ["java", "-jar", "/app.war"]