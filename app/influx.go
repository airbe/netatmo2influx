package app

import (
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

func getInfluxClient() (client.Client, error) {
	influxClient, influxClientErr := client.NewHTTPClient(client.HTTPConfig{
		Addr:     Config.Influx.Address,
		Username: Config.Influx.Username,
		Password: Config.Influx.Password,
	})

	return influxClient, influxClientErr
}

func setPoints(cl client.Client, err error, values NetatmoValues) ([]*client.Point, error) {
	if err != nil {
		return make([]*client.Point, 0), err
	}

	var points []*client.Point
	for _, value := range values.Values {
		tags := map[string]string{
			"modulename": value.ModuleName,
		}
		fields := make(map[string]interface{})
		fields[value.MetricName] = value.Value
		pt, err := client.NewPoint(
			Config.Influx.MetricPrefix,
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			return make([]*client.Point, 0), err
		}
		points = append(points, pt)
	}

	return points, err
}

func prepareBatchPoints(points []*client.Point, err error) (client.BatchPoints, error) {
	bp, bperr := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Config.Influx.Database,
		Precision: Config.Influx.Precision,
	})

	if err != nil {
		return bp, err
	}

	if bperr != nil {
		return bp, bperr
	}

	for _, pt := range points {
		bp.AddPoint(pt)
	}

	return bp, err
}

func writePoints(bp client.BatchPoints, cl client.Client, err error) error {
	if err != nil {
		return err
	}

	return cl.Write(bp)
}

func SendToInflux(values NetatmoValues) error {
	client, err := getInfluxClient()
	points, err := setPoints(client, err, values)
	batch, err := prepareBatchPoints(points, err)
	return writePoints(batch, client, err)

}
