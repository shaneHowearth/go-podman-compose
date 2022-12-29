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
				assert.Equal(t, tc.expected, output)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

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
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context:    ".",
						Dockerfile: "webapp.Dockerfile",
					},
				},
				},
			},
		},
		"Build args map": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           args:
             GIT_COMMIT: cdc3b19`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Args:    map[interface{}]interface{}{"GIT_COMMIT": "cdc3b19"},
					},
				},
				},
			},
		},
		"Build args list": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           args:
             - GIT_COMMIT=cdc3b19`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Args:    []interface{}{"GIT_COMMIT=cdc3b19"},
					},
				},
				},
			},
		},
		"Build args list no val": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           args:
             - GIT_COMMIT`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Args:    []interface{}{"GIT_COMMIT"},
					},
				},
				},
			},
		},
		"Build ssh list": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           ssh:
             - default`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						SSH:     []string{"default"},
					},
				},
				},
			},
		},
		"Build ssh": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           ssh: ["default"]`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						SSH:     []string{"default"},
					},
				},
				},
			},
		},
		"Build ssh custom id": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           ssh:
             - myproject=~/.ssh/myproject.pem`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						SSH:     []string{"myproject=~/.ssh/myproject.pem"},
					},
				},
				},
			},
		},
		"Build cache from": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           cache_from:
             - alpine:latest
             - type=local,src=path/to/cache
             - type=gha`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context:   ".",
						CacheFrom: []string{"alpine:latest", "type=local,src=path/to/cache", "type=gha"},
					},
				},
				},
			},
		},
		"Build cache to": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           cache_to:
             - user/app:cache
             - type=local,src=path/to/cache`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						CacheTo: []string{"user/app:cache", "type=local,src=path/to/cache"},
					},
				},
				},
			},
		},
		"Build extra hosts": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           extra_hosts:
             - "somehost:162.242.195.82"
             - "otherhost:50.31.209.229"`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context:    ".",
						ExtraHosts: []string{"somehost:162.242.195.82", "otherhost:50.31.209.229"},
					},
				},
				},
			},
		},
		"Build isolation": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           isolation:
             - stuff`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context:   ".",
						Isolation: []string{"stuff"},
					},
				},
				},
			},
		},
		"Build privileged": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           privileged: true`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context:    ".",
						Privileged: true,
					},
				},
				},
			},
		},
		"Build labels map": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           labels:
             com.example.description: "Accounting webapp"
             com.example.department: "Finance"
             com.example.label-with-empty-value: ""`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Labels:  map[interface{}]interface{}{"com.example.description": "Accounting webapp", "com.example.department": "Finance", "com.example.label-with-empty-value": ""},
					},
				},
				},
			},
		},
		"Build labels list": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           labels:
             - "com.example.description=Accounting webapp"
             - "com.example.department=Finance"
             - "com.example.label-with-empty-value"`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Labels:  []interface{}{"com.example.description=Accounting webapp", "com.example.department=Finance", "com.example.label-with-empty-value"},
					},
				},
				},
			},
		},
		"Build no_cache": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           no_cache: true`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						NoCache: true},
				},
				},
			},
		},
		"Build pull": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           pull: true`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Pull:    true},
				},
				},
			},
		},
		"Build shm size string": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           shm_size: 2gb`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						ShmSize: "2gb"},
				},
				},
			},
		},
		"Build shm size int": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           shm_size: 10000000`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						ShmSize: 10000000},
				},
				},
			},
		},
		"Build target": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           target: prod`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Target:  "prod"},
				},
				},
			},
		},
		"Build secrets short syntax": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           secrets:
             - server-certificate
secrets:
  server-certificate:
      file: ./server.cert`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Secrets: []podmancomposer.MapOrListOfString{"server-certificate"},
					},
				},
				},
				Secrets: map[string]podmancomposer.Secret{
					"server-certificate": {
						File: "./server.cert",
					},
				},
			},
		},
		"Build secrets long syntax": {
			input: `services:
     frontend:
         image: awesome/webapp
         build:
           context: .
           secrets:
             - source: server-certificate
               target: server.cert
               uid: "103"
               gid: "103"
               mode: 0440
secrets:
  server-certificate:
      external: true`,
			expected: podmancomposer.Compose{
				Services: map[string]podmancomposer.Service{"frontend": {
					Image: "awesome/webapp",
					Build: podmancomposer.Build{
						Context: ".",
						Secrets: []podmancomposer.MapOrListOfString{map[interface{}]interface{}{
							"source": "server-certificate",
							"target": "server.cert",
							"uid":    "103",
							"gid":    "103",
							"mode":   0440,
						}},
					},
				},
				},
				Secrets: map[string]podmancomposer.Secret{
					"server-certificate": {
						External: true,
					},
				},
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
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
		})
	}
}
