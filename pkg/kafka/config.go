package kafka

type Config struct {
	Brokers    []string `mapstructure:"brokers" validate:"required"`
	GroupID    string   `mapstructure:"group_id" validate:"required,gte=0"`
	InitTopics bool     `mapstructure:"init_topics"`
}

type TopicConfig struct {
	TopicName         string `mapstructure:"topic_name" validate:"required,gte=0"`
	Partitions        int    `mapstructure:"partitions" validate:"required"`
	ReplicationFactor int    `mapstructure:"replication_factor" validate:"required"`
}
