package api

type Config struct {
	Port            int
	LoggingEnabled  bool
	TimingEnabled   bool
	TimingThreshold int64
}
