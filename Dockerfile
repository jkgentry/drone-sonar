FROM golang:1.10.2

WORKDIR /app/

COPY main.go .
COPY plugin.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sonar

FROM openjdk:8-jre-alpine

WORKDIR /bin/
RUN wget http://repo1.maven.org/maven2/org/codehaus/sonar/runner/sonar-runner-dist/2.4/sonar-runner-dist-2.4.jar -O ./sonar-runner.jar
COPY --from=0 /app/sonar .
CMD ["/bin/sonar"]
