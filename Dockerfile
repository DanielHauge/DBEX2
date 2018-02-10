FROM golang:jessie


RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get -u gopkg.in/russross/blackfriday.v2
RUN go get github.com/shurcooL/github_flavored_markdown

COPY Handlers.go Handlers.go
COPY Logger.go Logger.go
COPY Main.go Main.go
COPY MongoManual.go ongoManual.go
COPY Router.go Router.go
COPY Routes.go Routes.go

EXPOSE 9191

RUN go build

ENTRYPOINT ./go
