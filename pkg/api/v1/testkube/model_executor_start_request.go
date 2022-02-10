/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// scripts execution request body when calling new execution request
type ExecutorStartRequest struct {
	// ID of script execution to handle by executor, execution need to be able to return execution info based on this ID
	Id string `json:"id,omitempty"`
	// script type
	Type_ string `json:"type,omitempty"`
	// script execution custom name
	Name string `json:"name,omitempty"`
	// execution params passed to executor
	Params  map[string]string `json:"params,omitempty"`
	Content *TestContent      `json:"content,omitempty"`
}
