FROM golang:1.18 as build
WORKDIR /root/
ADD . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/ .

FROM alpine:latest
WORKDIR /root/
COPY --from=build /root/build/pet-reminder /root/pet-reminder
CMD ["/root/pet-reminder"]