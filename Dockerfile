FROM mkenney/chromium-headless:latest

# Install htmltox
COPY ./app /go/src/app
RUN cd /go/src/app \
    && dep ensure \
    && go build -o /go/bin/app

# Grant the process permission to bind to port 80
#RUN setcap 'cap_net_bind_service=+ep' /go/bin/app

WORKDIR /go/src/app
EXPOSE 80
EXPOSE 9222
CMD /go/bin/app
