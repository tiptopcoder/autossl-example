FROM openresty/openresty:centos

RUN yum update -y \
    && yum install -y gcc \
    && luarocks install lua-resty-auto-ssl \
    && luarocks install lua-resty-http

# Create /etc/resty-auto-ssl and make sure it's writable by whichever user your
# nginx workers run as
RUN mkdir /etc/resty-auto-ssl
RUN chmod 777 /etc/resty-auto-ssl
# RUN chown nginx /etc/resty-auto-ssl
RUN rm /etc/nginx/conf.d/default.conf

COPY ./autossl.conf /etc/nginx/conf.d/autossl.conf
COPY ./ssl /etc/ssl
