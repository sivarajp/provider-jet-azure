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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-azure/apis/aad/v1alpha1"
	v1alpha1aadiam "github.com/crossplane-contrib/provider-jet-azure/apis/aadiam/v1alpha1"
	v1alpha1alertsmanagement "github.com/crossplane-contrib/provider-jet-azure/apis/alertsmanagement/v1alpha1"
	v1alpha1analysisservices "github.com/crossplane-contrib/provider-jet-azure/apis/analysisservices/v1alpha1"
	v1alpha1apimanagement "github.com/crossplane-contrib/provider-jet-azure/apis/apimanagement/v1alpha1"
	v1alpha1appconfiguration "github.com/crossplane-contrib/provider-jet-azure/apis/appconfiguration/v1alpha1"
	v1alpha1appplatform "github.com/crossplane-contrib/provider-jet-azure/apis/appplatform/v1alpha1"
	v1alpha1attestation "github.com/crossplane-contrib/provider-jet-azure/apis/attestation/v1alpha1"
	v1alpha1authorization "github.com/crossplane-contrib/provider-jet-azure/apis/authorization/v1alpha1"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-azure/apis/authorization/v1alpha2"
	v1alpha1automation "github.com/crossplane-contrib/provider-jet-azure/apis/automation/v1alpha1"
	v1alpha1avs "github.com/crossplane-contrib/provider-jet-azure/apis/avs/v1alpha1"
	v1alpha1azure "github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1"
	v1alpha2azure "github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2"
	v1alpha1azurestackhci "github.com/crossplane-contrib/provider-jet-azure/apis/azurestackhci/v1alpha1"
	v1alpha1batch "github.com/crossplane-contrib/provider-jet-azure/apis/batch/v1alpha1"
	v1alpha1blueprint "github.com/crossplane-contrib/provider-jet-azure/apis/blueprint/v1alpha1"
	v1alpha1botservice "github.com/crossplane-contrib/provider-jet-azure/apis/botservice/v1alpha1"
	v1alpha2cache "github.com/crossplane-contrib/provider-jet-azure/apis/cache/v1alpha2"
	v1alpha1cdn "github.com/crossplane-contrib/provider-jet-azure/apis/cdn/v1alpha1"
	v1alpha1certificateregistration "github.com/crossplane-contrib/provider-jet-azure/apis/certificateregistration/v1alpha1"
	v1alpha1cognitiveservices "github.com/crossplane-contrib/provider-jet-azure/apis/cognitiveservices/v1alpha1"
	v1alpha1communication "github.com/crossplane-contrib/provider-jet-azure/apis/communication/v1alpha1"
	v1alpha1compute "github.com/crossplane-contrib/provider-jet-azure/apis/compute/v1alpha1"
	v1alpha1consumption "github.com/crossplane-contrib/provider-jet-azure/apis/consumption/v1alpha1"
	v1alpha1containerregistry "github.com/crossplane-contrib/provider-jet-azure/apis/containerregistry/v1alpha1"
	v1alpha2containerservice "github.com/crossplane-contrib/provider-jet-azure/apis/containerservice/v1alpha2"
	v1alpha2cosmosdb "github.com/crossplane-contrib/provider-jet-azure/apis/cosmosdb/v1alpha2"
	v1alpha1costmanagement "github.com/crossplane-contrib/provider-jet-azure/apis/costmanagement/v1alpha1"
	v1alpha1customproviders "github.com/crossplane-contrib/provider-jet-azure/apis/customproviders/v1alpha1"
	v1alpha1databoxedge "github.com/crossplane-contrib/provider-jet-azure/apis/databoxedge/v1alpha1"
	v1alpha1databricks "github.com/crossplane-contrib/provider-jet-azure/apis/databricks/v1alpha1"
	v1alpha1datafactory "github.com/crossplane-contrib/provider-jet-azure/apis/datafactory/v1alpha1"
	v1alpha1datalakeanalytics "github.com/crossplane-contrib/provider-jet-azure/apis/datalakeanalytics/v1alpha1"
	v1alpha1datalakestore "github.com/crossplane-contrib/provider-jet-azure/apis/datalakestore/v1alpha1"
	v1alpha1datamigration "github.com/crossplane-contrib/provider-jet-azure/apis/datamigration/v1alpha1"
	v1alpha1dataprotection "github.com/crossplane-contrib/provider-jet-azure/apis/dataprotection/v1alpha1"
	v1alpha1datashare "github.com/crossplane-contrib/provider-jet-azure/apis/datashare/v1alpha1"
	v1alpha1dbformariadb "github.com/crossplane-contrib/provider-jet-azure/apis/dbformariadb/v1alpha1"
	v1alpha1dbformysql "github.com/crossplane-contrib/provider-jet-azure/apis/dbformysql/v1alpha1"
	v1alpha2dbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/apis/dbforpostgresql/v1alpha2"
	v1alpha1devices "github.com/crossplane-contrib/provider-jet-azure/apis/devices/v1alpha1"
	v1alpha2devices "github.com/crossplane-contrib/provider-jet-azure/apis/devices/v1alpha2"
	v1alpha1devspaces "github.com/crossplane-contrib/provider-jet-azure/apis/devspaces/v1alpha1"
	v1alpha1devtestlab "github.com/crossplane-contrib/provider-jet-azure/apis/devtestlab/v1alpha1"
	v1alpha1digitaltwins "github.com/crossplane-contrib/provider-jet-azure/apis/digitaltwins/v1alpha1"
	v1alpha1eventgrid "github.com/crossplane-contrib/provider-jet-azure/apis/eventgrid/v1alpha1"
	v1alpha1eventhub "github.com/crossplane-contrib/provider-jet-azure/apis/eventhub/v1alpha1"
	v1alpha1hardwaresecuritymodules "github.com/crossplane-contrib/provider-jet-azure/apis/hardwaresecuritymodules/v1alpha1"
	v1alpha1hdinsight "github.com/crossplane-contrib/provider-jet-azure/apis/hdinsight/v1alpha1"
	v1alpha1healthbot "github.com/crossplane-contrib/provider-jet-azure/apis/healthbot/v1alpha1"
	v1alpha1healthcareapis "github.com/crossplane-contrib/provider-jet-azure/apis/healthcareapis/v1alpha1"
	v1alpha1insights "github.com/crossplane-contrib/provider-jet-azure/apis/insights/v1alpha1"
	v1alpha2insights "github.com/crossplane-contrib/provider-jet-azure/apis/insights/v1alpha2"
	v1alpha1iotcentral "github.com/crossplane-contrib/provider-jet-azure/apis/iotcentral/v1alpha1"
	v1alpha1keyvault "github.com/crossplane-contrib/provider-jet-azure/apis/keyvault/v1alpha1"
	v1alpha2keyvault "github.com/crossplane-contrib/provider-jet-azure/apis/keyvault/v1alpha2"
	v1alpha1kusto "github.com/crossplane-contrib/provider-jet-azure/apis/kusto/v1alpha1"
	v1alpha2loganalytics "github.com/crossplane-contrib/provider-jet-azure/apis/loganalytics/v1alpha2"
	v1alpha1logic "github.com/crossplane-contrib/provider-jet-azure/apis/logic/v1alpha1"
	v1alpha1machinelearningservices "github.com/crossplane-contrib/provider-jet-azure/apis/machinelearningservices/v1alpha1"
	v1alpha1maintenance "github.com/crossplane-contrib/provider-jet-azure/apis/maintenance/v1alpha1"
	v1alpha1managedidentity "github.com/crossplane-contrib/provider-jet-azure/apis/managedidentity/v1alpha1"
	v1alpha1managedservices "github.com/crossplane-contrib/provider-jet-azure/apis/managedservices/v1alpha1"
	v1alpha2management "github.com/crossplane-contrib/provider-jet-azure/apis/management/v1alpha2"
	v1alpha1maps "github.com/crossplane-contrib/provider-jet-azure/apis/maps/v1alpha1"
	v1alpha1marketplaceordering "github.com/crossplane-contrib/provider-jet-azure/apis/marketplaceordering/v1alpha1"
	v1alpha1media "github.com/crossplane-contrib/provider-jet-azure/apis/media/v1alpha1"
	v1alpha1mixedreality "github.com/crossplane-contrib/provider-jet-azure/apis/mixedreality/v1alpha1"
	v1alpha1netapp "github.com/crossplane-contrib/provider-jet-azure/apis/netapp/v1alpha1"
	v1alpha1network "github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha1"
	v1alpha2network "github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2"
	v1alpha1notificationhubs "github.com/crossplane-contrib/provider-jet-azure/apis/notificationhubs/v1alpha1"
	v1alpha1operationalinsights "github.com/crossplane-contrib/provider-jet-azure/apis/operationalinsights/v1alpha1"
	v1alpha1policyinsights "github.com/crossplane-contrib/provider-jet-azure/apis/policyinsights/v1alpha1"
	v1alpha1portal "github.com/crossplane-contrib/provider-jet-azure/apis/portal/v1alpha1"
	v1alpha1powerbidedicated "github.com/crossplane-contrib/provider-jet-azure/apis/powerbidedicated/v1alpha1"
	v1alpha1purview "github.com/crossplane-contrib/provider-jet-azure/apis/purview/v1alpha1"
	v1alpha1recoveryservices "github.com/crossplane-contrib/provider-jet-azure/apis/recoveryservices/v1alpha1"
	v1alpha1relay "github.com/crossplane-contrib/provider-jet-azure/apis/relay/v1alpha1"
	v1alpha1resources "github.com/crossplane-contrib/provider-jet-azure/apis/resources/v1alpha1"
	v1alpha2resources "github.com/crossplane-contrib/provider-jet-azure/apis/resources/v1alpha2"
	v1alpha1search "github.com/crossplane-contrib/provider-jet-azure/apis/search/v1alpha1"
	v1alpha1security "github.com/crossplane-contrib/provider-jet-azure/apis/security/v1alpha1"
	v1alpha1securityinsights "github.com/crossplane-contrib/provider-jet-azure/apis/securityinsights/v1alpha1"
	v1alpha1servicebus "github.com/crossplane-contrib/provider-jet-azure/apis/servicebus/v1alpha1"
	v1alpha1servicefabric "github.com/crossplane-contrib/provider-jet-azure/apis/servicefabric/v1alpha1"
	v1alpha1servicefabricmesh "github.com/crossplane-contrib/provider-jet-azure/apis/servicefabricmesh/v1alpha1"
	v1alpha1signalrservice "github.com/crossplane-contrib/provider-jet-azure/apis/signalrservice/v1alpha1"
	v1alpha1solutions "github.com/crossplane-contrib/provider-jet-azure/apis/solutions/v1alpha1"
	v1alpha1sql "github.com/crossplane-contrib/provider-jet-azure/apis/sql/v1alpha1"
	v1alpha2sql "github.com/crossplane-contrib/provider-jet-azure/apis/sql/v1alpha2"
	v1alpha1sqlvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/apis/sqlvirtualmachine/v1alpha1"
	v1alpha1storage "github.com/crossplane-contrib/provider-jet-azure/apis/storage/v1alpha1"
	v1alpha2storage "github.com/crossplane-contrib/provider-jet-azure/apis/storage/v1alpha2"
	v1alpha1storagecache "github.com/crossplane-contrib/provider-jet-azure/apis/storagecache/v1alpha1"
	v1alpha1storagesync "github.com/crossplane-contrib/provider-jet-azure/apis/storagesync/v1alpha1"
	v1alpha1streamanalytics "github.com/crossplane-contrib/provider-jet-azure/apis/streamanalytics/v1alpha1"
	v1alpha1synapse "github.com/crossplane-contrib/provider-jet-azure/apis/synapse/v1alpha1"
	v1alpha1timeseriesinsights "github.com/crossplane-contrib/provider-jet-azure/apis/timeseriesinsights/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-azure/apis/v1alpha1"
	v1alpha1web "github.com/crossplane-contrib/provider-jet-azure/apis/web/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1aadiam.SchemeBuilder.AddToScheme,
		v1alpha1alertsmanagement.SchemeBuilder.AddToScheme,
		v1alpha1analysisservices.SchemeBuilder.AddToScheme,
		v1alpha1apimanagement.SchemeBuilder.AddToScheme,
		v1alpha1appconfiguration.SchemeBuilder.AddToScheme,
		v1alpha1appplatform.SchemeBuilder.AddToScheme,
		v1alpha1attestation.SchemeBuilder.AddToScheme,
		v1alpha1authorization.SchemeBuilder.AddToScheme,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha1automation.SchemeBuilder.AddToScheme,
		v1alpha1avs.SchemeBuilder.AddToScheme,
		v1alpha1azure.SchemeBuilder.AddToScheme,
		v1alpha2azure.SchemeBuilder.AddToScheme,
		v1alpha1azurestackhci.SchemeBuilder.AddToScheme,
		v1alpha1batch.SchemeBuilder.AddToScheme,
		v1alpha1blueprint.SchemeBuilder.AddToScheme,
		v1alpha1botservice.SchemeBuilder.AddToScheme,
		v1alpha2cache.SchemeBuilder.AddToScheme,
		v1alpha1cdn.SchemeBuilder.AddToScheme,
		v1alpha1certificateregistration.SchemeBuilder.AddToScheme,
		v1alpha1cognitiveservices.SchemeBuilder.AddToScheme,
		v1alpha1communication.SchemeBuilder.AddToScheme,
		v1alpha1compute.SchemeBuilder.AddToScheme,
		v1alpha1consumption.SchemeBuilder.AddToScheme,
		v1alpha1containerregistry.SchemeBuilder.AddToScheme,
		v1alpha2containerservice.SchemeBuilder.AddToScheme,
		v1alpha2cosmosdb.SchemeBuilder.AddToScheme,
		v1alpha1costmanagement.SchemeBuilder.AddToScheme,
		v1alpha1customproviders.SchemeBuilder.AddToScheme,
		v1alpha1databoxedge.SchemeBuilder.AddToScheme,
		v1alpha1databricks.SchemeBuilder.AddToScheme,
		v1alpha1datafactory.SchemeBuilder.AddToScheme,
		v1alpha1datalakeanalytics.SchemeBuilder.AddToScheme,
		v1alpha1datalakestore.SchemeBuilder.AddToScheme,
		v1alpha1datamigration.SchemeBuilder.AddToScheme,
		v1alpha1dataprotection.SchemeBuilder.AddToScheme,
		v1alpha1datashare.SchemeBuilder.AddToScheme,
		v1alpha1dbformariadb.SchemeBuilder.AddToScheme,
		v1alpha1dbformysql.SchemeBuilder.AddToScheme,
		v1alpha2dbforpostgresql.SchemeBuilder.AddToScheme,
		v1alpha1devices.SchemeBuilder.AddToScheme,
		v1alpha2devices.SchemeBuilder.AddToScheme,
		v1alpha1devspaces.SchemeBuilder.AddToScheme,
		v1alpha1devtestlab.SchemeBuilder.AddToScheme,
		v1alpha1digitaltwins.SchemeBuilder.AddToScheme,
		v1alpha1eventgrid.SchemeBuilder.AddToScheme,
		v1alpha1eventhub.SchemeBuilder.AddToScheme,
		v1alpha1hardwaresecuritymodules.SchemeBuilder.AddToScheme,
		v1alpha1hdinsight.SchemeBuilder.AddToScheme,
		v1alpha1healthbot.SchemeBuilder.AddToScheme,
		v1alpha1healthcareapis.SchemeBuilder.AddToScheme,
		v1alpha1insights.SchemeBuilder.AddToScheme,
		v1alpha2insights.SchemeBuilder.AddToScheme,
		v1alpha1iotcentral.SchemeBuilder.AddToScheme,
		v1alpha1keyvault.SchemeBuilder.AddToScheme,
		v1alpha2keyvault.SchemeBuilder.AddToScheme,
		v1alpha1kusto.SchemeBuilder.AddToScheme,
		v1alpha2loganalytics.SchemeBuilder.AddToScheme,
		v1alpha1logic.SchemeBuilder.AddToScheme,
		v1alpha1machinelearningservices.SchemeBuilder.AddToScheme,
		v1alpha1maintenance.SchemeBuilder.AddToScheme,
		v1alpha1managedidentity.SchemeBuilder.AddToScheme,
		v1alpha1managedservices.SchemeBuilder.AddToScheme,
		v1alpha2management.SchemeBuilder.AddToScheme,
		v1alpha1maps.SchemeBuilder.AddToScheme,
		v1alpha1marketplaceordering.SchemeBuilder.AddToScheme,
		v1alpha1media.SchemeBuilder.AddToScheme,
		v1alpha1mixedreality.SchemeBuilder.AddToScheme,
		v1alpha1netapp.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha2network.SchemeBuilder.AddToScheme,
		v1alpha1notificationhubs.SchemeBuilder.AddToScheme,
		v1alpha1operationalinsights.SchemeBuilder.AddToScheme,
		v1alpha1policyinsights.SchemeBuilder.AddToScheme,
		v1alpha1portal.SchemeBuilder.AddToScheme,
		v1alpha1powerbidedicated.SchemeBuilder.AddToScheme,
		v1alpha1purview.SchemeBuilder.AddToScheme,
		v1alpha1recoveryservices.SchemeBuilder.AddToScheme,
		v1alpha1relay.SchemeBuilder.AddToScheme,
		v1alpha1resources.SchemeBuilder.AddToScheme,
		v1alpha2resources.SchemeBuilder.AddToScheme,
		v1alpha1search.SchemeBuilder.AddToScheme,
		v1alpha1security.SchemeBuilder.AddToScheme,
		v1alpha1securityinsights.SchemeBuilder.AddToScheme,
		v1alpha1servicebus.SchemeBuilder.AddToScheme,
		v1alpha1servicefabric.SchemeBuilder.AddToScheme,
		v1alpha1servicefabricmesh.SchemeBuilder.AddToScheme,
		v1alpha1signalrservice.SchemeBuilder.AddToScheme,
		v1alpha1solutions.SchemeBuilder.AddToScheme,
		v1alpha1sql.SchemeBuilder.AddToScheme,
		v1alpha2sql.SchemeBuilder.AddToScheme,
		v1alpha1sqlvirtualmachine.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha2storage.SchemeBuilder.AddToScheme,
		v1alpha1storagecache.SchemeBuilder.AddToScheme,
		v1alpha1storagesync.SchemeBuilder.AddToScheme,
		v1alpha1streamanalytics.SchemeBuilder.AddToScheme,
		v1alpha1synapse.SchemeBuilder.AddToScheme,
		v1alpha1timeseriesinsights.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1web.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
