FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
COPY main /app
COPY ../join_script/join_script /app
ENTRYPOINT ./main -conductor_ip=conductor
EXPOSE 62071
