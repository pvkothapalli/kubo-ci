package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	Iaas             string           `json:"iaas"`
	TimeoutScale     float64          `json:"timeout_scale"`
	AWS              AWS              `json:"aws"`
	Bosh             Bosh             `json:"bosh"`
	Turbulence       Turbulence       `json:"turbulence"`
	TurbulenceTests  TurbulenceTests  `json:"turbulence_tests"`
	Cf               Cf               `json:"cf"`
	Kubernetes       Kubernetes       `json:"kubernetes"`
	CFCR             CFCR             `json:"cfcr"`
	IntegrationTests IntegrationTests `json:"integration_tests"`
	Conformance      Conformance      `json:"conformance"`
}

type AWS struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Region          string `json:"region"`
	IngressGroupID  string `json:"ingress_group_id"`
}

type Bosh struct {
	Environment  string `json:"environment"`
	CaCert       string `json:"ca_cert"`
	Client       string `json:"client"`
	ClientSecret string `json:"client_secret"`
	Deployment   string `json:"deployment"`
}

type Turbulence struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	CaCert   string `json:"ca_cert"`
}

type TurbulenceTests struct {
	IncludeWorkerDrain        bool `json:"include_worker_drain"`
	IncludeWorkerFailure      bool `json:"include_worker_failure"`
	IncludeMasterFailure      bool `json:"include_master_failure"`
	IncludePersistenceFailure bool `json:"include_persistence_failure"`
}

type Cf struct {
	AppsDomain string `json:"apps_domain"`
}

type IntegrationTests struct {
	IncludeAPIExtensions    bool `json:"include_api_extensions"`
	IncludeGeneric          bool `json:"include_generic"`
	IncludeAddons           bool `json:"include_addons"`
	IncludeOSSOnly          bool `json:"include_oss_only"`
	IncludePodLogs          bool `json:"include_pod_logs"`
	IncludeRBAC             bool `json:"include_rbac"`
	IncludeCloudFoundry     bool `json:"include_cloudfoundry"`
	IncludeMultiAZ          bool `json:"include_multiaz"`
	IncludeK8SLB            bool `json:"include_k8s_lb"`
	IncludePersistentVolume bool `json:"include_persistent_volume"`
}

type Kubernetes struct {
	AuthorizationMode string `json:"authorization_mode"`
	MasterHost        string `json:"master_host"`
	MasterPort        int    `json:"master_port"`
	PathToKubeConfig  string `json:"path_to_kube_config"`
	TLSCert           string `json:"tls_cert"`
	TLSPrivateKey     string `json:"tls_private_key"`
}

type CFCR struct {
	DeploymentPath           string `json:"deployment_path"`
	UpgradeToStemcellVersion string `json:"upgrade_to_stemcell_version"`
}

type Conformance struct {
	ResultsDir     string `json:"results_dir"`
	ReleaseVersion string `json:"release_version"`
}

func InitConfig() (*Config, error) {
	var config Config
	var configPath = os.Getenv("CONFIG")

	if configPath == "" {
		return nil, errors.New("config path must be defined")
	}

	configJSON, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return nil, err
	}

	// Do not allow zero for timeout scale as it would fail all the time.
	if config.TimeoutScale == 0 {
		config.TimeoutScale = 1
	}

	return &config, nil
}
