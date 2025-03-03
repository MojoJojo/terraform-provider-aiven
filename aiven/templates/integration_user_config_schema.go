// Code generated by go generate; DO NOT EDIT.
// This file was generated at Thu Sep 30 12:46:55 PM CEST 2021

package templates

import "encoding/json"

var (
  // inline JSON was taken from a file:
  // integrations_user_config_schema.json
  integrationJSON = `{
  "alertmanager": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "dashboard": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "datadog": {
    "additionalProperties": false,
    "properties": {
      "datadog_tags": {
        "example": [
          {
            "tag": "foo"
          },
          {
            "comment": "Useful tag",
            "tag": "bar:buzz"
          }
        ],
        "items": {
          "properties": {
            "comment": {
              "example": "Used to tag primary replica metrics",
              "maxLength": 1024,
              "title": "Optional tag explanation",
              "type": "string"
            },
            "tag": {
              "description": "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.",
              "example": "replica:primary",
              "maxLength": 200,
              "minLength": 1,
              "pattern": "^(?!aiven-)[^\\W\\d_](?:[:\\w./-]*[\\w./-])?$",
              "title": "Tag value",
              "type": "string",
              "user_error": "Tags must start with a letter and after that may contain the characters listed below:\nalphanumerics, underscores, minuses, colons, periods, slashes.\nA tag cannot end with a colon.\nTags can be up to 200 characters long and support Unicode.\nTags with prefix 'aiven-' are reserved for Aiven.\nMore info: https://docs.datadoghq.com/getting_started/tagging.\n"
            }
          },
          "required": [
            "tag"
          ],
          "title": "Datadog tag defined by user",
          "type": "object"
        },
        "maxItems": 32,
        "title": "Custom tags provided by user",
        "type": "array"
      },
      "exclude_consumer_groups": {
        "items": {
          "example": "[ group_a, group_b ]",
          "maxLength": 1024,
          "title": "Consumer groups to exclude",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of custom metrics",
        "type": "array"
      },
      "exclude_topics": {
        "items": {
          "example": "[ topic_x, topic_y ]",
          "maxLength": 1024,
          "title": "Topics to exclude",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of topics to exclude",
        "type": "array"
      },
      "include_consumer_groups": {
        "items": {
          "example": "[ group_a, group_b ]",
          "maxLength": 1024,
          "title": "Consumer groups to include",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of custom metrics",
        "type": "array"
      },
      "include_topics": {
        "items": {
          "example": "[ topic_x, topic_y ]",
          "maxLength": 1024,
          "title": "Topics to include",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of topics to include",
        "type": "array"
      },
      "kafka_custom_metrics": {
        "items": {
          "example": "kafka.log.log_size",
          "maxLength": 1024,
          "title": "Metric name",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of custom metrics",
        "type": "array"
      },
      "max_jmx_metrics": {
        "example": 2000,
        "maximum": 100000,
        "minimum": 10,
        "title": "Maximum number of JMX metrics to send",
        "type": "integer"
      }
    },
    "type": "object"
  },
  "datasource": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "external_aws_cloudwatch_logs": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "external_aws_cloudwatch_metrics": {
    "additionalProperties": false,
    "properties": {
      "dropped_metrics": {
        "example": [
          {
            "field": "evicted_keys",
            "metric": "redis"
          }
        ],
        "items": {
          "additionalProperties": false,
          "properties": {
            "field": {
              "example": "used",
              "maxLength": 1000,
              "title": "Identifier of a value in the metric",
              "type": "string"
            },
            "metric": {
              "example": "java.lang:Memory",
              "maxLength": 1000,
              "title": "Identifier of the metric",
              "type": "string"
            }
          },
          "required": [
            "metric",
            "field"
          ],
          "title": "Metric name and subfield",
          "type": "object"
        },
        "maxItems": 1024,
        "title": "Metrics to not send to AWS CloudWatch (takes precedence over extra_metrics)",
        "type": "array"
      },
      "extra_metrics": {
        "example": [
          {
            "field": "evicted_keys",
            "metric": "redis"
          }
        ],
        "items": {
          "additionalProperties": false,
          "properties": {
            "field": {
              "example": "used",
              "maxLength": 1000,
              "title": "Identifier of a value in the metric",
              "type": "string"
            },
            "metric": {
              "example": "java.lang:Memory",
              "maxLength": 1000,
              "title": "Identifier of the metric",
              "type": "string"
            }
          },
          "required": [
            "metric",
            "field"
          ],
          "title": "Metric name and subfield",
          "type": "object"
        },
        "maxItems": 1024,
        "title": "Metrics to allow through to AWS CloudWatch (in addition to default metrics)",
        "type": "array"
      }
    },
    "title": "External AWS CloudWatch Metrics integration user config",
    "type": "object"
  },
  "external_elasticsearch_logs": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "external_google_cloud_logging": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "flink": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "internal_connectivity": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "jolokia": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "kafka_connect": {
    "additionalProperties": false,
    "properties": {
      "kafka_connect": {
        "additionalProperties": false,
        "properties": {
          "config_storage_topic": {
            "example": "__connect_configs",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration data are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          },
          "group_id": {
            "example": "connect",
            "maxLength": 249,
            "title": "A unique string that identifies the Connect cluster group this worker belongs to.",
            "type": "string"
          },
          "offset_storage_topic": {
            "example": "__connect_offsets",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration offsets are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          },
          "status_storage_topic": {
            "example": "__connect_status",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration status updates are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          }
        },
        "title": "Kafka Connect service configuration values",
        "type": "object"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "kafka_logs": {
    "additionalProperties": false,
    "properties": {
      "kafka_topic": {
        "example": "mytopic",
        "maxLength": 249,
        "minLength": 1,
        "pattern": "^(?!\\.$|\\.\\.$)[-_.A-Za-z0-9]+$",
        "title": "Topic name",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, underscores, dashes or dots, max 249 characters"
      }
    },
    "type": "object"
  },
  "kafka_mirrormaker": {
    "additionalProperties": false,
    "properties": {
      "cluster_alias": {
        "description": "The alias under which the Kafka cluster is known to MirrorMaker. Can contain the following symbols: ASCII alphanumerics, '.', '_', and '-'.",
        "example": "kafka-abc",
        "maxLength": 128,
        "pattern": "^[a-zA-Z0-9_.-]+$",
        "title": "Kafka cluster alias",
        "type": "string",
        "user_error": "Must consist of alpha-numeric ASCII characters, dashes, underscores, and dots."
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "logs": {
    "additionalProperties": false,
    "properties": {
      "elasticsearch_index_days_max": {
        "default": 3,
        "example": 5,
        "maximum": 10000,
        "minimum": 1,
        "title": "Elasticsearch index retention limit",
        "type": "integer"
      },
      "elasticsearch_index_prefix": {
        "default": "logs",
        "example": "logs",
        "maxLength": 1024,
        "minLength": 1,
        "title": "Elasticsearch index prefix",
        "type": "string"
      }
    },
    "type": "object"
  },
  "m3aggregator": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "m3coordinator": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "metrics": {
    "additionalProperties": false,
    "properties": {
      "database": {
        "example": "metrics",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-_A-Za-z0-9]{0,39}$",
        "title": "Name of the database where to store metric datapoints. Only affects PostgreSQL destinations. Defaults to 'metrics'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, underscores or dashes, may not start with dash, max 40 characters"
      },
      "retention_days": {
        "example": 30,
        "maximum": 10000,
        "minimum": 0,
        "title": "Number of days to keep old metrics. Only affects PostgreSQL destinations. Set to 0 for no automatic cleanup. Defaults to 30 days.",
        "type": "integer"
      },
      "ro_username": {
        "example": "metrics_reader",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-._A-Za-z0-9]{0,39}$",
        "title": "Name of a user that can be used to read metrics. This will be used for Grafana integration (if enabled) to prevent Grafana users from making undesired changes. Only affects PostgreSQL destinations. Defaults to 'metrics_reader'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, dots, underscores or dashes, may not start with dash or dot, max 40 characters"
      },
      "source_mysql": {
        "additionalProperties": false,
        "properties": {
          "telegraf": {
            "additionalProperties": false,
            "properties": {
              "gather_event_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS",
                "type": "boolean"
              },
              "gather_file_events_stats": {
                "example": false,
                "title": "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME",
                "type": "boolean"
              },
              "gather_index_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE",
                "type": "boolean"
              },
              "gather_info_schema_auto_inc": {
                "example": false,
                "title": "Gather auto_increment columns and max values from information schema",
                "type": "boolean"
              },
              "gather_innodb_metrics": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS",
                "type": "boolean"
              },
              "gather_perf_events_statements": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST",
                "type": "boolean"
              },
              "gather_process_list": {
                "example": true,
                "title": "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST",
                "type": "boolean"
              },
              "gather_slave_status": {
                "example": true,
                "title": "Gather metrics from SHOW SLAVE STATUS command output",
                "type": "boolean"
              },
              "gather_table_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE",
                "type": "boolean"
              },
              "gather_table_lock_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS",
                "type": "boolean"
              },
              "gather_table_schema": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.TABLES",
                "type": "boolean"
              },
              "perf_events_statements_digest_text_limit": {
                "example": 120,
                "maximum": 2048,
                "minimum": 1,
                "title": "Truncates digest text from perf_events_statements into this many characters",
                "type": "integer"
              },
              "perf_events_statements_limit": {
                "example": 250,
                "maximum": 4000,
                "minimum": 1,
                "title": "Limits metrics from perf_events_statements",
                "type": "integer"
              },
              "perf_events_statements_time_limit": {
                "example": 86400,
                "maximum": 2592000,
                "minimum": 1,
                "title": "Only include perf_events_statements whose last seen is less than this many seconds",
                "type": "integer"
              }
            },
            "title": "Configuration options for Telegraf MySQL input plugin",
            "type": "object"
          }
        },
        "title": "Configuration options for metrics where source service is MySQL",
        "type": "object"
      },
      "username": {
        "example": "metrics_writer",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-._A-Za-z0-9]{0,39}$",
        "title": "Name of the user used to write metrics. Only affects PostgreSQL destinations. Defaults to 'metrics_writer'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, dots, underscores or dashes, may not start with dash or dot, max 40 characters"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "mirrormaker": {
    "additionalProperties": false,
    "properties": {
      "mirrormaker_whitelist": {
        "default": ".*",
        "example": ".*",
        "maxLength": 1000,
        "minLength": 1,
        "title": "Mirrormaker topic whitelist",
        "type": "string"
      }
    },
    "type": "object"
  },
  "prometheus": {
    "additionalProperties": false,
    "properties": {
      "source_mysql": {
        "additionalProperties": false,
        "properties": {
          "telegraf": {
            "additionalProperties": false,
            "properties": {
              "gather_event_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS",
                "type": "boolean"
              },
              "gather_file_events_stats": {
                "example": false,
                "title": "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME",
                "type": "boolean"
              },
              "gather_index_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE",
                "type": "boolean"
              },
              "gather_info_schema_auto_inc": {
                "example": false,
                "title": "Gather auto_increment columns and max values from information schema",
                "type": "boolean"
              },
              "gather_innodb_metrics": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS",
                "type": "boolean"
              },
              "gather_perf_events_statements": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST",
                "type": "boolean"
              },
              "gather_process_list": {
                "example": true,
                "title": "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST",
                "type": "boolean"
              },
              "gather_slave_status": {
                "example": true,
                "title": "Gather metrics from SHOW SLAVE STATUS command output",
                "type": "boolean"
              },
              "gather_table_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE",
                "type": "boolean"
              },
              "gather_table_lock_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS",
                "type": "boolean"
              },
              "gather_table_schema": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.TABLES",
                "type": "boolean"
              },
              "perf_events_statements_digest_text_limit": {
                "example": 120,
                "maximum": 2048,
                "minimum": 1,
                "title": "Truncates digest text from perf_events_statements into this many characters",
                "type": "integer"
              },
              "perf_events_statements_limit": {
                "example": 250,
                "maximum": 4000,
                "minimum": 1,
                "title": "Limits metrics from perf_events_statements",
                "type": "integer"
              },
              "perf_events_statements_time_limit": {
                "example": 86400,
                "maximum": 2592000,
                "minimum": 1,
                "title": "Only include perf_events_statements whose last seen is less than this many seconds",
                "type": "integer"
              }
            },
            "title": "Configuration options for Telegraf MySQL input plugin",
            "type": "object"
          }
        },
        "title": "Configuration options for metrics where source service is MySQL",
        "type": "object"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "read_replica": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "rsyslog": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "schema_registry_proxy": {
    "additionalProperties": false,
    "properties": {},
    "title": "Integration user config",
    "type": "object"
  },
  "signalfx": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  }
}`
)

func init() {
  var integrationSchema map[string]interface{}
  if err := json.Unmarshal([]byte(integrationJSON), &integrationSchema); err != nil {
    panic("cannot unmarshal user configuration options integration JSON', error :" + err.Error())
  }
  userConfigSchemas["integration"] = integrationSchema
}
