package opslevel_jq_parser_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/opslevel/opslevel-go/v2024"
	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

var k8sResource = `{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "annotations": {
			"deployment.kubernetes.io/revision": "243",
			"kots.io/app-slug": "opslevel",
			"opslevel.com/description": "this is a description",
			"opslevel.com/framework": "rails",
			"opslevel.com/language": "ruby",
			"opslevel.com/lifecycle": "alpha",
			"opslevel.com/owner": "velero",
			"opslevel.com/product": "jklabs",
			"opslevel.com/repo.terraform.clusters.dev.opslevel": "gitlab.com:opslevel/terraform",
			"opslevel.com/system": "monolith",
			"opslevel.com/tier": "tier_1",
			"opslevel.com/tools.logs.my-ci": "https://circleci.com",
			"opslevel.com/tools.logs.my-graphs": "https://datadog.com",
			"opslevel.com/tools.logs.my-logs": "https://splunk.com",
			"opslevel.com/tools.logs.my-logs": "https://splunk.com",
			"opslevel.com/tools.logs.my-schedule": "https://pagerduty.com",
			"opslevel.com/tools.logs.my-schedule": "https://pagerduty.com",
			"prop_bool": true,
			"prop_empty_object": {},
			"prop_empty_string": "",
			"prop_object": {"message": "hello world", "condition": true},
			"prop_string": "hello world",
			"repo": "github.com:hashicorp/vault"
        },
        "creationTimestamp": "2023-07-19T18:04:03Z",
        "generation": 243,
        "labels": {
            "app.kubernetes.io/instance": "web",
            "app.kubernetes.io/part-of": "opslevel",
            "kots.io/app-slug": "opslevel",
            "kots.io/backup": "velero"
        },
        "name": "web",
        "namespace": "self-hosted",
        "resourceVersion": "383293724",
        "uid": "19d729de-f708-437c-8b65-10fa06d5dfd5"
    },
    "spec": {
        "progressDeadlineSeconds": 600,
        "replicas": 2,
        "revisionHistoryLimit": 3,
        "selector": {
            "matchLabels": {
                "app.kubernetes.io/instance": "web",
                "app.kubernetes.io/part-of": "opslevel"
            }
        },
        "strategy": {
            "rollingUpdate": {
                "maxSurge": "1",
                "maxUnavailable": 0
            },
            "type": "RollingUpdate"
        },
        "template": {
            "metadata": {
                "annotations": {
                    "kots.io/app-slug": "opslevel"
                },
                "creationTimestamp": null,
                "labels": {
                    "app.kubernetes.io/instance": "web",
                    "app.kubernetes.io/part-of": "opslevel",
					"environment": "dev",
                    "collect-logs": "true"
                }
            },
            "spec": {
                "containers": [
                    {
                        "args": [
                            "bundle",
                            "exec",
                            "puma",
                            "-C ./config/puma.rb"
                        ],
                        "env": [],
                        "envFrom": [
                            {
                                "configMapRef": {
                                    "name": "opslevel"
                                }
                            },
                            {
                                "secretRef": {
                                    "name": "opslevel"
                                }
                            }
                        ],
                        "image": "opslevel/opslevel:main-240131e5",
                        "imagePullPolicy": "Always",
                        "lifecycle": {
                            "preStop": {
                                "exec": {
                                    "command": [
                                        "sleep",
                                        "15"
                                    ]
                                }
                            }
                        },
                        "livenessProbe": {
                            "failureThreshold": 3,
                            "initialDelaySeconds": 3,
                            "periodSeconds": 20,
                            "successThreshold": 1,
                            "tcpSocket": {
                                "port": "opslevel"
                            },
                            "timeoutSeconds": 1
                        },
                        "name": "web",
                        "ports": [
                            {
                                "containerPort": 3000,
                                "name": "opslevel",
                                "protocol": "TCP"
                            }
                        ],
                        "readinessProbe": {
                            "failureThreshold": 3,
                            "httpGet": {
                                "path": "/api/ping",
                                "port": "opslevel",
                                "scheme": "HTTP"
                            },
                            "initialDelaySeconds": 5,
                            "periodSeconds": 10,
                            "successThreshold": 2,
                            "timeoutSeconds": 1
                        },
                        "resources": {
                            "limits": {
                                "cpu": "1",
                                "memory": "1536Mi"
                            },
                            "requests": {
                                "cpu": "500m",
                                "memory": "500Mi"
                            }
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "imagePullSecrets": [
                    {
                        "name": "opslevel-registry"
                    }
                ],
                "initContainers": [
                    {
                        "args": [
                            "bundle",
                            "exec",
                            "rake",
                            "db:abort_if_pending_migrations"
                        ],
                        "envFrom": [
                            {
                                "configMapRef": {
                                    "name": "opslevel"
                                }
                            },
                            {
                                "secretRef": {
                                    "name": "opslevel"
                                }
                            }
                        ],
                        "image": "opslevel/opslevel:main-240131e5",
                        "imagePullPolicy": "Always",
                        "name": "migrations",
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "nodeSelector": {
                    "kubernetes.io/os": "linux"
                },
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "terminationGracePeriodSeconds": 30,
                "topologySpreadConstraints": [
                    {
                        "labelSelector": {
                            "matchLabels": {
                                "app.kubernetes.io/name": "web",
                                "app.kubernetes.io/part-of": "opslevel"
                            }
                        },
                        "maxSkew": 1,
                        "topologyKey": "topology.kubernetes.io/zone",
                        "whenUnsatisfiable": "ScheduleAnyway"
                    }
                ]
            }
        }
    },
    "status": {
        "availableReplicas": 2,
        "conditions": [
            {
                "lastTransitionTime": "2023-07-19T18:05:54Z",
                "lastUpdateTime": "2023-07-19T18:05:54Z",
                "message": "Deployment has minimum availability.",
                "reason": "MinimumReplicasAvailable",
                "status": "True",
                "type": "Available"
            },
            {
                "lastTransitionTime": "2023-08-31T09:00:52Z",
                "lastUpdateTime": "2023-09-25T16:26:30Z",
                "message": "ReplicaSet \"web-6fd48cb855\" has successfully progressed.",
                "reason": "NewReplicaSetAvailable",
                "status": "True",
                "type": "Progressing"
            }
        ],
        "observedGeneration": 243,
        "readyReplicas": 2,
        "replicas": 2,
        "updatedReplicas": 2
    }
}
`

