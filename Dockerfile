FROM golang:1.14.2

WORKDIR /app/

COPY main.go .
COPY plugin.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sonar

FROM adoptopenjdk:11-jre-openj9

WORKDIR /bin/
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.3.0.2102-linux.zip ./sonarscanner.zip

RUN apt-get update \ 
    && curl -sL https://deb.nodesource.com/setup_12.x | bash - \
    && apt-get install -y nodejs unzip \
    && unzip sonarscanner.zip \
    && rm sonarscanner.zip

ENV SONAR_RUNNER_HOME=/bin/sonar-scanner-4.3.0.2102-linux
ENV PATH $PATH:/bin/sonar-scanner-4.3.0.2102-linux/bin

COPY --from=0 /app/sonar .
CMD ["/bin/sonar"]
