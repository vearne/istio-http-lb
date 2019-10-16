FROM ubuntu:16.04


RUN apt-get update 
RUN apt install -y curl python3 
RUN apt-get install -y supervisor dnsutils
WORKDIR /data
CMD ["supervisord", "--nodaemon", "--configuration", "/etc/supervisor/supervisord.conf"]
