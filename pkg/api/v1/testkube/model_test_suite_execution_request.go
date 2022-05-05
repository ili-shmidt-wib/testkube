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

// test suite execution request body
type TestSuiteExecutionRequest struct {
	// test execution custom name
	Name string `json:"name,omitempty"`
	// test kubernetes namespace (\"testkube\" when not set)
	Namespace string               `json:"namespace,omitempty"`
	Params    *map[string]Variable `json:"params,omitempty"`
	// execution labels
	Labels map[string]string `json:"labels,omitempty"`
}
