/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package connector_utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"time"

	gouuid "github.com/satori/go.uuid"
	"log"
	"net"
)

// NewDatabaseObjectFromPrincipal creates a new object with default values, out of principle object
// func NewDatabaseObjectFromPrincipal(principal interface{}, refType string) *DatabaseObject {
// 	// Get user object
// 	Key, _ := PrincipalMarshalling(principal)

// 	// Generate DatabaseObject without JSON-object in it.
// 	key := NewDatabaseObject(Key.Uuid, refType)

// 	return key
// }

// PrincipalMarshalling Marhshall and Unmarshall Principal and Principals Objects
func PrincipalMarshalling(Object interface{}) (Key, KeyPermissions) {
	// marshall principal
	principalMarshall, _ := json.Marshal(Object)
	var key Key
	json.Unmarshal(principalMarshall, &key)

	return key, key.Permissions
}

// CreateFirstUserObject creates a new user with new API key when none exists when starting server
func CreateFirstUserObject() Key {
	key := Key{}

	// Create key token
	key.KeyToken = fmt.Sprintf("%v", gouuid.NewV4())

	// Uuid + name
	uuid := fmt.Sprintf("%v", gouuid.NewV4())

	// Auto set the parent ID to root *
	key.Parent = "*"

	// Set Uuid
	key.UUID = uuid

	// Set expiry to unlimited
	key.KeyExpiresUnix = -1

	// Set chmod variables
	keyPermissions := KeyPermissions{}
	keyPermissions.Read = true
	keyPermissions.Write = true
	keyPermissions.Delete = true
	keyPermissions.Execute = true

	// Get ips as v6
	var ips []string
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipv6 := ip.To16()
			ips = append(ips, ipv6.String())
		}
	}

	keyPermissions.IPOrigin = ips

	// Marshall and add to object
	key.Permissions = keyPermissions

	// Print the key
	log.Println("INFO: No root key was found, a new root key is created. More info: https://github.com/weaviate/weaviate/blob/develop/README.md#authentication")
	log.Println("INFO: Auto set allowed IPs to: ", keyPermissions.IPOrigin)
	log.Println("ROOTKEY=" + key.KeyToken)

	return key
}

// NowUnix returns the current Unix time
func NowUnix() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GenerateUUID returns a new UUID
func GenerateUUID() strfmt.UUID {
	return strfmt.UUID(fmt.Sprintf("%v", gouuid.NewV4()))
}