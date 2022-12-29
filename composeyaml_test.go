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
					Deploy: podmancomposer.Deploy{
						// }{
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
					Deploy: podmancomposer.Deploy{
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
					Deploy: podmancomposer.Deploy{
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
					Deploy: podmancomposer.Deploy{
						Placement: podmancomposer.DeployPlacement{
							Constraints: []interface{}{"disktype=ssd"}},
					},
				},
				},
			},
		},
		"deploy placement constraints map": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      placement:
        constraints:
          disktype: ssd`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Placement: podmancomposer.DeployPlacement{
							Constraints: map[interface{}]interface{}{"disktype": "ssd"}},
					},
				},
				},
			},
		},
		"deploy preferences list": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      placement:
        preferences:
          - datacenter=us-east`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Placement: podmancomposer.DeployPlacement{
							Preferences: []interface{}{"datacenter=us-east"}},
					},
				},
				},
			},
		},
		"deploy preferences map": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      placement:
        preferences:
          datacenter: us-east`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Placement: podmancomposer.DeployPlacement{
							Preferences: map[interface{}]interface{}{"datacenter": "us-east"},
						},
					},
				},
				},
			},
		},
		"deploy replicas": {
			input: `services:
  frontend:
    image: awesome/webapp
    deploy:
      mode: replicated
      replicas: 6`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Mode:     "replicated",
						Replicas: 6,
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Limits: podmancomposer.Resources{
								CPUs:   "0.50",
								Memory: "50M",
								PIDs:   1,
							},
							Reservations: podmancomposer.Resources{
								CPUs:   "0.25",
								Memory: "20M",
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"nvidia-compute"},
								},
								},
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"nvidia-compute"},
									Driver:       "nvidia",
								},
								},
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"nvidia-compute"},
									Count:        2,
								},
								},
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"nvidia-compute"},
									Count:        "all",
								},
								},
							},
						},
					},
				},
				},
			},
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
            device_ids: ["GPU-f123d1c9-26bb-df9b-1c23-4a731f61d8c7", "second"]`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"gpu"},
									DeviceIDs:    []string{"GPU-f123d1c9-26bb-df9b-1c23-4a731f61d8c7", "second"},
								},
								},
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						Resources: podmancomposer.DeployResources{
							Reservations: podmancomposer.Resources{
								Devices: []podmancomposer.ResourceDevice{{
									Capabilities: []interface{}{"gpu"},
									Driver:       "gpuvendor",
									Options:      map[string]string{"virtualization": "false"},
								},
								},
							},
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						RestartPolicy: podmancomposer.DeployRestartPolicy{
							Condition:   "on-failure",
							Delay:       time.Duration(5 * time.Second),
							MaxAttempts: 3,
							Window:      time.Duration(120 * time.Second),
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						RollbackConfig: podmancomposer.DeployConfig{
							Parallelism:     0,
							Delay:           0,
							FailureAction:   "pause",
							Monitor:         time.Duration(1 * time.Minute),
							MaxFailureRatio: 0,
							Order:           "stop-first",
						},
					},
				},
				},
			},
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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Deploy: podmancomposer.Deploy{
						UpdateConfig: podmancomposer.DeployConfig{
							Parallelism:     0,
							Delay:           0,
							FailureAction:   "pause",
							Monitor:         time.Duration(1 * time.Minute),
							MaxFailureRatio: 0,
							Order:           "stop-first",
						},
					},
				},
				},
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			var output podmancomposer.Compose
			err := yaml.Unmarshal([]byte(tc.input), &output)
			if tc.err == nil {
				assert.Nil(t, err, "got an unexpected error %v", err)
				// fmt.Printf("Expected: %#v", tc.expected)
				// fmt.Printf("Output: %#v", output)
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
