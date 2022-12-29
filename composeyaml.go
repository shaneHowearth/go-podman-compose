package podmancomposer

import "time"

type MapOrListOfString interface{}
type StringOrInt interface{}

type ResourceDevice struct {
	Capabilities MapOrListOfString `yaml:"capabilities,omitempty"`
	Driver       string            `yaml:"driver,omitempty"`
	Count        StringOrInt       `yaml:"count,omitempty"`
	DeviceIDs    []string          `yaml:"device_ids,omitempty"`
	Options      map[string]string `yaml:"options,omitempty"`
}

type Resources struct {
	CPUs    StringOrInt      `yaml:"cpus,omitempty"`
	Memory  StringOrInt      `yaml:"memory,omitempty"`
	PIDs    int32            `yaml:"pids,omitempty"`
	Devices []ResourceDevice `yaml:"devices,omitempty"`
}

type DeployConfig struct {
	Parallelism     int           `yaml:"parallelism,omitempty"`
	Delay           time.Duration `yaml:"delay,omitempty"`
	FailureAction   string        `yaml:"failure_action,omitempty"`
	Monitor         time.Duration `yaml:"monitor,omitempty"`
	MaxFailureRatio int           `yaml:"max_failure_ratio,omitempty"`
	Order           string        `yaml:"order,omitempty"`
}

type DeployRestartPolicy struct {
	Condition   string        `yaml:"condition,omitempty"`
	Delay       time.Duration `yaml:"delay,omitempty"`
	MaxAttempts int           `yaml:"max_attempts,omitempty"`
	Window      time.Duration `yaml:"window,omitempty"`
}

type DeployPlacement struct {
	Constraints MapOrListOfString `yaml:"constraints,omitempty"`
	Preferences MapOrListOfString `yaml:"preferences,omitempty"`
}

type DeployResources struct {
	Limits       Resources `yaml:"limits,omitempty"`
	Reservations Resources `yaml:"reservations,omitempty"`
}

type Deploy struct {
	EndpointMode   string              `yaml:"endpoint_mode,omitempty"`
	Labels         map[string]string   `yaml:"labels,omitempty"`
	Mode           string              `yaml:"mode,omitempty"`
	Placement      DeployPlacement     `yaml:"placement,omitempty"`
	Replicas       int32               `yaml:"replicas,omitempty"`
	Resources      DeployResources     `yaml:"resources,omitempty"`
	RestartPolicy  DeployRestartPolicy `yaml:"restart_policy,omitempty"`
	RollbackConfig DeployConfig        `yaml:"rollback_config,omitempty"`
	UpdateConfig   DeployConfig        `yaml:"update_config,omitempty"`
}

type BuildSecret struct {
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	UID    string `yaml:"uid"`
	GID    string `yaml:"gid"`
	Mode   int    `yaml:"mode"` // This is octal.
}

type Build struct {
	Context    string              `yaml:"context"`
	Dockerfile string              `yaml:"dockerfile,omitempty"`
	Args       MapOrListOfString   `yaml:"args,omitempty"`
	SSH        []string            `yaml:"ssh,omitempty"`
	CacheFrom  []string            `yaml:"cache_from,omitempty"`  // [NAME|type=TYPE[,KEY=VALUE]].
	CacheTo    []string            `yaml:"cache_to,omitempty"`    // [NAME|type=TYPE[,KEY=VALUE]].
	ExtraHosts []string            `yaml:"extra_hosts,omitempty"` // Compose implementations MUST create matching entry with the IP address and hostname in the container's network configuration.
	Isolation  []string            `yaml:"isolation,omitempty"`
	Privileged bool                `yaml:"privileged,omitempty"`
	Labels     MapOrListOfString   `yaml:"labels,omitempty"`
	NoCache    bool                `yaml:"no_cache,omitempty"`
	Pull       bool                `yaml:"pull,omitempty"`
	ShmSize    StringOrInt         `yaml:"shm_size"`
	Target     string              `yaml:"target,omitempty"`
	Secrets    []MapOrListOfString `yaml:"secrets,omitempty"`
	Tags       []string            `yaml:"tags,omitempty"`
	Platforms  []string            `yaml:"platforms,omitempty"`
	Networks   []string            `yaml:"networks,omitempty"`
	Volumes    []string            `yaml:"volumes,omitempty"`
}

type Service struct {
	Image       string            `yaml:"image,omitempty"`
	Ports       MapOrListOfString `yaml:"ports,omitempty"`
	Deploy      Deploy            `yaml:"deploy,omitempty"`
	Profiles    []string          `yaml:"profiles,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Environment []string          `yaml:"environment,omitempty"`
	Build       Build             `yaml:"build,omitempty"`
}

type IPAMConfig struct {
	Subnet       string            `yaml:"subnet"`
	IPRange      string            `yaml:"ip_range"`
	Gateway      string            `yaml:"gateway"`
	AuxAddresses map[string]string `yaml:"aux_addresses"`
	Options      map[string]any    `yaml:"options"`
}

type IPAM struct {
	Driver string     `yaml:"driver"`
	Config IPAMConfig `yaml:"config"`
}

type Network struct {
	Driver     string            `yaml:"driver"`
	External   bool              `yaml:"external"`
	Name       string            `yaml:"name"`
	DriverOpts []map[string]any  `yaml:"driver_opts"`
	Attachable bool              `yaml:"attachable"`
	EnableIPv6 bool              `yaml:"enable_ipv6"`
	IPAM       IPAM              `yaml:"ipam"`
	Internal   bool              `yaml:"internal"`
	Labels     MapOrListOfString `yaml:"labels"`
}

type Volume struct {
	Driver     string            `yaml:"driver"`
	DriverOpts map[string]string `yaml:"driver_opts"`
	External   bool              `yaml:"external"`
	Labels     MapOrListOfString `yaml:"labels"`
	Name       string            `yaml:"name"`
}

type Config struct {
	File     string `yaml:"file"`
	External bool   `yaml:"external"`
	Name     string `yaml:"name"`
}

type Secret struct {
	File        string `yaml:"file,omitempty"`
	Environment string `yaml:"environment,omitempty"`
	External    bool   `yaml:"external,omitempty"`
	Name        string `yaml:"name,omitempty"`
}

type Compose struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]Service `yaml:"services"`
	Networks []Network          `yaml:"networks,omitempty"`
	Volumes  []Volume           `yaml:"volumes,omitempty"`
	Configs  []Config           `yaml:"configs,omitempty"`
	Secrets  map[string]Secret  `yaml:"secrets,omitempty"`
}
