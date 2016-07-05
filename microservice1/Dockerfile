# Standard Dockerfile for Codefresh test microservices

# Instructions are for Docker daemon, Docker Compose configuration available.
# Make sure you have a Docker network called codefresh_test up:
# docker network create --driver bridge codefresh_test

FROM golang:alpine

# First run: 
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Then:
# docker build -t microservice1 -f Dockerfile.standard .

RUN mkdir /microservice 
ADD main /microservice/

EXPOSE 1323

# Run the executable by default when the container starts.
CMD ["/microservice/main"]

# Then run:
# docker run --net=codefresh_test -p 1323:1323 --rm --name ms1 -it microservice1
