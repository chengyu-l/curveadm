/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-12-24
 * Author: Jingli Chen (Wine93)
 */

package topology

import (
	"fmt"

	"github.com/opencurve/curveadm/internal/utils"
)

const (
	REQUIRE_ANY = iota
	REQUIRE_INT
	REQUIRE_STRING
	REQUIRE_BOOL
	REQUIRE_POSITIVE_INTEGER

	// default value
	DEFAULT_SSH_PORT                        = 22
	DEFAULT_SSH_TIMEOUT_SECONDS             = 10
	DEFAULT_REPORT_USAGE                    = true
	DEFAULT_CURVEBS_CONTAINER_IMAGE         = "opencurvedocker/curvebs:latest"
	DEFAULT_CURVEFS_CONTAINER_IMAGE         = "opencurvedocker/curvefs:latest"
	DEFAULT_ETCD_LISTEN_PEER_PORT           = 2380
	DEFAULT_ETCD_LISTEN_CLIENT_PORT         = 2379
	DEFAULT_MDS_LISTEN_PORT                 = 6700
	DEFAULT_MDS_LISTEN_DUMMY_PORT           = 7700
	DEFAULT_CHUNKSERVER_LISTN_PORT          = 8200
	DEFAULT_SNAPSHOTCLONE_LISTEN_PORT       = 5555
	DEFAULT_SNAPSHOTCLONE_LISTEN_DUMMY_PORT = 8081
	DEFAULT_SNAPSHOTCLONE_LISTEN_PROXY_PORT = 8080
	DEFAULT_METASERVER_LISTN_PORT           = 6800
	DEFAULT_METASERVER_LISTN_EXTARNAL_PORT  = 7800
	DEFAULT_CHUNKSERVER_COPYSETS            = 100 // copysets per chunkserver
	DEFAULT_METASERVER_COPYSETS             = 100 // copysets per metaserver
)

type (
	// config item
	item struct {
		key          string
		require      int
		exclude      bool        // exclude for service config
		defaultValue interface{} // nil means no default value
	}

	itemSet struct {
		items    []*item
		key2item map[string]*item
	}
)

// you should add config item to itemset iff you want to:
//   (1) check the configuration item value, like type, valid value OR
//   (2) filter out the configuration item for service config OR
//   (3) set the default value for configuration item
var (
	itemset = &itemSet{
		items:    []*item{},
		key2item: map[string]*item{},
	}

	CONFIG_USER = itemset.insert(
		"user",
		REQUIRE_STRING,
		true,
		func(dc *DeployConfig) interface{} {
			return utils.GetCurrentUser()
		},
	)

	CONFIG_SSH_PORT = itemset.insert(
		"ssh_port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		DEFAULT_SSH_PORT,
	)

	CONFIG_PRIVATE_CONFIG_FILE = itemset.insert(
		"private_key_file",
		REQUIRE_STRING,
		true,
		func(dc *DeployConfig) interface{} {
			return fmt.Sprintf("/home/%s/.ssh/id_rsa", dc.GetUser())
		},
	)

	CONFIG_REPORT_USAGE = itemset.insert(
		"report_usage",
		REQUIRE_BOOL,
		true,
		DEFAULT_REPORT_USAGE,
	)

	CONFIG_CONTAINER_IMAGE = itemset.insert(
		"container_image",
		REQUIRE_STRING,
		true,
		func(dc *DeployConfig) interface{} {
			if dc.GetKind() == KIND_CURVEBS {
				return DEFAULT_CURVEBS_CONTAINER_IMAGE
			}
			return DEFAULT_CURVEFS_CONTAINER_IMAGE
		},
	)

	CONFIG_LOG_DIR = itemset.insert(
		"log_dir",
		REQUIRE_STRING,
		true,
		nil,
	)

	CONFIG_DATA_DIR = itemset.insert(
		"data_dir",
		REQUIRE_STRING,
		true,
		nil,
	)

	CONFIG_CORE_DIR = itemset.insert(
		"core_dir",
		REQUIRE_STRING,
		true,
		nil,
	)

	CONFIG_LISTEN_IP = itemset.insert(
		"listen.ip",
		REQUIRE_STRING,
		true,
		func(dc *DeployConfig) interface{} {
			return dc.GetHost()
		},
	)

	CONFIG_LISTEN_PORT = itemset.insert(
		"listen.port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		func(dc *DeployConfig) interface{} {
			switch dc.GetRole() {
			case ROLE_ETCD:
				return DEFAULT_ETCD_LISTEN_PEER_PORT
			case ROLE_MDS:
				return DEFAULT_MDS_LISTEN_PORT
			case ROLE_CHUNKSERVER:
				return DEFAULT_CHUNKSERVER_LISTN_PORT
			case ROLE_SNAPSHOTCLONE:
				return DEFAULT_SNAPSHOTCLONE_LISTEN_PORT
			case ROLE_METASERVER:
				return DEFAULT_METASERVER_LISTN_PORT
			}
			return nil
		},
	)

	CONFIG_LISTEN_CLIENT_PORT = itemset.insert(
		"listen.client_port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		DEFAULT_ETCD_LISTEN_CLIENT_PORT,
	)

	CONFIG_LISTEN_DUMMY_PORT = itemset.insert(
		"listen.dummy_port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		func(dc *DeployConfig) interface{} {
			switch dc.GetRole() {
			case ROLE_MDS:
				return DEFAULT_MDS_LISTEN_DUMMY_PORT
			case ROLE_SNAPSHOTCLONE:
				return DEFAULT_SNAPSHOTCLONE_LISTEN_DUMMY_PORT
			}
			return nil
		},
	)

	CONFIG_LISTEN_PROXY_PORT = itemset.insert(
		"listen.proxy_port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		DEFAULT_SNAPSHOTCLONE_LISTEN_PROXY_PORT,
	)

	CONFIG_LISTEN_EXTERNAL_IP = itemset.insert(
		"listen.external_ip",
		REQUIRE_STRING,
		true,
		func(dc *DeployConfig) interface{} {
			return dc.GetHost()
		},
	)

	CONFIG_LISTEN_EXTERNAL_PORT = itemset.insert(
		"listen.external_port",
		REQUIRE_POSITIVE_INTEGER,
		true,
		func(dc *DeployConfig) interface{} {
			if dc.GetRole() == ROLE_METASERVER {
				return DEFAULT_METASERVER_LISTN_EXTARNAL_PORT
			}
			return dc.GetListenPort()
		},
	)

	CONFIG_COPYSETS = itemset.insert(
		"copysets",
		REQUIRE_POSITIVE_INTEGER,
		true,
		func(dc *DeployConfig) interface{} {
			if dc.GetRole() == ROLE_CHUNKSERVER {
				return DEFAULT_CHUNKSERVER_COPYSETS
			}
			return DEFAULT_METASERVER_COPYSETS
		},
	)

	CONFIG_VARIABLE = itemset.insert(
		"variable",
		REQUIRE_STRING,
		true,
		nil,
	)
)

func (i *item) Key() string {
	return i.key
}

func (itemset *itemSet) insert(key string, require int, exclude bool, defaultValue interface{}) *item {
	i := &item{key, require, exclude, defaultValue}
	itemset.key2item[key] = i
	itemset.items = append(itemset.items, i)
	return i
}

func (itemset *itemSet) get(key string) *item {
	return itemset.key2item[key]
}

func (itemset *itemSet) getAll() []*item {
	return itemset.items
}
