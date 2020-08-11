package huaweicloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for HuaweiCloud.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_ACCESS_KEY", nil),
				Description:  descriptions["access_key"],
				RequiredWith: []string{"secret_key"},
			},

			"secret_key": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_SECRET_KEY", nil),
				Description:  descriptions["secret_key"],
				RequiredWith: []string{"access_key"},
			},

			"auth_url": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.EnvDefaultFunc(
					"OS_AUTH_URL", "https://iam.myhuaweicloud.com:443/v3"),
				Description: descriptions["auth_url"],
			},

			"region": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  descriptions["region"],
				DefaultFunc:  schema.EnvDefaultFunc("OS_REGION_NAME", nil),
				InputDefault: "cn-north-1",
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_TOKEN", ""),
				Description: descriptions["token"],
			},

			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
					"OS_DOMAIN_ID",
				}, ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_DEFAULT_DOMAIN",
				}, ""),
				Description: descriptions["domain_name"],
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", false),
				Description: descriptions["insecure"],
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"agency_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_NAME", nil),
				Description:  descriptions["agency_name"],
				RequiredWith: []string{"agency_domain_name"},
			},

			"agency_domain_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_DOMAIN_NAME", nil),
				Description:  descriptions["agency_domain_name"],
				RequiredWith: []string{"agency_name"},
			},
			"delegated_project": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELEGATED_PROJECT", ""),
				Description: descriptions["delegated_project"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"huaweicloud_availability_zones":            dataSourceAvailabilityZones(),
			"huaweicloud_compute_availability_zones_v2": dataSourceComputeAvailabilityZonesV2(),
			"huaweicloud_images_image_v2":               dataSourceImagesImageV2(),
			"huaweicloud_networking_network_v2":         dataSourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":          dataSourceNetworkingSubnetV2(),
			"huaweicloud_networking_port_v2":            dataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2":        dataSourceNetworkingSecGroupV2(),
			"huaweicloud_s3_bucket_object":              dataSourceS3BucketObject(),
			"huaweicloud_kms_key_v1":                    dataSourceKmsKeyV1(),
			"huaweicloud_kms_data_key_v1":               dataSourceKmsDataKeyV1(),
			"huaweicloud_gaussdb_mysql_flavors":         dataSourceGaussdbMysqlFlavors(),
			"huaweicloud_rds_flavors_v1":                dataSourceRdsFlavorV1(),
			"huaweicloud_rds_flavors_v3":                dataSourceRdsFlavorV3(),
			"huaweicloud_sfs_file_system_v2":            dataSourceSFSFileSystemV2(),
			"huaweicloud_rts_stack_v1":                  dataSourceRTSStackV1(),
			"huaweicloud_rts_stack_resource_v1":         dataSourceRTSStackResourcesV1(),
			"huaweicloud_iam_role_v3":                   dataSourceIAMRoleV3(),
			"huaweicloud_vpc_v1":                        dataSourceVirtualPrivateCloudVpcV1(),
			"huaweicloud_vpc_ids_v1":                    dataSourceVirtualPrivateCloudVpcIdsV1(),
			"huaweicloud_vpc_peering_connection_v2":     dataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_route_v2":                  dataSourceVPCRouteV2(),
			"huaweicloud_vpc_route_ids_v2":              dataSourceVPCRouteIdsV2(),
			"huaweicloud_vpc_subnet_v1":                 dataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids_v1":             dataSourceVpcSubnetIdsV1(),
			"huaweicloud_cce_cluster_v3":                dataSourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":                   dataSourceCCENodeV3(),
			"huaweicloud_rts_software_config_v1":        dataSourceRtsSoftwareConfigV1(),
			"huaweicloud_csbs_backup_v1":                dataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1":         dataSourceCSBSBackupPolicyV1(),
			"huaweicloud_dms_az_v1":                     dataSourceDmsAZV1(),
			"huaweicloud_dms_product_v1":                dataSourceDmsProductV1(),
			"huaweicloud_dms_maintainwindow_v1":         dataSourceDmsMaintainWindowV1(),
			"huaweicloud_vbs_backup_policy_v2":          dataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":                 dataSourceVBSBackupV2(),
			"huaweicloud_cts_tracker_v1":                dataSourceCTSTrackerV1(),
			"huaweicloud_antiddos_v1":                   dataSourceAntiDdosV1(),
			"huaweicloud_dcs_az_v1":                     dataSourceDcsAZV1(),
			"huaweicloud_dcs_maintainwindow_v1":         dataSourceDcsMaintainWindowV1(),
			"huaweicloud_dcs_product_v1":                dataSourceDcsProductV1(),
			"huaweicloud_dds_flavors_v3":                dataSourceDDSFlavorV3(),
			"huaweicloud_identity_role_v3":              dataSourceIdentityRoleV3(),
			"huaweicloud_cdm_flavors_v1":                dataSourceCdmFlavorV1(),
			"huaweicloud_dis_partition_v2":              dataSourceDisPartitionV2(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"huaweicloud_antiddos":                           resourceAntiDdosV1(),
			"huaweicloud_api_gateway_group":                  resourceAPIGatewayGroup(),
			"huaweicloud_api_gateway_api":                    resourceAPIGatewayAPI(),
			"huaweicloud_blockstorage_volume_v2":             resourceBlockStorageVolumeV2(),
			"huaweicloud_compute_instance_v2":                resourceComputeInstanceV2(),
			"huaweicloud_compute_interface_attach_v2":        resourceComputeInterfaceAttachV2(),
			"huaweicloud_compute_keypair_v2":                 resourceComputeKeypairV2(),
			"huaweicloud_compute_secgroup_v2":                resourceComputeSecGroupV2(),
			"huaweicloud_compute_servergroup_v2":             resourceComputeServerGroupV2(),
			"huaweicloud_compute_floatingip_v2":              resourceComputeFloatingIPV2(),
			"huaweicloud_compute_floatingip_associate_v2":    resourceComputeFloatingIPAssociateV2(),
			"huaweicloud_compute_eip_associate":              resourceComputeFloatingIPAssociateV2(),
			"huaweicloud_compute_volume_attach_v2":           resourceComputeVolumeAttachV2(),
			"huaweicloud_dns_ptrrecord_v2":                   resourceDNSPtrRecordV2(),
			"huaweicloud_dns_recordset_v2":                   resourceDNSRecordSetV2(),
			"huaweicloud_dns_zone_v2":                        resourceDNSZoneV2(),
			"huaweicloud_dcs_instance_v1":                    resourceDcsInstanceV1(),
			"huaweicloud_dds_instance_v3":                    resourceDdsInstanceV3(),
			"huaweicloud_ecs_instance_v1":                    resourceEcsInstanceV1(),
			"huaweicloud_evs_snapshot":                       resourceEvsSnapshotV2(),
			"huaweicloud_fw_firewall_group_v2":               resourceFWFirewallGroupV2(),
			"huaweicloud_fw_policy_v2":                       resourceFWPolicyV2(),
			"huaweicloud_fw_rule_v2":                         resourceFWRuleV2(),
			"huaweicloud_kms_key_v1":                         resourceKmsKeyV1(),
			"huaweicloud_dms_queue_v1":                       resourceDmsQueuesV1(),
			"huaweicloud_dms_group_v1":                       resourceDmsGroupsV1(),
			"huaweicloud_dms_instance_v1":                    resourceDmsInstancesV1(),
			"huaweicloud_lb_certificate_v2":                  resourceCertificateV2(),
			"huaweicloud_elb_loadbalancer":                   resourceELBLoadBalancer(),
			"huaweicloud_elb_listener":                       resourceELBListener(),
			"huaweicloud_elb_healthcheck":                    resourceELBHealthCheck(),
			"huaweicloud_elb_backendecs":                     resourceELBBackendECS(),
			"huaweicloud_images_image_v2":                    resourceImagesImageV2(),
			"huaweicloud_lb_loadbalancer_v2":                 resourceLoadBalancerV2(),
			"huaweicloud_lb_listener_v2":                     resourceListenerV2(),
			"huaweicloud_lb_pool_v2":                         resourcePoolV2(),
			"huaweicloud_lb_member_v2":                       resourceMemberV2(),
			"huaweicloud_lb_monitor_v2":                      resourceMonitorV2(),
			"huaweicloud_lb_l7policy_v2":                     resourceL7PolicyV2(),
			"huaweicloud_lb_l7rule_v2":                       resourceL7RuleV2(),
			"huaweicloud_lb_whitelist_v2":                    resourceWhitelistV2(),
			"huaweicloud_mrs_cluster_v1":                     resourceMRSClusterV1(),
			"huaweicloud_mrs_job_v1":                         resourceMRSJobV1(),
			"huaweicloud_networking_network_v2":              resourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":               resourceNetworkingSubnetV2(),
			"huaweicloud_networking_floatingip_v2":           resourceNetworkingFloatingIPV2(),
			"huaweicloud_networking_port_v2":                 resourceNetworkingPortV2(),
			"huaweicloud_networking_router_v2":               resourceNetworkingRouterV2(),
			"huaweicloud_networking_router_interface_v2":     resourceNetworkingRouterInterfaceV2(),
			"huaweicloud_networking_router_route_v2":         resourceNetworkingRouterRouteV2(),
			"huaweicloud_networking_secgroup_v2":             resourceNetworkingSecGroupV2(),
			"huaweicloud_networking_secgroup_rule_v2":        resourceNetworkingSecGroupRuleV2(),
			"huaweicloud_networking_floatingip_associate_v2": resourceNetworkingFloatingIPAssociateV2(),
			"huaweicloud_networking_eip_associate":           resourceNetworkingFloatingIPAssociateV2(),
			"huaweicloud_s3_bucket":                          resourceS3Bucket(),
			"huaweicloud_s3_bucket_policy":                   resourceS3BucketPolicy(),
			"huaweicloud_s3_bucket_object":                   resourceS3BucketObject(),
			"huaweicloud_obs_bucket":                         resourceObsBucket(),
			"huaweicloud_obs_bucket_object":                  resourceObsBucketObject(),
			"huaweicloud_obs_bucket_policy":                  resourceObsBucketPolicy(),
			"huaweicloud_smn_topic_v2":                       resourceTopic(),
			"huaweicloud_smn_subscription_v2":                resourceSubscription(),
			"huaweicloud_rds_instance_v1":                    resourceRdsInstance(),
			"huaweicloud_rds_instance_v3":                    resourceRdsInstanceV3(),
			"huaweicloud_rds_parametergroup_v3":              resourceRdsConfigurationV3(),
			"huaweicloud_gaussdb_cassandra_instance":         resourceGeminiDBInstanceV3(),
			"huaweicloud_gaussdb_mysql_instance":             resourceGaussDBInstance(),
			"huaweicloud_gaussdb_opengauss_instance":         resourceOpenGaussInstance(),
			"huaweicloud_nat_gateway_v2":                     resourceNatGatewayV2(),
			"huaweicloud_nat_snat_rule_v2":                   resourceNatSnatRuleV2(),
			"huaweicloud_nat_dnat_rule_v2":                   resourceNatDnatRuleV2(),
			"huaweicloud_sfs_file_system_v2":                 resourceSFSFileSystemV2(),
			"huaweicloud_rts_stack_v1":                       resourceRTSStackV1(),
			"huaweicloud_iam_agency_v3":                      resourceIAMAgencyV3(),
			"huaweicloud_vpc_v1":                             resourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_bandwidth_v2":                   resourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip_v1":                         resourceVpcEIPV1(),
			"huaweicloud_vpc_peering_connection_v2":          resourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter_v2": resourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route_v2":                       resourceVPCRouteV2(),
			"huaweicloud_vpc_subnet_v1":                      resourceVpcSubnetV1(),
			"huaweicloud_cce_cluster_v3":                     resourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":                        resourceCCENodeV3(),
			"huaweicloud_cci_network_v1":                     resourceCCINetworkV1(),
			"huaweicloud_rts_software_config_v1":             resourceSoftwareConfigV1(),
			"huaweicloud_ces_alarmrule":                      resourceAlarmRule(),
			"huaweicloud_as_configuration_v1":                resourceASConfiguration(),
			"huaweicloud_as_group_v1":                        resourceASGroup(),
			"huaweicloud_as_policy_v1":                       resourceASPolicy(),
			"huaweicloud_csbs_backup_v1":                     resourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1":              resourceCSBSBackupPolicyV1(),
			"huaweicloud_vbs_backup_policy_v2":               resourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":                      resourceVBSBackupV2(),
			"huaweicloud_cts_tracker_v1":                     resourceCTSTrackerV1(),
			"huaweicloud_maas_task_v1":                       resourceMaasTaskV1(),
			"huaweicloud_dws_cluster":                        resourceDwsCluster(),
			"huaweicloud_mls_instance":                       resourceMlsInstance(),
			"huaweicloud_identity_project_v3":                resourceIdentityProjectV3(),
			"huaweicloud_identity_role_assignment_v3":        resourceIdentityRoleAssignmentV3(),
			"huaweicloud_identity_user_v3":                   resourceIdentityUserV3(),
			"huaweicloud_identity_group_v3":                  resourceIdentityGroupV3(),
			"huaweicloud_identity_group_membership_v3":       resourceIdentityGroupMembershipV3(),
			"huaweicloud_cdm_cluster_v1":                     resourceCdmClusterV1(),
			"huaweicloud_ges_graph_v1":                       resourceGesGraphV1(),
			"huaweicloud_cloudtable_cluster_v2":              resourceCloudtableClusterV2(),
			"huaweicloud_css_cluster_v1":                     resourceCssClusterV1(),
			"huaweicloud_dis_stream_v2":                      resourceDisStreamV2(),
			"huaweicloud_cs_cluster_v1":                      resourceCsClusterV1(),
			"huaweicloud_cs_peering_connect_v1":              resourceCsPeeringConnectV1(),
			"huaweicloud_vpnaas_ipsec_policy_v2":             resourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service_v2":                  resourceVpnServiceV2(),
			"huaweicloud_vpnaas_ike_policy_v2":               resourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_endpoint_group_v2":           resourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_site_connection_v2":          resourceVpnSiteConnectionV2(),
			"huaweicloud_dli_queue_v1":                       resourceDliQueueV1(),
			"huaweicloud_cs_route_v1":                        resourceCsRouteV1(),
			"huaweicloud_networking_vip_v2":                  resourceNetworkingVIPV2(),
			"huaweicloud_networking_vip_associate_v2":        resourceNetworkingVIPAssociateV2(),
			"huaweicloud_fgs_function_v2":                    resourceFgsFunctionV2(),
			"huaweicloud_cdn_domain_v1":                      resourceCdnDomainV1(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return configureProvider(d, terraformVersion)
	}

	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The HuaweiCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"agency_name": "The name of agency",

		"agency_domain_name": "The name of domain who created the agency (Identity v3).",

		"delegated_project": "The name of delegated project (Identity v3).",
	}
}

