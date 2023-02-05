package configuration

type MonitoringConfiguration struct {
	MonitorTags map[string]string `json:"monitorTags,omitempty"`
}
