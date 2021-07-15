package settingtypes

type TomlApp struct {
	Runmode    string
	FilscanApi string
}

type TomlServer struct {
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

type TomlDatabase struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

type TomlEnv struct {
	Server   TomlServer
	Database TomlDatabase
}

type TomlConfig struct {
	App         TomlApp
	Development TomlEnv
	Product     TomlEnv
	Test        TomlEnv
}

type AppConfig struct {
	App TomlApp
	TomlEnv
}
