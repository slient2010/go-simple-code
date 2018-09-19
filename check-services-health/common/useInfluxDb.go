package common

import (
	"log"
	"time"

	"check-services-health/res"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	DbUrl     = "http://127.0.0.1:8086"
	DbName    = "servicesStatus"
	Username  = "ops"
	Password  = "devops"
	Precision = "s"
)

func SaveToInfluxDb(logSrcData *res.LogSrv) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     DbUrl,
		Username: Username,
		Password: Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DbName,
		Precision: Precision,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	// tags := map[string]string{"cpu": "cpu-total"}
	tags := map[string]string{"logsrv": "logsrv-total"}
	fields := map[string]interface{}{
		"code":   logSrcData.Code,
		"status": logSrcData.Data.Status,
	}

	// pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	pt, err := client.NewPoint("log_srv_status", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
