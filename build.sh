#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

echo ""
echo "Executing ./microservice1/build.sh"
echo ""
cd microservice1
./build.sh | sed 's/^/\t/'
cd ..

echo ""
echo "Executing ./microservice2/build.sh"
echo ""
cd microservice2
./build.sh | sed 's/^/\t/'
cd ..

echo "Generating ./run-dc.sh (run script for Docker Compose)"
rm -f ./run-dc.sh
echo "#!/bin/bash" >> ./run-dc.sh
echo "docker-compose up" >> ./run-dc.sh
chmod +x ./run-dc.sh

echo "Generating ./run-dd.sh (run script for Docker daemon)"
rm -f ./run-dd.sh
echo "#!/bin/bash" >> ./run-dd.sh
echo "docker network create --driver bridge codefresh_test" >> ./run-dd.sh
echo "echo 'Now run ./microservice1/run.sh and ./microservice2/run.sh in two different consoles.'" >> ./run-dd.sh
echo "echo 'The gateway is accessible on http://0.0.0.0:1323/'" >> ./run-dd.sh
chmod +x ./run-dd.sh

echo ""
echo "Now run ./run-dc.sh or ./run-dd.sh"

