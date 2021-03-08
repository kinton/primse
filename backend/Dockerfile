FROM golang:1.16.0

WORKDIR /backend/app
COPY ./ /backend

RUN go get github.com/gorilla/mux

# RUN go mod init backend — creates go.mod
RUN go build
RUN go mod tidy

ARG ENV
ENTRYPOINT if [ $${ENV} = production ]; \
	then \
		./builds/builded_backend; \
	else \
		# go get github.com/pilu/fresh && \
		# fresh; \
		go get github.com/githubnemo/CompileDaemon && \
		CompileDaemon --build="go build -o ./builds/backend" --command="./builds/backend"; \
	fi

EXPOSE 3000