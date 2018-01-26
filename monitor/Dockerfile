FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
COPY main /app
ENTRYPOINT ./main -conductor_ip=conductor
EXPOSE 62071
