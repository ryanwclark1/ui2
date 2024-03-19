package config

import (
    "time"
    "flag"
    "log"
    "os"
    "gopkg.in/yaml.v3"
)

// Config represents the top-level configuration structure.
// Defaults are set programmatically when initializing a Config instance.
type Config struct {
    ConfigFile         string            `yaml:"config_file"`
    ExtraConfigFiles   string            `yaml:"extra_config_files"`
    Debug              bool              `yaml:"debug"`
    LogLevel           string            `yaml:"log_level"`
    User               string
    LogFilename        string            `yaml:"log_filename"`
    SessionLifetime    time.Duration     `yaml:"session_lifetime"`
    HTTP               HTTPConfig        `yaml:"http"`
    AMID               ServiceConfig     `yaml:"amid"`
    Auth               ServiceConfig     `yaml:"auth"`
    CallLogd           ServiceConfig     `yaml:"call-logd"`
    Confd              ServiceConfig     `yaml:"confd"`
    Dird               ServiceConfig     `yaml:"dird"`
    Plugind            ServiceConfig     `yaml:"plugind"`
    Provd              ServiceConfig     `yaml:"provd"`
    Webhookd           ServiceConfig     `yaml:"webhookd"`
    Websocketd         WebSocketDConfig  `yaml:"websocketd"`
    EnabledPlugins     map[string]bool   `yaml:"enabled_plugins"`
}

// HTTPConfig represents the HTTP configuration.
type HTTPConfig struct {
    Listen       string `yaml:"listen"`
    Port         int    `yaml:"port"`
    Certificate  *string `yaml:"certificate"`
    PrivateKey   *string `yaml:"private_key"`
}

// ServiceConfig represents the generic service configuration.
type ServiceConfig struct {
    Host    string  `yaml:"host"`
    Port    int     `yaml:"port"`
    Prefix  *string `yaml:"prefix"`
    HTTPS   bool    `yaml:"https"`
}

// WebSocketDConfig represents the WebSocket daemon configuration.
type WebSocketDConfig struct {
    Host               *string `yaml:"host"`
    Port               int     `yaml:"port"`
    Prefix             string  `yaml:"prefix"`
    VerifyCertificate  bool    `yaml:"verify_certificate"`
}


// DefaultConfig creates a new Config instance with default values.
func DefaultConfig() Config {
    return Config{
        ConfigFile:       "/etc/accent-ui/config.yml",
        ExtraConfigFiles: "/etc/accent-ui/conf.d",
        Debug:            false,
        LogLevel:         "info",
        // TODO: Revaulate the need for this field
        User:             "accent",
        LogFilename:      "/var/log/accent-ui.log",
        SessionLifetime:  8 * time.Hour, // Equals to 60 * 60 * 8
        HTTP: HTTPConfig{
            Listen: "127.0.0.1",
            Port:   9296,
        },
        AMID: ServiceConfig{
            Host:  "localhost",
            Port:  9491,
            HTTPS: false,
        },
        Auth: ServiceConfig{
            Host:  "localhost",
            Port:  9497,
            HTTPS: false,
        },
        CallLogd: ServiceConfig{
            Host:  "localhost",
            Port:  9298,
            HTTPS: false,
        },
        Confd: ServiceConfig{
            Host:  "localhost",
            Port:  9486,
            HTTPS: false,
        },
        Dird: ServiceConfig{
            Host:  "localhost",
            Port:  9489,
            HTTPS: false,
        },
        Plugind: ServiceConfig{
            Host:  "localhost",
            Port:  9503,
            HTTPS: false,
        },
        Provd: ServiceConfig{
            Host:  "localhost",
            Port:  8666,
            HTTPS: false,
        },
        Webhookd: ServiceConfig{
            Host:  "localhost",
            Port:  9300,
            HTTPS: false,
        },
        Websocketd: WebSocketDConfig{
            Host:               nil,
            Port:               443,
            Prefix:             "/api/websocketd",
            VerifyCertificate:  false,
        },
        EnabledPlugins: map[string]bool{
            "access_feature":    true,
            "authentication":    true,
            "index":             true,
            "application":       true,
            "agent":             true,
            "cli":               true,
            "call_filter":       true,
            "call_permission":   true,
            "call_pickup":       true,
            "cdr":               true,
            "conference":        true,
            "context":           true,
            "device":            true,
            "dird_profile":      true,
            "dird_source":       true,
            "dhcp":              true,
            "extension":         true,
            "external_auth":     true,
            "funckey":           true,
            "general_settings":  true,
            "group":             true,
            "global_settings":   true,
            "ha":                true,
            "hep":               true,
            "identity":          true,
            "incall":            true,
            "ivr":               true,
            "line":              true,
            "moh":               true,
            "outcall":           true,
            "paging":            true,
            "parking_lot":       true,
            "phonebook":         true,
            "plugin":            true,
            "provisioning":      true,
            "queue":             true,
            "rtp":               true,
            "schedule":          true,
            "sip_template":      true,
            "skill":             true,
            "skillrule":         true,
            "sound":             true,
            "switchboard":       true,
            "transport":         true,
            "trunk":             true,
            "user":              true,
            "voicemail":         true,
            "webhook":           true,
        },
    }
}

func parseCLIArgs() *Config {
    var cfg Config

    flag.StringVar(&cfg.ConfigFile, "config-file", "/etc/accent-ui/config.yml", "The path to the config file.")
    flag.BoolVar(&cfg.Debug, "debug", false, "Log debug messages.")
    flag.StringVar(&cfg.LogLevel, "log-level", "info", "Logs messages with LOG_LEVEL details. Must be one of: critical, error, warning, info, debug.")
    // flag.StringVar(&cfg.User, "user", "", "The owner of the process")

    flag.Parse()

    return &cfg
}

// loadConfiguration reads the YAML configuration from the given file path.
func loadConfiguration(filePath string) (*Config, error) {
    bytes, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var cfg Config
    err = yaml.Unmarshal(bytes, &cfg)
    if err != nil {
        return nil, err
    }

    return &cfg, nil
}

// load takes the initial Config from CLI, reads the file configuration, and applies any CLI overrides.
func load(cliConfig *Config) *Config {
    fileConfig, err := loadConfiguration(cliConfig.ConfigFile)
    if err != nil {
        log.Fatalf("Error loading configuration file: %v", err)
    }

    // Apply CLI overrides
    if cliConfig.Debug {
        fileConfig.Debug = cliConfig.Debug
    }
    if cliConfig.LogLevel != "info" { // Assuming "info" is the default value
        fileConfig.LogLevel = cliConfig.LogLevel
    }
    if cliConfig.User != "" {
        fileConfig.User = cliConfig.User
    }

    return fileConfig
}
