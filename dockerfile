
FROM golang:1.17.8-alpine

RUN mkdir /restapi


WORKDIR /restapi

# ADD ./ /restapi

COPY  , /restapi/

# COPY /home/sudhir/go/src/university/go.sum /app/ 

RUN go mod download

 
RUN go build -o main

EXPOSE 3000

# ENTRYPOINT [ "/restapi/main" ]

CMD ["/restapi/main"]