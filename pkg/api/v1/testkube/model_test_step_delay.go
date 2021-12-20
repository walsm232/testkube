/*
 * TestKube API
 *
 * TestKube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

type TestStepDelay struct {
	// step name (in case of script it'll be script name)
	Name              string `json:"name"`
	StopTestOnFailure bool   `json:"stopTestOnFailure,omitempty"`
	// delay duration in milliseconds
	Duration int32 `json:"duration"`
}