FROM debian:stable-slim
ADD hello-app /
EXPOSE 8080
ENTRYPOINT ["/hello-app"]
