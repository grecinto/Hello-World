Welcome to Hello World Microservice!

Pls. find below the build and run instructions

A. Download the source code to your local folder, e.g. - /hello

B. Build and Run in local machine
* Navigate to your download folder, e.g. /hello
$ cd /hello

* Build the binary executable using this cmd:
$ go build

* Run the app (e.g. - HelloWorld):
$ ./HelloWorld

* Using your browser, enter and browse these URLs to test the app:
http://localhost:8080 - this URL should display "Hello!" message on your browser
http://localhost:8080/healthz - this URL should display JSON string such as below, on your browser
{
  "service": "Hello World Service"
  "status": "OK",
  "version": "0.0.1",
  "uptime": "up since 2020-08-04 08:00:05"
}
NOTE: I added "service" name attribute. Usage scenario: so the Application or microservice (if you may) can be identified among potentially many different Web Service/Apps in the enterprise. In real world, a microservice or app is just one of the potentially many Apps running in the Enterprise, so to distinguish one, from the other, a name attribute is required.

C. Build and Run in docker container
* Ensure to install (and running) the "Docker Desktop" to your host machine.

* Build the docker image with the "Hello World" app (in it) using following cmd line instruction:
$ docker build -t helloworld .

NOTE: where helloworld = name of the image


* Run the docker container & the "Hello World" app (in it) using following cmd line instruction:
$ docker run -d -p 3333:8080 --name helloWorldContainer helloworld

NOTE: Where helloWorldContainer = name of the docker container
      helloworld = name of the image as built on previous step
      3333 = available port in your local host where the application listening port will be mapped to
      8080 = the Hello World app's listening port #


* Using your browser, enter and browse these URLs to test the app:
http://localhost:3333 - this URL should display "Hello!" message on your browser
http://localhost:3333/healthz - this URL should display JSON string such as below, on your browser
{
  "service": "Hello World Service"
  "status": "OK",
  "version": "0.0.1",
  "uptime": "up since 2020-08-04 08:00:05"
}
