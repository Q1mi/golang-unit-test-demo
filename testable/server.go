package main

// Config 配置项结构体
type Config struct {
	// ...
}

// LoadConfFromFile 从配置文件中加载配置
func LoadConfFromFile(filename string) *Config {
	return &Config{}
}

// Server server 程序
type Server struct {
	Config *Config
}

//// NewServer Server 构造函数
//func NewServer() *Server {
//	return &Server{
//		// 隐式创建依赖项
//		Config: LoadConfFromFile("./config.toml"),
//	}
//}

// NewServer Server 构造函数
func NewServer(conf *Config) *Server {
	return &Server{
		// 隐式创建依赖项
		Config: conf,
	}
}
