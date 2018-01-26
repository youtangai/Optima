FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
RUN yum update -y && yum install -y openssh-clients
COPY main /app
COPY join_script.sh /app
ENV CONDUCTOR_IP=conductor
ENTRYPOINT ./main -conductor_ip=conductor
EXPOSE 62071
