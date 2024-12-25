# Etapa 1: Compilar a aplicação
FROM maven:3.9.7-eclipse-temurin-21-alpine AS build

WORKDIR /app
COPY pom.xml . 
COPY src ./src
RUN mvn package -DskipTests

# Verificar se o JAR foi gerado corretamente
RUN ls -la /app/target

# Etapa 2: Criar a imagem final para execução
FROM eclipse-temurin:21-jre-alpine

WORKDIR /app
COPY --from=build /app/target/*.jar app.jar
RUN ls -la .
EXPOSE 3000
CMD ["java", "-jar", "app.jar"]
