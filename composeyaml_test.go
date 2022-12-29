package podmancomposer_test

import (
	"testing"
	"time"

	podmancomposer "github.com/shanehowearth/go-podman-compose"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

/*
func Test_First(t *testing.T) {
	testcases := map[string]struct {
		input    string
		err      error
		expected podmancomposer.Compose
	}{
		"First": {
			input: `services:
  frontend:
      image: awesome/webapp
      ports:
        - "443:8043"
    networks:
      - front-tier
        - back-tier
    configs:
      - httpd-config
      secrets:
        - server-certificate

  backend:
      image: awesome/database
      volumes:
        - db-data:/etc/data
    networks:
      - back-tier

  volumes:
    db-data:
    driver: flocker
    driver_opts:
      size: "10GiB"

  configs:
    httpd-config:
    external: true

secrets:
  server-certificate:
      external: true

  networks:
    # The presence of these objects is sufficient to define them
  front-tier: {}
    back-tier: {}`,
			expected: podmancomposer.Compose{},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				var output podmancomposer.Compose
				err := yaml.Unmarshal([]byte(tc.input), &output)
				if tc.err == nil {
					assert.Nil(t, err, "got an unexpected error %v", err)
				} else {
					assert.NotNil(t, err)
				}
			})

		})
	}
}
*/

func Test_Deploy(t *testing.T) {
	testcases := map[string]struct {
		input    string
		err      error
		expected podmancomposer.Compose
	}{
		"deploy endpoint mode": {
			input: `services:
  frontend:
    image: awesome/webapp
    ports:
      - "8080:80"
    deploy:
      mode: replicated
      replicas: 2
      endpoint_mode: vip`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Ports: []interface{}{"8080:80"},
					Deploy: struct {
						EndpointMode string            `yaml:"endpoint_mode,omitempty"`
						Labels       map[string]string `yaml:"labels,omitempty"`
						Mode         string            `yaml:"mode,omitempty"`
						Placement    struct {
							Constraints podmancomposer.MapOrListOfString `yaml:"constraints,omitempty"`
							Preferences podmancomposer.MapOrListOfString `yaml:"preferences,omitempty"`
						} `yaml:"placement,omitempty"`
						Replicas  int32 `yaml:"replicas,omitempty"`
						Resources struct {
							Limits       podmancomposer.Resources `yaml:"limits,omitempty"`
							Reservations podmancomposer.Resources `yaml:"reservations,omitempty"`
						} `yaml:"resources,omitempty"`
						RestartPolicy struct {
							Condition   string        `yaml:"condition,omitempty"`
							Delay       time.Duration `yaml:"delay,omitempty"`
							MaxAttempts int           `yaml:"max_attempts,omitempty"`
							Window      time.Duration `yaml:"window,omitempty"`
						} `yaml:"restart_policy,omitempty"`
						RollbackConfig podmancomposer.DeployConfig `yaml:"rollback_config,omitempty"`
						UpdateConfig   podmancomposer.DeployConfig `yaml:"update_config,omitempty"`
					}{
						Mode:         "replicated",
						EndpointMode: "vip",
						Replicas:     2,
					},
				},
				},
			},
		},
		"deploy labels": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      labels:
        com.example.description: "This label will appear on the web service"`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: struct {
						EndpointMode string            `yaml:"endpoint_mode,omitempty"`
						Labels       map[string]string `yaml:"labels,omitempty"`
						Mode         string            `yaml:"mode,omitempty"`
						Placement    struct {
							Constraints podmancomposer.MapOrListOfString `yaml:"constraints,omitempty"`
							Preferences podmancomposer.MapOrListOfString `yaml:"preferences,omitempty"`
						} `yaml:"placement,omitempty"`
						Replicas  int32 `yaml:"replicas,omitempty"`
						Resources struct {
							Limits       podmancomposer.Resources `yaml:"limits,omitempty"`
							Reservations podmancomposer.Resources `yaml:"reservations,omitempty"`
						} `yaml:"resources,omitempty"`
						RestartPolicy struct {
							Condition   string        `yaml:"condition,omitempty"`
							Delay       time.Duration `yaml:"delay,omitempty"`
							MaxAttempts int           `yaml:"max_attempts,omitempty"`
							Window      time.Duration `yaml:"window,omitempty"`
						} `yaml:"restart_policy,omitempty"`
						RollbackConfig podmancomposer.DeployConfig `yaml:"rollback_config,omitempty"`
						UpdateConfig   podmancomposer.DeployConfig `yaml:"update_config,omitempty"`
					}{
						Labels: map[string]string{"com.example.description": "This label will appear on the web service"},
					},
				},
				},
			},
		},
		"deploy mode": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      mode: global`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: struct {
						EndpointMode string            `yaml:"endpoint_mode,omitempty"`
						Labels       map[string]string `yaml:"labels,omitempty"`
						Mode         string            `yaml:"mode,omitempty"`
						Placement    struct {
							Constraints podmancomposer.MapOrListOfString `yaml:"constraints,omitempty"`
							Preferences podmancomposer.MapOrListOfString `yaml:"preferences,omitempty"`
						} `yaml:"placement,omitempty"`
						Replicas  int32 `yaml:"replicas,omitempty"`
						Resources struct {
							Limits       podmancomposer.Resources `yaml:"limits,omitempty"`
							Reservations podmancomposer.Resources `yaml:"reservations,omitempty"`
						} `yaml:"resources,omitempty"`
						RestartPolicy struct {
							Condition   string        `yaml:"condition,omitempty"`
							Delay       time.Duration `yaml:"delay,omitempty"`
							MaxAttempts int           `yaml:"max_attempts,omitempty"`
							Window      time.Duration `yaml:"window,omitempty"`
						} `yaml:"restart_policy,omitempty"`
						RollbackConfig podmancomposer.DeployConfig `yaml:"rollback_config,omitempty"`
						UpdateConfig   podmancomposer.DeployConfig `yaml:"update_config,omitempty"`
					}{
						Mode: "global",
					},
				},
				},
			},
		},
		"deploy placement constraints list": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      placement:
        constraints:
          - disktype=ssd`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: struct {
						EndpointMode string            `yaml:"endpoint_mode,omitempty"`
						Labels       map[string]string `yaml:"labels,omitempty"`
						Mode         string            `yaml:"mode,omitempty"`
						Placement    struct {
							Constraints podmancomposer.MapOrListOfString `yaml:"constraints,omitempty"`
							Preferences podmancomposer.MapOrListOfString `yaml:"preferences,omitempty"`
						} `yaml:"placement,omitempty"`
						Replicas  int32 `yaml:"replicas,omitempty"`
						Resources struct {
							Limits       podmancomposer.Resources `yaml:"limits,omitempty"`
							Reservations podmancomposer.Resources `yaml:"reservations,omitempty"`
						} `yaml:"resources,omitempty"`
						RestartPolicy struct {
							Condition   string        `yaml:"condition,omitempty"`
							Delay       time.Duration `yaml:"delay,omitempty"`
							MaxAttempts int           `yaml:"max_attempts,omitempty"`
							Window      time.Duration `yaml:"window,omitempty"`
						} `yaml:"restart_policy,omitempty"`
						RollbackConfig podmancomposer.DeployConfig `yaml:"rollback_config,omitempty"`
						UpdateConfig   podmancomposer.DeployConfig `yaml:"update_config,omitempty"`
					}{
						Placement: struct {
							Constraints podmancomposer.MapOrListOfString `yaml:"constraints,omitempty"`
							Preferences podmancomposer.MapOrListOfString `yaml:"preferences,omitempty"`
						}{Constraints: []interface{}{"disktype=ssd"}},
					},
				},
				},
			},
		},
		/*
					"deploy placement constraints map": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			    placement:
			      constraints:
			        disktype: ssd`,
						expected: podmancomposer.Compose{},
					},
					"deploy preferences list": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			    placement:
			      preferences:
			        - datacenter=us-east`,
						expected: podmancomposer.Compose{},
					},
					"deploy preferences map": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			    placement:
			      preferences:
			        datacenter: us-east`,
						expected: podmancomposer.Compose{},
					},
					"deploy replicas": {
						input: `services:
			  fronted:
			    image: awesome/webapp
			    deploy:
			      mode: replicated
			      replicas: 6`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			      limits:
			        cpus: '0.50'
			        memory: 50M
			        pids: 1
			      reservations:
			        cpus: '0.25'
			        memory: 20M`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations capabilities": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			            - capabilities: ["nvidia-compute"]`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations devices driver": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			          - capabilities: ["nvidia-compute"]
			            driver: nvidia`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations devices count int ": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			          - capabilities: ["nvidia-compute"]
			            count: 2`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations devices count string ": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			          - capabilities: ["nvidia-compute"]
			            count: all`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations devices device_ids": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			          - capabilities: ["gpu"]
			            device_ids: ["GPU-f123d1c9-26bb-df9b-1c23-4a731f61d8c7"]`,
						expected: podmancomposer.Compose{},
					},
					"deploy resources reservations devices options": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      resources:
			        reservations:
			          devices:
			          - capabilities: ["gpu"]
			            driver: gpuvendor
			            options:
			              virtualization: false`,
						expected: podmancomposer.Compose{},
					},
					"deploy restart policy": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      restart_policy:
			        condition: on-failure
			        delay: 5s
			        max_attempts: 3
			        window: 120s`,
						expected: podmancomposer.Compose{},
					},
					"deploy rollback config": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      rollback_config:
			        parallelism: 0
			        delay: 0
			        failure_action: pause
			        monitor: 1m
			        max_failure_rate: 0
			        order: stop-first`,
						expected: podmancomposer.Compose{},
					},
					"update rollback config": {
						input: `services:
			  frontend:
			    image: awesome/webapp
			    deploy:
			      update_config:
			        parallelism: 0
			        delay: 0
			        failure_action: pause
			        monitor: 1m
			        max_failure_rate: 0
			        order: stop-first`,
						expected: podmancomposer.Compose{},
					},
		*/
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			var output podmancomposer.Compose
			err := yaml.Unmarshal([]byte(tc.input), &output)
			if tc.err == nil {
				assert.Nil(t, err, "got an unexpected error %v", err)
				assert.Equal(t, tc.expected, output)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

/*
func Test_Build(t *testing.T) {
	testcases := map[string]struct {
		input    string
		err      error
		expected podmancomposer.Compose
	}{
		"Build dockerfile": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        dockerfile: webapp.Dockerfile`,
			expected: podmancomposer.Compose{},
		},
		"Build args map": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        args:
          GIT_COMMIT: cdc3b19`,
			expected: podmancomposer.Compose{},
		},
		"Build args list": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        args:
          - GIT_COMMIT=cdc3b19`,
			expected: podmancomposer.Compose{},
		},
		"Build args list no val": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        args:
          - GIT_COMMIT`,
			expected: podmancomposer.Compose{},
		},
		"Build ssh list": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        ssh:
          - default`,
			expected: podmancomposer.Compose{},
		},
		"Build ssh": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        ssh: ["default"]`,
			expected: podmancomposer.Compose{},
		},
		"Build ssh custom id": {
			input: `services:
  frontend:
      image: awesome/webapp
      build:
        context: .
        ssh:
          - myproject=~/.ssh/myproject.pem`,
			expected: podmancomposer.Compose{},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				var output podmancomposer.Compose
				err := yaml.Unmarshal([]byte(tc.input), &output)
				if tc.err == nil {
					assert.Nil(t, err, "got an unexpected error %v", err)
				} else {
					assert.NotNil(t, err)
				}
			})

		})
	}
}
*/
