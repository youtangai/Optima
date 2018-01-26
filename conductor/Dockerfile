FROM centos:latest
# RUN apk --no-cache add ca-certificates 
# ↑https通信を可能にする為
WORKDIR /app
COPY main /app
RUN mkdir -p /var/optima/test && touch /var/optima/test/optima_key.pub
RUN echo "kizunaai" >> /var/optima/test/optima_key.pub
ENTRYPOINT ./main
ENV OPTIMA_DB_HOST=db
EXPOSE 62070
