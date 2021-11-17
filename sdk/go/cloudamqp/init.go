// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "cloudamqp:index/alarm:Alarm":
		r = &Alarm{}
	case "cloudamqp:index/customDomain:CustomDomain":
		r = &CustomDomain{}
	case "cloudamqp:index/instance:Instance":
		r = &Instance{}
	case "cloudamqp:index/integrationLog:IntegrationLog":
		r = &IntegrationLog{}
	case "cloudamqp:index/integrationMetric:IntegrationMetric":
		r = &IntegrationMetric{}
	case "cloudamqp:index/notification:Notification":
		r = &Notification{}
	case "cloudamqp:index/plugin:Plugin":
		r = &Plugin{}
	case "cloudamqp:index/pluginCommunity:PluginCommunity":
		r = &PluginCommunity{}
	case "cloudamqp:index/securityFirewall:SecurityFirewall":
		r = &SecurityFirewall{}
	case "cloudamqp:index/vpcPeering:VpcPeering":
		r = &VpcPeering{}
	case "cloudamqp:index/webhook:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:cloudamqp" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/alarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/customDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/integrationLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/integrationMetric",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/notification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/plugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/pluginCommunity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/securityFirewall",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/vpcPeering",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"cloudamqp",
		"index/webhook",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"cloudamqp",
		&pkg{version},
	)
}
