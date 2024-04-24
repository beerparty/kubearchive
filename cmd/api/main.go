// Copyright KubeArchive Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FIXME This will be taken from a shared pkg with sink based on the DB schema
// Just for the first approach
type resources struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Items      []string `json:"items"`
}

// getAllResources responds with the list of resources of a specific type across all namespaces
func getAllResources(c *gin.Context) {
	group := c.Param("group")
	version := c.Param("version")
	pluralResourceType := c.Param("resourceType")
	response := resources{
		// FIXME The plural of the resource type can be different of adding an `s` to the kind
		Kind:       pluralResourceType[:len(pluralResourceType)-1],
		APIVersion: fmt.Sprintf("%s/%s", group, version),
		Items:      []string{"resource1", "resource2"}}
	c.IndentedJSON(http.StatusOK, response)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/apis/:group/:version/:resourceType", getAllResources)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8081")
}