package config

type storageConfig struct {
	Type                    string                      `json:"type"`
	Address                 string                      `json:"address"`
	Password                string                      `json:"password"`
	DB                      int                         `json:"db"`
}

type systemConfig struct {
	Parallelism             int                         `json:"parallelism"`
	WaitPeriod              int                         `json:"waitPeriod"`
	ExecutionMinutes        int                         `json:"executionMinutes"`
	Storage                 storageConfig               `json:"storage"`
}

type systemConfigOverride struct {
	Parallelism             int                         `json:"parallelism,omitempty"`
	WaitPeriod              int                         `json:"waitPeriod,omitempty"`
	ExecutionMinutes        int                         `json:"executionMinutes,omitempty"`
}

type httpAppConfig struct {
	Addr                    string                      `json:"addr"`
	Method                  string                      `json:"method"`
	Headers                 map[string]string           `json:"header"`
	Body                    interface{}                 `json:"body"`
}

type httpConfig struct {
	ConfigOverride          systemConfigOverride        `json:"configOverride,omitempty"`
	Services                map[string]httpAppConfig    `json:"services"`
}

type config struct {
	Config                  systemConfig                `json:"config"`
	HTTP                    httpConfig                  `json:"http"`
}