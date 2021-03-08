FROM golang:1.16.0

ENV GO111MODULE=auto

WORKDIR /backend
COPY ./ /backend

RUN go get github.com/gorilla/mux

# RUN go get github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="go build -o ./app ./app/main.go" --command="./app"

RUN go build -o builded_app ./app

ARG ENV
CMD if [ ${ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	cd app && \
	fresh ./app; \
	fi

EXPOSE 3000