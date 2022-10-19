package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"sync"
)

const EnvLocal = "local"
const EnvTest = "test"
const EnvStage = "stage"
const EnvProduction = "production"

type App struct {
	Name string `yaml:"name"`
	Ver  string `yaml:"ver"`
	Env  string `yaml:"env"`
}
type Http struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
	Log  string `yaml:"log"`
}
type Pprof struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}
type Websocket struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
	Log  string `yaml:"log"`
}
type Rpc struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
	Ttl  int64  `yaml:"ttl"`
	Log  string `yaml:"log"`
}
type Ps struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Database string `yaml:"database"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	db       int    `yaml:"db"`
}
type Jaeger struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}
type Jwt struct {
	Secret       string `yaml:"secret"`
	DefaultToken string `yaml:"default_token"`
}
type Pager struct {
	PageSize uint `yaml:"page_size"`
}
type Etcd struct {
	Host    []string `yaml:"host"`
	Timeout uint64   `yaml:"timeout"`
}
type Smms struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Email struct {
	Smtp     string `yaml:"smtp"`
	Port     uint64 `yaml:"port"`
	Ssl      bool   `yaml:"ssl"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}
type PubSub struct {
	Brokers []string `yaml:"brokers"`
}
type Manager struct {
	Email map[string]string `yaml:"email"`
}
type Gateway struct {
	Id []string `yaml:"id"`
}

type Config struct {
	App       App
	Http      Http
	Pprof     Pprof
	Websocket Websocket
	Rpc       Rpc
	Ps        Ps
	Mysql     Mysql
	Redis     Redis
	Jaeger    Jaeger
	Jwt       Jwt
	Pager     Pager
	Etcd      Etcd
	Smms      Smms
	Email     Email
	PubSub    PubSub
	Manager   Manager
	Gateway   Gateway
}

var (
	onceConfig     sync.Once
	instanceConfig *Config
)

func GetConfigInstance(root string) *Config {
	onceConfig.Do(func() {
		instanceConfig = &Config{}
		v := viper.New()
		v.AddConfigPath(root)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
		})
		if err := v.ReadInConfig(); err != nil {
		}
		if err:= v.Unmarshal(instanceConfig); err != nil {
		}
	})
	return instanceConfig
}

var Provider=wire.NewSet(GetConfigInstance)
