// Copyright © 2019 Kerem Karatal
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Creator struct {
	Name        string
	Title       string
	Email       string
	GithubToken string
}

type Task struct {
	Level           int8
	Title           string
	DescriptionFile string
}

type Challenge struct {
	Discipline       string
	TemplateRepoName string
	Reviewers        []string
	Tasks            []Task
}

type Config struct {
	TrackingRepoName string
	Organization     string
	Owner            string
	Creator          Creator
	Challenges       []Challenge
}

var instance *Config
var once sync.Once

func GetConfigInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		err := viper.UnmarshalKey("config", instance)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})
	return instance
}

func (config *Config) FindChallenge(discipline string) (*Challenge, error) {
	for _, challenge := range config.Challenges {
		if challenge.Discipline == discipline {
			return &challenge, nil
		}
	}
	return nil, errors.New("Unknown discipline")
}
