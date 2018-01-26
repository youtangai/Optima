FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
RUN yum install -y openssh-clients
COPY monitor /app
COPY join_script.sh /app
ENV CONDUCTOR_IP=conductor
ENTRYPOINT ./monitor -conductor_ip=conductor
EXPOSE 62071
