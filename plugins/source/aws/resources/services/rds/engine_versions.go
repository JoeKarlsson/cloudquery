// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_engine_versions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html`,
		Resolver:    fetchRdsEngineVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "custom_db_engine_version_manifest",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomDBEngineVersionManifest"),
			},
			{
				Name:     "db_engine_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBEngineDescription"),
			},
			{
				Name:     "db_engine_media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBEngineMediaType"),
			},
			{
				Name:     "db_engine_version_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBEngineVersionArn"),
			},
			{
				Name:     "db_engine_version_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBEngineVersionDescription"),
			},
			{
				Name:     "db_parameter_group_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupFamily"),
			},
			{
				Name:     "database_installation_files_s3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseInstallationFilesS3BucketName"),
			},
			{
				Name:     "database_installation_files_s3_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseInstallationFilesS3Prefix"),
			},
			{
				Name:     "default_character_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultCharacterSet"),
			},
			{
				Name:     "exportable_log_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExportableLogTypes"),
			},
			{
				Name:     "image",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Image"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KMSKeyId"),
			},
			{
				Name:     "major_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MajorEngineVersion"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "supported_ca_certificate_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedCACertificateIdentifiers"),
			},
			{
				Name:     "supported_character_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SupportedCharacterSets"),
			},
			{
				Name:     "supported_engine_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedEngineModes"),
			},
			{
				Name:     "supported_feature_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedFeatureNames"),
			},
			{
				Name:     "supported_nchar_character_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SupportedNcharCharacterSets"),
			},
			{
				Name:     "supported_timezones",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SupportedTimezones"),
			},
			{
				Name:     "supports_babelfish",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsBabelfish"),
			},
			{
				Name:     "supports_certificate_rotation_without_restart",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsCertificateRotationWithoutRestart"),
			},
			{
				Name:     "supports_global_databases",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsGlobalDatabases"),
			},
			{
				Name:     "supports_log_exports_to_cloudwatch_logs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsLogExportsToCloudwatchLogs"),
			},
			{
				Name:     "supports_parallel_query",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsParallelQuery"),
			},
			{
				Name:     "supports_read_replica",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsReadReplica"),
			},
			{
				Name:     "tag_list",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
			{
				Name:     "valid_upgrade_target",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ValidUpgradeTarget"),
			},
		},

		Relations: []*schema.Table{
			ClusterParameters(),
		},
	}
}
