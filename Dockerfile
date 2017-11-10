FROM scratch
ADD etc/ssl/cert.pem /etc/ssl/certs/ca-certificates.crt
ADD config/config.yml /config/
ADD netatmo2influx /
CMD ["/netatmo2influx"]
