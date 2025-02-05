/*
 * Copyright (C) 2021 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package confgen

import (
	"github.com/netobserv/flowlogs2metrics/pkg/api"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Options struct {
	DestConfFile             string
	DestDocFile              string
	DestGrafanaJsonnetFolder string
	SrcFolder                string
	SkipWithLabels           []string
}

var (
	Opt = Options{}
)

type ConfigIngest struct {
	Collector api.IngestCollector `yaml:"collector"`
}

type ConfigTransform struct {
	Generic api.TransformGeneric `yaml:"generic"`
}

type ConfigEncode struct {
	Prom api.PromEncode `yaml:"prom"`
}

type ConfigWrite struct {
	Loki api.WriteLoki `yaml:"loki"`
	Type string        `yaml:"type"`
}

type ConfigVisualization struct {
	Grafana ConfigVisualizationGrafana `yaml:"grafana"`
}

type Config struct {
	Description   string              `yaml:"description"`
	Ingest        ConfigIngest        `yaml:"ingest"`
	Transform     ConfigTransform     `yaml:"transform"`
	Write         ConfigWrite         `yaml:"write"`
	Encode        ConfigEncode        `yaml:"encode"`
	Visualization ConfigVisualization `yaml:"visualization"`
}

func (cg *ConfGen) parseConfigFile(fileName string) (*Config, error) {
	// parse config file yaml
	var config Config
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debugf("ioutil.ReadFile err: %v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Debugf("Unmarshal err: %v ", err)
		return nil, err
	}

	return &config, nil
}
