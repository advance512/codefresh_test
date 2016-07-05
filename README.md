Wat iz?
-------

An example composition of 2 Go microservices (using labstack/echo as the RESTful framework)
that are Dockerized, for the intent of loading to Codefresh.

Requirements:
-------------

Make sure Docker (and optionally Docker Compose) are properly installed.

NOTE: This was built and tested on an Ubuntu Linux machine, so no Docker Machine.
If you're using OSX/Windows, you might need to modify the build scripts.

How to run:
-----------

Just execute ./build.sh and if all works well, execute ./run-dc.sh to create the two
microservice containers