func TestJQServiceParserSimpleConfig(t *testing.T) {
	// Arrange
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(opslevel_jq_parser.SimpleConfig)
	if err != nil {
		t.Error(err)
	}
	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	// Act
	service, err := parser.Run(k8sResource)
	if err != nil {
		t.Error(err)
	}
	// Assert
	autopilot.Equals(t, "web", service.Name)
	autopilot.Equals(t, "self-hosted", service.Owner)
	autopilot.Equals(t, "", service.Lifecycle)
	autopilot.Equals(t, "", service.Tier)
	autopilot.Equals(t, "", service.Product)
	autopilot.Equals(t, "", service.Language)
	autopilot.Equals(t, "", service.Framework)
	// autopilot.Equals(t, "", service.System)
	autopilot.Equals(t, 1, len(service.Aliases))
	autopilot.Equals(t, "k8s:web-self-hosted", service.Aliases[0])
	autopilot.Equals(t, 1, len(service.TagCreates))
	autopilot.Equals(t, opslevel.TagInput{Key: "environment", Value: "dev"}, service.TagCreates[0])
	autopilot.Equals(t, 5, len(service.TagAssigns))
	autopilot.Equals(t, opslevel.TagInput{Key: "imported", Value: "kubectl-opslevel"}, service.TagAssigns[0])
	autopilot.Equals(t, 0, len(service.Tools))
	autopilot.Equals(t, 0, len(service.Repositories))
	// property assignment
	fmt.Println(service.Properties)
	autopilot.Equals(t, 5, len(service.Properties))
	autopilot.Equals(t, "\"true\"", string(service.Properties["prop_bool"]))
	autopilot.Equals(t, "{}", string(service.Properties["prop_empty_object"]))
	autopilot.Equals(t, "\"\"", string(service.Properties["prop_empty_string"]))
	autopilot.Equals(t, `{"condition":true,"message":"hello world"}`, string(service.Properties["prop_object"]))
	autopilot.Equals(t, "\"hello world\"", string(service.Properties["prop_string"]))
}

