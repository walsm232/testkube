/*
 * Kubetest
 *
 * Efficient testing of k8s applications mandates a k8s native approach to test mgmt/definition/execution - kubetest provides a “quality control plane” that natively integrates testing activities into k8s development and operational workflows
 *
 * API version: 0.0.1
 * Contact: api@kubetest.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kubetest
import (
	"time"
)

type Execution struct {
	// execution id (UUID?)
	Id string `json:"id,omitempty"`
	// script type e.g. postman/collection
	ScriptType string `json:"script-type,omitempty"`
	// script metadata content
	ScriptContent string `json:"script-content,omitempty"`
	// execution status
	Status string `json:"status,omitempty"`
	// RAW Script execution output, depends of reporter used in particular tool
	Output string `json:"output,omitempty"`
	// output type depends of reporter used in partucular tool
	OutputType string `json:"output-type,omitempty"`
	// error message when status is failed, separate to output as output can be partial in case of error
	ErrorMessage string `json:"error-message,omitempty"`
	// test start time
	StartTime time.Time `json:"start-time,omitempty"`
	// test end time
	EndTime time.Time `json:"end-time,omitempty"`
}
