# docker-debug-server
A damn small http server that returns and logs every request

The server listens to port 8080.

Start a docker container with 
```
docker build -t debug-server .
docker run -p 8080:8080 debug-server
```
