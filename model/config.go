package model

type Config struct {
	ConfigBaseRule     ConfigBaseRules  `yaml:"base-rule"`
	PullProxySource    PullProxySources `yaml:"pull-proxy-source"`
	FilterProxyName    []string         `yaml:"filter-proxy-name,omitempty"`
	FilterProxyServer  []string         `yaml:"filter-proxy-server,omitempty"`
	Port               int              `yaml:"port,omitempty"`
	SocksPort          int              `yaml:"socks-port,omitempty"`
	AllowLan           bool             `yaml:"allow-lan"`
	Mode               string           `yaml:"mode,omitempty"`
	LogLevel           string           `yaml:"log-level,omitempty"`
	ExternalController string           `yaml:"external-controller,omitempty"`
	ExternalUi         string           `yaml:"external-ui,omitempty"`
	Secret             string           `yaml:"secret,omitempty"`
	Experimental       Experimental     `yaml:"experimental,omitempty"`
	Dns                Dns              `yaml:"dns,omitempty"`
	FallbackFilter     FallbackFilter   `yaml:"fallback-filter,omitempty"`     // 这个会最终合并到dns 里面去
	CfwBypass          []string         `yaml:"cfw-bypass,omitempty"`          // 仅在windows下有效
	CfwLatencyTimeout  int              `yaml:"cfw-latency-timeout,omitempty"` // 仅在windows下有效
	Proxy              []Proxy          `yaml:"Proxy,omitempty"`               // 这里还可以放一些自定义的代理信息上去哟
	Proxies            []Proxy          `yaml:"proxies,omitempty"`             // 这里还可以放一些自定义的代理信息上去哟 新版配置
	ProxyGroup         []ProxyGroup     `yaml:"Proxy Group,omitempty"`         // 最终会合并到配置文件中，并且这里面的 所有，都会合并到 proxy 组中去
	ProxyGroups        []ProxyGroup     `yaml:"proxy-groups,omitempty"`        // 最终会合并到配置文件中，并且这里面的 所有，都会合并到 proxy 组中去 新版配置
	Rule               []string         `yaml:"Rule,omitempty"`                // 自定义规则列表 最终会合并到配置文件中
	Rules              []string         `yaml:"rules,omitempty"`               // 自定义规则列表 最终会合并到配置文件中 新版配置
	UploadConfig       string           `yaml:"uploadConfig,omitempty"`        // 是否上传，如果不上传就为空，如果上传就写入上传配置前缀
	UpYunConfig        UpYunConfig      `yaml:"upyunConfig,omitempty"`         // 又拍云上传配置
}
