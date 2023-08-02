FROM golang:1.20 as build

WORKDIR /notifier

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags lambda.norpc -o main

# ---

FROM public.ecr.aws/lambda/provided:al2

COPY --from=build /notifier/main ./main
ENTRYPOINT ["./main"]
