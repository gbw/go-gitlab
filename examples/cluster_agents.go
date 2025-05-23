//
// Copyright 2021, Timo Furrer <tuxtimo@gmail.com>
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

// This example demonstrates how to manage Cluster Agents in GitLab using the GitLab API client.
// It covers the following steps:
// 1. Initialize the GitLab client with a personal access token.
// 2. Register a new Cluster Agent for a specific project by providing its name.
// 3. List all Cluster Agents associated with the specified project and log their details.

package main

import (
	"log"

	"gitlab.com/gitlab-org/api/client-go"
)

func clusterAgentsExample() {
	git, err := gitlab.NewClient("tokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	projectID := 33
	opt := &gitlab.RegisterAgentOptions{
		Name: gitlab.Ptr("agent-2"),
	}

	// Register Cluster Agent
	clusterAgent, _, err := git.ClusterAgents.RegisterAgent(projectID, opt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cluster Agent: %+v\n", clusterAgent)

	// List Cluster Agents
	clusterAgents, _, err := git.ClusterAgents.ListAgents(projectID, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cluster Agents: %+v", clusterAgents)
}
