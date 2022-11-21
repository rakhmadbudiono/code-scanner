package configs

type Kafka struct {
	Servers       string `envconfig:"KAFKA_SERVERS" default:"localhost"`
	ScanRepoTopic string `envconfig:"KAFKA_SCAN_REPO_TOPIC" default:"code-scanner.repository.scan"`
}
