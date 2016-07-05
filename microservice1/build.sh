#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

echo "Building microservice1."
echo "Getting all Go dependencies..."
go get

echo "Building Go binary..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

echo "Building Docker image microservice1"
docker build -t microservice1 -f Dockerfile .

echo "Generating ./run.sh"
rm -f ./run.sh
echo "#!/bin/bash" >> ./run.sh
echo "docker run --net=codefresh_test -p 1323:1323 --rm --name ms1 -it microservice1" >> ./run.sh
chmod +x ./run.sh

echo "You can now run ./run.sh to start microservice1."



