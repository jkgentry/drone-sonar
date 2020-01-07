FROM golang:1.13.5

WORKDIR /app/

COPY main.go .
COPY plugin.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sonar

FROM openjdk:11-jdk

WORKDIR /bin/
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.2.0.1873.zip ./sonarscanner.zip

RUN apt-get update \ 
    && apt-get install -y nodejs unzip \
    && unzip sonarscanner.zip \
    && rm sonarscanner.zip

ENV SONAR_RUNNER_HOME=/bin/sonar-scanner-4.2.0.1873
ENV PATH $PATH:/bin/sonar-scanner-4.2.0.1873/bin

COPY --from=0 /app/sonar .
CMD ["/bin/sonar"]
