/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"testing"

	etcdintegration "github.com/coreos/etcd/integration"
)

// EtcdTestServer encapsulates the datastructures needed to start local instance for testing
type EtcdTestServer struct {
	Cluster *etcdintegration.ClusterV3
}

// Terminate will shutdown the running etcd server
func (m *EtcdTestServer) Terminate(t *testing.T) {
	m.Cluster.Terminate(t)
}

// NewEtcdTestClientServer creates a new client and server for testing
func NewEtcdTestClientServer(t *testing.T) *EtcdTestServer {
	clus := etcdintegration.NewClusterV3(t, &etcdintegration.ClusterConfig{Size: 1})
	return &EtcdTestServer{clus}
}
