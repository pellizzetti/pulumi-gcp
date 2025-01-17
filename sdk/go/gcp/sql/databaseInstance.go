// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Google SQL Database Instance. For more information, see the [official documentation](https://cloud.google.com/sql/),
// or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/instances).
// 
// > **NOTE on `sql.DatabaseInstance`:** - Second-generation instances include a
// default 'root'@'%' user with no password. This user will be deleted by this provider on
// instance creation. You should use `sql.User` to define a custom user with
// a restricted host and strong password.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_database_instance.html.markdown.
type DatabaseInstance struct {
	s *pulumi.ResourceState
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOpt) (*DatabaseInstance, error) {
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["databaseVersion"] = nil
		inputs["masterInstanceName"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["replicaConfiguration"] = nil
		inputs["rootPassword"] = nil
		inputs["settings"] = nil
	} else {
		inputs["databaseVersion"] = args.DatabaseVersion
		inputs["masterInstanceName"] = args.MasterInstanceName
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["replicaConfiguration"] = args.ReplicaConfiguration
		inputs["rootPassword"] = args.RootPassword
		inputs["settings"] = args.Settings
	}
	inputs["connectionName"] = nil
	inputs["firstIpAddress"] = nil
	inputs["ipAddresses"] = nil
	inputs["privateIpAddress"] = nil
	inputs["publicIpAddress"] = nil
	inputs["selfLink"] = nil
	inputs["serverCaCert"] = nil
	inputs["serviceAccountEmailAddress"] = nil
	s, err := ctx.RegisterResource("gcp:sql/databaseInstance:DatabaseInstance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatabaseInstance{s: s}, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatabaseInstanceState, opts ...pulumi.ResourceOpt) (*DatabaseInstance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["connectionName"] = state.ConnectionName
		inputs["databaseVersion"] = state.DatabaseVersion
		inputs["firstIpAddress"] = state.FirstIpAddress
		inputs["ipAddresses"] = state.IpAddresses
		inputs["masterInstanceName"] = state.MasterInstanceName
		inputs["name"] = state.Name
		inputs["privateIpAddress"] = state.PrivateIpAddress
		inputs["project"] = state.Project
		inputs["publicIpAddress"] = state.PublicIpAddress
		inputs["region"] = state.Region
		inputs["replicaConfiguration"] = state.ReplicaConfiguration
		inputs["rootPassword"] = state.RootPassword
		inputs["selfLink"] = state.SelfLink
		inputs["serverCaCert"] = state.ServerCaCert
		inputs["serviceAccountEmailAddress"] = state.ServiceAccountEmailAddress
		inputs["settings"] = state.Settings
	}
	s, err := ctx.ReadResource("gcp:sql/databaseInstance:DatabaseInstance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatabaseInstance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DatabaseInstance) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DatabaseInstance) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The connection name of the instance to be used in
// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
func (r *DatabaseInstance) ConnectionName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionName"])
}

// The MySQL, PostgreSQL or MS SQL Server (beta) version to
// use. Can be `MYSQL_5_6`, `MYSQL_5_7`, `POSTGRES_9_6` or `POSTGRES_11` (beta) for second-generation
// instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
// MS SQL Server supported versions: `SQLSERVER_2017_STANDARD`, `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`, `SQLSERVER_ENTERPRISE_2016`.
// See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
// for more information.
func (r *DatabaseInstance) DatabaseVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["databaseVersion"])
}

func (r *DatabaseInstance) FirstIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["firstIpAddress"])
}

func (r *DatabaseInstance) IpAddresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipAddresses"])
}

// The name of the instance that will act as
// the master in the replication setup. Note, this requires the master to have
// `binaryLogEnabled` set, as well as existing backups.
func (r *DatabaseInstance) MasterInstanceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["masterInstanceName"])
}