func TestJQServiceParserSampleConfig(t *testing.T) {
	// Arrange
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(opslevel_jq_parser.SampleConfig)
	if err != nil {
		t.Error(err)
	}

	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	// Act
	service, err := parser.Run(k8sResource)
	if err != nil {
		t.Error(err)
	}
	// Assert
	autopilot.Equals(t, "web", service.Name)
	autopilot.Equals(t, "this is a description", service.Description)
	autopilot.Equals(t, "velero", service.Owner)
	autopilot.Equals(t, "alpha", service.Lifecycle)
	autopilot.Equals(t, "tier_1", service.Tier)
	autopilot.Equals(t, "jklabs", service.Product)
	autopilot.Equals(t, "ruby", service.Language)
	autopilot.Equals(t, "rails", service.Framework)
	// autopilot.Equals(t, "monolith", service.System)
	autopilot.Equals(t, "k8s:web-self-hosted", service.Aliases[0])
	autopilot.Equals(t, "self-hosted-web", service.Aliases[1])
	autopilot.Equals(t, 1, len(service.TagCreates))
	autopilot.Equals(t, opslevel.TagInput{Key: "environment", Value: "dev"}, service.TagCreates[0])
	autopilot.Equals(t, 5, len(service.TagAssigns))
	autopilot.Equals(t, opslevel.TagInput{Key: "imported", Value: "kubectl-opslevel"}, service.TagAssigns[0])
	autopilot.Equals(t, 4, len(service.Tools))
	autopilot.Equals(t, 3, len(service.Repositories))
}

func BenchmarkJQParser_New(b *testing.B) {
	config, _ := opslevel_jq_parser.NewServiceRegistrationConfig(opslevel_jq_parser.SampleConfig)
	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = parser.Run(k8sResource)
	}
}

type Beverage struct {
	Name string
	Oz   int
}

func DeduplicatedBeverages(objects []Beverage) []Beverage {
	return opslevel_jq_parser.Deduplicated(objects, func(b Beverage) string {
		return fmt.Sprintf("%s%d", b.Name, b.Oz)
	})
}

func BeveragesEqual(b1 []Beverage, b2 []Beverage) bool {
	return slices.EqualFunc(b1, b2, func(b1, b2 Beverage) bool {
		return b1.Name == b2.Name && b1.Oz == b2.Oz
	})
}

func TestDeduplicated(t *testing.T) {
	emptyList := []Beverage{}
	emptyDedup := DeduplicatedBeverages(emptyList)
	if !BeveragesEqual(emptyList, emptyDedup) {
		t.Error("an empty list deduplicated should be equal to itself")
	}

	oneElem := []Beverage{
		{Name: "Energy Drink", Oz: 10},
	}
	oneElemDedup := DeduplicatedBeverages(oneElem)
	if !BeveragesEqual(oneElem, oneElemDedup) {
		t.Error("a single element list deduplicated should be equal to itself")
	}

	list := []Beverage{
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Soda", Oz: 12},
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Iced Tea", Oz: 24},
		{Name: "Soda", Oz: 24},
		{Name: "Energy Drink", Oz: 10},
	}
	listDedup := DeduplicatedBeverages(list)
	listDedupExp := []Beverage{
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Iced Tea", Oz: 24},
		{Name: "Soda", Oz: 24},
		{Name: "Energy Drink", Oz: 10},
	}
	if BeveragesEqual(list, listDedup) {
		t.Error("long list deduplicated should NOT be equal to itself")
	}
	if !BeveragesEqual(listDedup, listDedupExp) {
		t.Error("long list deduplicated should be equal to the expected list")
	}
}
