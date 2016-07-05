Wat iz?
-------

An example composition of 2 Go microservices (using labstack/echo as the RESTful framework)
that are Dockerized, for the intent of loading to Codefresh.

Requirements:
-------------

Make sure Docker and Docker Compose are properly installed.

Since this this builds static Go binaries and includes them in the containers,
go 1.6 must be installed. 

NOTE: This was built and tested on an Ubuntu Linux machine, so no need for Docker Machine.

How to run:
-----------

Just execute ./build.sh and if all works well, execute ./run-dc.sh to create the two
microservice containers

If you're using OSX/Windows, you might need to modify the build scripts.