// The name of the instance. If the name is left
// blank, this provider will randomly generate one when the instance is first
// created. This is done because after a name is used, it cannot be reused for
// up to [one week](https://cloud.google.com/sql/docs/delete-instance).
func (r *DatabaseInstance) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The first private (`PRIVATE`) IPv4 address assigned. This provides a convenient way to access an IP of a specific type without
// performing filtering.
func (r *DatabaseInstance) PrivateIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["privateIpAddress"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *DatabaseInstance) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The first public (`PRIMARY`) IPv4 address assigned. This provides a convenient way to access an IP of a specific type without
// performing filtering.
func (r *DatabaseInstance) PublicIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["publicIpAddress"])
}

// The region the instance will sit in. Note, first-generation Cloud SQL instance
// regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
// the provider region will be used instead, but this will be an apply-time error for all first-generation
// instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
// If you choose not to provide the `region` argument for this resource, make sure you understand this.
func (r *DatabaseInstance) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The configuration for replication. The
// configuration is detailed below.
func (r *DatabaseInstance) ReplicaConfiguration() pulumi.Output {
	return r.s.State["replicaConfiguration"]
}

// ) Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
func (r *DatabaseInstance) RootPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rootPassword"])
}

// The URI of the created resource.
func (r *DatabaseInstance) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *DatabaseInstance) ServerCaCert() pulumi.Output {
	return r.s.State["serverCaCert"]
}

// The service account email address assigned to the
// instance. This property is applicable only to Second Generation instances.
func (r *DatabaseInstance) ServiceAccountEmailAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceAccountEmailAddress"])
}

// The settings to use for the database. The
// configuration is detailed below.
func (r *DatabaseInstance) Settings() pulumi.Output {
	return r.s.State["settings"]
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type DatabaseInstanceState struct {
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName interface{}
	// The MySQL, PostgreSQL or MS SQL Server (beta) version to
	// use. Can be `MYSQL_5_6`, `MYSQL_5_7`, `POSTGRES_9_6` or `POSTGRES_11` (beta) for second-generation
	// instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
	// MS SQL Server supported versions: `SQLSERVER_2017_STANDARD`, `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`, `SQLSERVER_ENTERPRISE_2016`.
	// See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
	// for more information.
	DatabaseVersion interface{}
	FirstIpAddress interface{}
	IpAddresses interface{}
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName interface{}
	// The name of the instance. If the name is left
	// blank, this provider will randomly generate one when the instance is first
	// created. This is done because after a name is used, it cannot be reused for
	// up to [one week](https://cloud.google.com/sql/docs/delete-instance).
	Name interface{}
	// The first private (`PRIVATE`) IPv4 address assigned. This provides a convenient way to access an IP of a specific type without
	// performing filtering.
	PrivateIpAddress interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The first public (`PRIMARY`) IPv4 address assigned. This provides a convenient way to access an IP of a specific type without
	// performing filtering.
	PublicIpAddress interface{}
	// The region the instance will sit in. Note, first-generation Cloud SQL instance
	// regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for all first-generation
	// instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
	// If you choose not to provide the `region` argument for this resource, make sure you understand this.
	Region interface{}
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration interface{}
	// ) Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword interface{}
	// The URI of the created resource.
	SelfLink interface{}
	ServerCaCert interface{}
	// The service account email address assigned to the
	// instance. This property is applicable only to Second Generation instances.
	ServiceAccountEmailAddress interface{}
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings interface{}
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// The MySQL, PostgreSQL or MS SQL Server (beta) version to
	// use. Can be `MYSQL_5_6`, `MYSQL_5_7`, `POSTGRES_9_6` or `POSTGRES_11` (beta) for second-generation
	// instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
	// MS SQL Server supported versions: `SQLSERVER_2017_STANDARD`, `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`, `SQLSERVER_ENTERPRISE_2016`.
	// See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
	// for more information.
	DatabaseVersion interface{}
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName interface{}
	// The name of the instance. If the name is left
	// blank, this provider will randomly generate one when the instance is first
	// created. This is done because after a name is used, it cannot be reused for
	// up to [one week](https://cloud.google.com/sql/docs/delete-instance).
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region the instance will sit in. Note, first-generation Cloud SQL instance
	// regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for all first-generation
	// instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
	// If you choose not to provide the `region` argument for this resource, make sure you understand this.
	Region interface{}
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration interface{}
	// ) Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword interface{}
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings interface{}
}
