FROM golang:1.16.0

WORKDIR /backend/app
COPY ./ /backend

RUN go get github.com/gorilla/mux

# RUN go get github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="go build -o ./app ./app/main.go" --command="./app"

# RUN go mod init backend — creates go.mod
RUN go build

ARG ENV
CMD /bin/bash -c "if [[ $${ENV} == \"production\" ]]; \
	then \
		app; \
	else \
		go get github.com/pilu/fresh && \
		fresh; \
	fi"

EXPOSE 3000