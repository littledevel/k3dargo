FROM debian:12-slim
RUN apt-get update && apt-get install -y \
    bind9-utils \
    iproute2 \
    curl \
    netcat-openbsd \
    vim
RUN apt-get install -y nginx
COPY index.html /var/www/html/ 
COPY default /etc/nginx/sites-available/
COPY nginx.conf /etc/nginx/
CMD ["nginx"]
