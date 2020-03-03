package alerts

import (
	"time"
)

// Alert describes a notification alert
type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations,omitempty"`
	StartsAt     time.Time         `json:"startsAt,omitempty"`
	EndsAt       time.Time         `json:"endsAt,omitempty"`
	GeneratorURL string            `json:"generatorURL,omitempty"`
}

// Payload describes the webhook payload as send by the notification service
type Payload struct {
	Receiver          string            `json:"receiver,omitempty"`
	Status            string            `json:"status,omitempty"`
	Alerts            []Alert           `json:"alerts,omitempty"`
	GroupLabels       map[string]string `json:"groupLabels,omitempty"`
	CommonLabels      map[string]string `json:"commonLabels,omitempty"`
	CommonAnnotations map[string]string `json:"commonAnnotations,omitempty"`
	ExternalURL       string            `json:"externalURL,omitempty"`
	Version           string            `json:"version,omitempty"`
	GroupKey          string            `json:"groupKey,omitempty"`
	AlertName         string            `json:"alertName,omitempty"`
}

/*

{
  "receiver": "default",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "rds_test_alert",
        "broker_id": "prometheus-5d829584",
        "dbinstance_identifier": "postgres-4f8fa1ce-57be-4dd2-8e6b-94b4b08cea72",
        "exported_job": "aws_rds",
        "hsdp_instance_guid": "4f8fa1ce-57be-4dd2-8e6b-94b4b08cea72",
        "hsdp_instance_name": "postgres_mds",
        "instance": "localhost:9106",
        "job": "cloudwatch",
        "organization": "suite-phs",
        "region": "na1",
        "severity": "error",
        "space": "production",
        "space_id": "3fd5dd40-a241-41d0-8493-5bb46228d68b"
      },
      "annotations": {
        "description": "CPU is greater than 10.000000% for 30s! Current Value: 10.8196721313002%",
        "summary": "CPU for postgres_mds is too high"
      },
      "startsAt": "2020-02-10T20:03:48.418219825Z",
      "endsAt": "2020-02-10T20:06:48.418219825Z",
      "generatorURL": "http://prometheus-5d829584-0:9090/graph?g0.expr=%28aws_rds_cpuutilization_average+%2A+on%28hsdp_instance_guid%29+group_left%28hsdp_instance_name%29+%28cf_service_instance_info%7Bhsdp_instance_name%3D~%22postgres_mds%7Cpostgres-msps%7Cpostgres_mus%22%7D%29%29+%3E+10&g0.tab=1"
    }
  ],
  "groupLabels": {
    "alertname": "rds_test_alert",
    "hsdp_instance_name": "postgres_mds"
  },
  "commonLabels": {
    "alertname": "rds_test_alert",
    "broker_id": "prometheus-5d829584",
    "dbinstance_identifier": "postgres-4f8fa1ce-57be-4dd2-8e6b-94b4b08cea72",
    "exported_job": "aws_rds",
    "hsdp_instance_guid": "4f8fa1ce-57be-4dd2-8e6b-94b4b08cea72",
    "hsdp_instance_name": "postgres_mds",
    "instance": "localhost:9106",
    "job": "cloudwatch",
    "organization": "suite-phs",
    "region": "na1",
    "severity": "error",
    "space": "production",
    "space_id": "3fd5dd40-a241-41d0-8493-5bb46228d68b"
  },
  "commonAnnotations": {
    "description": "CPU is greater than 10.000000% for 30s! Current Value: 10.8196721313002%",
    "summary": "CPU for postgres_mds is too high"
  },
  "externalURL": "http://prometheus-5d829584-0:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"rds_test_alert\", hsdp_instance_name=\"postgres_mds\"}"
}

*/
