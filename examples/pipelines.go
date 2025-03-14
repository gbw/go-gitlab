//
// Copyright 2021, Sander van Harmelen
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
//

package main

import (
	"log"
	"time"

	"gitlab.com/gitlab-org/api/client-go"
)

func pipelineExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectPipelinesOptions{
		Scope:         gitlab.Ptr("branches"),
		Status:        gitlab.Ptr(gitlab.Running),
		Ref:           gitlab.Ptr("master"),
		YamlErrors:    gitlab.Ptr(true),
		Name:          gitlab.Ptr("name"),
		Username:      gitlab.Ptr("username"),
		UpdatedAfter:  gitlab.Ptr(time.Now().Add(-24 * 365 * time.Hour)),
		UpdatedBefore: gitlab.Ptr(time.Now().Add(-7 * 24 * time.Hour)),
		OrderBy:       gitlab.Ptr("status"),
		Sort:          gitlab.Ptr("asc"),
	}

	pipelines, _, err := git.Pipelines.ListProjectPipelines(2743054, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("Found pipeline: %v", pipeline)
	}
}
