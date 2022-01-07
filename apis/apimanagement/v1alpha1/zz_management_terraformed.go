/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Management
func (mg *Management) GetTerraformResourceType() string {
	return "azurerm_api_management"
}

// GetConnectionDetailsMapping for this Management
func (tr *Management) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate[*].certificate_password": "spec.forProvider.certificate[*].certificatePasswordSecretRef", "certificate[*].encoded_certificate": "spec.forProvider.certificate[*].encodedCertificateSecretRef", "hostname_configuration[*].developer_portal[*].certificate": "spec.forProvider.hostnameConfiguration[*].developerPortal[*].certificateSecretRef", "hostname_configuration[*].developer_portal[*].certificate_password": "spec.forProvider.hostnameConfiguration[*].developerPortal[*].certificatePasswordSecretRef", "hostname_configuration[*].management[*].certificate": "spec.forProvider.hostnameConfiguration[*].management[*].certificateSecretRef", "hostname_configuration[*].management[*].certificate_password": "spec.forProvider.hostnameConfiguration[*].management[*].certificatePasswordSecretRef", "hostname_configuration[*].portal[*].certificate": "spec.forProvider.hostnameConfiguration[*].portal[*].certificateSecretRef", "hostname_configuration[*].portal[*].certificate_password": "spec.forProvider.hostnameConfiguration[*].portal[*].certificatePasswordSecretRef", "hostname_configuration[*].proxy[*].certificate": "spec.forProvider.hostnameConfiguration[*].proxy[*].certificateSecretRef", "hostname_configuration[*].proxy[*].certificate_password": "spec.forProvider.hostnameConfiguration[*].proxy[*].certificatePasswordSecretRef", "hostname_configuration[*].scm[*].certificate": "spec.forProvider.hostnameConfiguration[*].scm[*].certificateSecretRef", "hostname_configuration[*].scm[*].certificate_password": "spec.forProvider.hostnameConfiguration[*].scm[*].certificatePasswordSecretRef", "tenant_access[*].primary_key": "status.atProvider.tenantAccess[*].primaryKey", "tenant_access[*].secondary_key": "status.atProvider.tenantAccess[*].secondaryKey"}
}

// GetObservation of this Management
func (tr *Management) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Management
func (tr *Management) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Management
func (tr *Management) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Management
func (tr *Management) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Management
func (tr *Management) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Management using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Management) LateInitialize(attrs []byte) (bool, error) {
	params := &ManagementParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Management) GetTerraformSchemaVersion() int {
	return 0
}
