//
// Copyright 2023, Joel Gerber
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
	"fmt"
	"log"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func dataDogExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Create new DataDog integration
	opts := &gitlab.SetDataDogServiceOptions{
		APIKey:             gitlab.Ptr("testing"),
		DataDogEnv:         gitlab.Ptr("sandbox"),
		DataDogService:     gitlab.Ptr("test"),
		DataDogSite:        gitlab.Ptr("datadoghq.com"),
		DataDogTags:        gitlab.Ptr("country:canada\nprovince:ontario"),
		ArchiveTraceEvents: gitlab.Ptr(true),
	}

	svc, _, err := git.Services.SetDataDogService(1, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"api_url: %s, datadog_env: %s, datadog_service: %s, datadog_site: %s, datadog_tags: %s",
		svc.Properties.APIURL, svc.Properties.DataDogEnv, svc.Properties.DataDogService,
		svc.Properties.DataDogSite, svc.Properties.DataDogTags,
	)

	// Delete the integration
	_, err = git.Services.DeleteDataDogService(1)
	if err != nil {
		log.Fatal(err)
	}
}
