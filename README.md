docker day 3
------------

* docker workflow
* `Dockerfile` commands
* `docker run` commands
* `docker logs`
* docker port mapping and networking
* multistage builds
* Learn about `FROM scratch`
* `Dockerfile` challenge


docker workflow
---------------

Getting your containerized app to production:

*Simplified - no CI/CD*

1. Locally build and test a Docker image on your development box.

2. Push image to docker registry.

3. Run your image on the target server which will pull your image from a registry. 

*Standard - with CI/CD*

1. Code your application normally.

2. Push your code to a git repo.

3. CI/CD automatically builds and pushes your image to a registry.

4. CI/CD automatically deploys your image to different environments: dev,test,prod,etc.

* [docker workflow - microsoft](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/docker-application-development-process/docker-app-development-workflow)
* [docker workflow - oreilly](https://www.oreilly.com/content/docker-in-production/)

docker build and Dockerfile
---------------------------

Docker can build images automatically by reading the instructions from a Dockerfile. A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. Using docker build users can create an automated build that executes several command-line instructions in succession.

* [Dockerfile build reference](https://docs.docker.com/engine/reference/builder/)

docker run
----------

Docker runs processes in isolated containers. A container is a process which runs on a host. The host may be local or remote. When an operator executes docker run, the container process that runs is isolated in that it has its own file system, its own networking, and its own isolated process tree separate from the host.

* [docker run](https://docs.docker.com/engine/reference/run/)

docker run - ports and networking
---------------------------------

By default, when you create a container, it does not publish any of its ports to the outside world. To make a port available to services outside of Docker, or
to Docker containers which are not connected to the container’s network, use the --publish or -p flag. This creates a firewall rule which maps a container
port to a port on the Docker host. 

* [docker container ports](https://docs.docker.com/config/containers/container-networking/)

docker logs
-----------

Fetch the logs of a container.

* [docker logs](https://docs.docker.com/engine/reference/commandline/logs/)


multistage builds
-----------------

One of the most challenging things about building images is keeping the image size down. Each instruction in the Dockerfile adds a layer to the image, and you need to remember to clean up any artifacts you don’t need before moving on to the next layer. To write a really efficient Dockerfile, you have traditionally needed to employ shell tricks and other logic to keep the layers as small as possible and to ensure that each layer has the artifacts it needs from the previous layer and nothing else.

It was actually very common to have one Dockerfile to use for development (which contained everything needed to build your application), and a slimmed-down one to use for production, which only contained your application and exactly what was needed to run it. This has been referred to as the “builder pattern”. Maintaining two Dockerfiles is not ideal.

* [multistage builds](https://docs.docker.com/develop/develop-images/multistage-build/)

FROM scratch
------------

*What is `FROM scratch`?*

Scratch is a reserved image name that is a stand in for a blank image.
All images must at one point inherit from scratch. See the latest [debian example Dockerfile](https://github.com/debuerreotype/docker-debian-artifacts/blob/3503997cf522377bc4e4967c7f0fcbcb18c69fc8/bullseye/Dockerfile).

*When should from scratch be used?*

* When you are building a base docker image
* When you have a statically compiled binary that doesn’t need any system dependencies
* You want the smallest possible docker image
* You want to remove all dependencies that could possible have security issues. This is the best way to always pass container security scans.

*When should you not use it?*

* If you app has any shared dependencies like libc. (If you didn’t explicitly compile it not to it probably does)
* Your app uses a runtime. Python, Java, etc
* You want to be able to exec into your image and execute debug and inspection commands while your process is running. This one really hurts when your trying to debug things in production.

* [docker base images](https://docs.docker.com/develop/develop-images/baseimages/)

dockerfile challenge
--------------------

Checkout this repo to your local system and change into the directory.

Read and understand what the golang code is doing for later reference See [hello-http.go](hello-http.go)

* Build the image locally
* Figure how to run it in detached mode
* Watch the logs of your container
* Make improvements to the dockerfile so that it is easier to run
* Access the http endpoint in the browser http://localhost:8080/ (make sure you map the port in your run
command)
* Make the hello endpoint say goodbye neighbor without changing the code
* Explain the difference between CMD and ENTRYPOINT and how they work together

docker 2 additional challenge
-----------------------------

Build the [docker alpine base image](https://github.com/alpinelinux/docker-alpine) locally on your machine.

Learning Objectives
* How are base images really created?
* Where do all the directories come from in a base image?
* What are the most important components of a base image?
* Could building my own base images improve security?
* Read and understand all the scripts that make the build possible. On Windows you will no be able to use the `prepare-branch.sh` script. Even if you are on Mac try to build the image without using the `prepare-branch.sh` script.


