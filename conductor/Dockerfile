FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
COPY main /app
ENTRYPOINT ./main
ENV OPTIMA_DB_HOST=db
EXPOSE 62070
