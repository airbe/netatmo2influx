# netatmo2influx

Send data from netatmo api to influxdb

## setup influxdb

```
CREATE DATABASE netatmo
CREATE USER netatmo WITH PASSWORD 'password'
GRANT WRITE ON netatmo TO netatmo
CREATE USER grafana WITH PASSWORD 'password'
GRANT READ ON netatmo TO grafana
```

## Todo

- add some log :)
