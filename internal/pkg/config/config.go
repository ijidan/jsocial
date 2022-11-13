package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/cast"
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

func GetConfigInstance(configPath string) *Config {
	onceConfig.Do(func() {
		instanceConfig = &Config{}
		v := viper.New()
		v.AddConfigPath(cast.ToString(configPath))
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
		})
		if err := v.ReadInConfig(); err != nil {
		}
		if err := v.Unmarshal(instanceConfig); err != nil {
		}
	})
	return instanceConfig
}

func NewApp(conf *Config) *App {
	return &conf.App
}
func NewHttp(conf *Config) *Http {
	return &conf.Http
}

func NewPprof(conf *Config) *Pprof {
	return &conf.Pprof
}
func NewWebsocket(conf *Config) *Websocket {
	return &conf.Websocket
}
func NewRpc(conf *Config) *Rpc {
	return &conf.Rpc
}
func NewPs(conf *Config) *Ps {
	return &conf.Ps
}
func NewMysql(conf *Config) *Mysql {
	return &conf.Mysql
}
func NewRedis(conf *Config) *Redis {
	return &conf.Redis
}
func NewJaeger(conf *Config) *Jaeger {
	return &conf.Jaeger
}
func NewJwt(conf *Config) *Jwt {
	return &conf.Jwt
}
func NewPager(conf *Config) *Pager {
	return &conf.Pager
}
func NewEtcd(conf *Config) *Etcd {
	return &conf.Etcd
}
func NewSmms(conf *Config) *Smms {
	return &conf.Smms
}
func NewEmail(conf *Config) *Email {
	return &conf.Email
}
func NewPubSub(conf *Config) *PubSub {
	return &conf.PubSub
}
func NewManager(conf *Config) *Manager {
	return &conf.Manager
}
func NewGateway(conf *Config) *Gateway {
	return &conf.Gateway
}

var AppProvider = wire.NewSet(NewApp)
var HttpProvider = wire.NewSet(NewHttp)
var PprofProvider = wire.NewSet(NewPprof)
var WebsocketProvider = wire.NewSet(NewWebsocket)
var RpcProvider = wire.NewSet(NewRpc)
var PsProvider = wire.NewSet(NewPs)
var MysqlProvider = wire.NewSet(NewMysql)
var RedisProvider = wire.NewSet(NewRedis)
var JaegerProvider = wire.NewSet(NewJaeger)
var JwtProvider = wire.NewSet(NewJwt)
var PagerProvider = wire.NewSet(NewPager)
var EtcdProvider = wire.NewSet(NewEtcd)
var SmMsProvider = wire.NewSet(NewSmms)
var EmailProvider = wire.NewSet(NewEmail)
var PubSubProvider = wire.NewSet(NewPubSub)
var ManagerProvider = wire.NewSet(NewManager)
var gatewayProvider = wire.NewSet(NewGateway)
