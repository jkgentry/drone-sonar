FROM golang:1.10.2

WORKDIR /app/

COPY main.go .
COPY plugin.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sonar

FROM openjdk:8-jre-alpine

WORKDIR /bin/
RUN wget https://sonarsource.bintray.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227.zip -O ./sonarscanner.zip  \
    && unzip sonarscanner.zip \
    && rm sonarscanner.zip

ENV SONAR_RUNNER_HOME=/bin/sonar-scanner-3.2.0.1227
ENV PATH $PATH:/bin/sonar-scanner-3.2.0.1227/bin

COPY --from=0 /app/sonar .
CMD ["/bin/sonar"]