func configureProvider(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	var tenant_name, delegated_project string

	// Use region as tenant_name if it's not set
	if v, ok := d.GetOk("tenant_name"); ok && v.(string) != "" {
		tenant_name = v.(string)
	} else {
		tenant_name = d.Get("region").(string)
	}

	// Use region as delegated_project if it's not set
	if v, ok := d.GetOk("delegated_project"); ok && v.(string) != "" {
		delegated_project = v.(string)
	} else {
		delegated_project = d.Get("region").(string)
	}

	config := Config{
		AccessKey:        d.Get("access_key").(string),
		SecretKey:        d.Get("secret_key").(string),
		CACertFile:       d.Get("cacert_file").(string),
		ClientCertFile:   d.Get("cert").(string),
		ClientKeyFile:    d.Get("key").(string),
		DomainID:         d.Get("domain_id").(string),
		DomainName:       d.Get("domain_name").(string),
		IdentityEndpoint: d.Get("auth_url").(string),
		Insecure:         d.Get("insecure").(bool),
		Password:         d.Get("password").(string),
		Region:           d.Get("region").(string),
		Token:            d.Get("token").(string),
		TenantID:         d.Get("tenant_id").(string),
		TenantName:       tenant_name,
		Username:         d.Get("user_name").(string),
		UserID:           d.Get("user_id").(string),
		AgencyName:       d.Get("agency_name").(string),
		AgencyDomainName: d.Get("agency_domain_name").(string),
		DelegatedProject: delegated_project,
		terraformVersion: terraformVersion,
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
