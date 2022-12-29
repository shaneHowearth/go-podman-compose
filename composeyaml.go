package podmancomposer

import "time"

type Resources struct {
	CPUs    int32 `yaml:"cpus,omitempty"`
	Memory  int32 `yaml:"memory,omitempty"`
	PIDs    int32 `yaml:"pids,omitempty"`
	Devices struct {
		Capabilities []map[string]string `yaml:"capabilities,omitempty"`
		Driver       string              `yaml:"driver,omitempty"`
		Count        StringOrInt         `yaml:"count,omitempty"`
		DeviceIDs    string              `yaml:"device_ids,omitempty"`
		Options      map[string]string   `yaml:"options,omitempty"`
	} `yaml:"devices,omitempty"`
}

type DeployConfig struct {
	Parallelism     int           `yaml:"parallelism,omitempty"`
	Delay           time.Duration `yaml:"delay,omitempty"`
	FailureAction   string        `yaml:"failure_action,omitempty"`
	Monitor         time.Duration `yaml:"monitor,omitempty"`
	MaxFailureRatio int           `yaml:"max_failure_ratio,omitempty"`
	Order           string        `yaml:"order,omitempty"`
}

type MapOrListOfString interface{}
type StringOrInt interface{}

type Service struct {
	Image  string            `yaml:"image,omitempty"`
	Ports  MapOrListOfString `yaml:"ports,omitempty"`
	Deploy struct {
		EndpointMode string            `yaml:"endpoint_mode,omitempty"`
		Labels       map[string]string `yaml:"labels,omitempty"`
		Mode         string            `yaml:"mode,omitempty"`
		Placement    struct {
			Constraints map[string]string `yaml:"constraints,omitempty"`
			Preferences map[string]string `yaml:"preferences,omitempty"`
		} `yaml:"placement,omitempty"`
		Replicas  int32 `yaml:"replicas,omitempty"`
		Resources struct {
			Limits       Resources `yaml:"limits,omitempty"`
			Reservations Resources `yaml:"reservations,omitempty"`
		} `yaml:"resources,omitempty"`
		RestartPolicy struct {
			Condition   string        `yaml:"condition,omitempty"`
			Delay       time.Duration `yaml:"delay,omitempty"`
			MaxAttempts int           `yaml:"max_attempts,omitempty"`
			Window      time.Duration `yaml:"window,omitempty"`
		} `yaml:"restart_policy,omitempty"`
		RollbackConfig DeployConfig `yaml:"rollback_config,omitempty"`
		UpdateConfig   DeployConfig `yaml:"update_config,omitempty"`
	} `yaml:"deploy,omitempty"`
	Profiles    []string `yaml:"profiles,omitempty"`
	Version     string   `yaml:"version,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Build       struct {
		Context    string            `yaml:"context"`
		Dockerfile string            `yaml:"dockerfile,omitempty"`
		Args       map[string]string `yaml:"args,omitempty"`
		SSH        []string          `yaml:"ssh,omitempty"`
		CacheFrom  []string          `yaml:"cache_from,omitempty"`
		CacheTo    []string          `yaml:"cache_to,omitempty"`
		ExtraHosts []string          `yaml:"extra_hosts,omitempty"`
		Isolation  []string          `yaml:"isolation,omitempty"`
		Privileged bool              `yaml:"privileged,omitempty"`
		Labels     []string          `yaml:"labels,omitempty"`
		NoCache    bool              `yaml:"no_cache,omitempty"`
		Pull       bool              `yaml:"pull,omitempty"`
		ShmSize    StringOrInt       `yaml:"shm_size"`
		Target     string            `yaml:"target,omitempty"`
		Secrets    []struct {
			Source string `yaml:"source"`
			Target string `yaml:"target"`
			UID    string `yaml:"uid"`
			GID    string `yaml:"gid"`
			Mode   int    `yaml:"mode"` // This is octal.
		} `yaml:"secrets,omitempty"`
		Tags      []string `yaml:"tags,omitempty"`
		Platforms []string `yaml:"platforms,omitempty"`
		Networks  []string `yaml:"networks,omitempty"`
		Volumes   []string `yaml:"volumes,omitempty"`
	} `yaml:"build,omitempty"`
}

type Network struct {
	Driver     string           `yaml:"driver"`
	External   bool             `yaml:"external"`
	Name       string           `yaml:"name"`
	DriverOpts []map[string]any `yaml:"driver_opts"`
	Attachable bool             `yaml:"attachable"`
	EnableIPv6 bool             `yaml:"enable_ipv6"`
	IPAM       struct {
		Driver string `yaml:"driver"`
		Config struct {
			Subnet       string            `yaml:"subnet"`
			IPRange      string            `yaml:"ip_range"`
			Gateway      string            `yaml:"gateway"`
			AuxAddresses map[string]string `yaml:"aux_addresses"`
			Options      map[string]any    `yaml:"options"`
		} `yaml:"config"`
	} `yaml:"ipam"`
	Internal bool              `yaml:"internal"`
	Labels   MapOrListOfString `yaml:"labels"`
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
	File        string `yaml:"file"`
	Environment string `yaml:"environment"`
	External    bool   `yaml:"external"`
	Name        string `yaml:"name"`
}

type Compose struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]Service `yaml:"services"`
	Networks []Network          `yaml:"networks,omitempty"`
	Volumes  []Volume           `yaml:"volumes,omitempty"`
	Configs  []Config           `yaml:"configs,omitempty"`
	Secrets  []Secret           `yaml:"secrets,omitempty"`
}
