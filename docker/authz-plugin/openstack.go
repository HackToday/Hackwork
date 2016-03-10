package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type keystone struct {
	cfgHash map[string]string
}


func (ks *keystone) userName() string {
	return ks.cfgHash["userName"]	
}

func (ks *keystone) password() string {
	return ks.cfgHash["password"]
}

func (ks *keystone) domainName() string {
	return ks.cfgHash["domainName"]
}

func (ks *keystone) projectName() string {
	return ks.cfgHash["projectName"]
}

func (ks *keystone) regionName() string {
	return ks.cfgHash["regionName"]
}

func (ks *keystone) authUrl() string {
	return ks.cfgHash["authUrl"]
}

func (ks *keystone) authVersion() string {
	return ks.cfgHash["authVersion"]
}

func init() {

}

func (ks *keystone) getKsClient() *gophercloud.ServiceClient {

	ksVersion := ks.authVersion()
	if ksVersion == "v2.0" {
		logrus.Infof("Keystone is set to work with v2.0")	
	} else if ksVersion == "v3" {
		logrus.Infof("Keystone is set to work with v3")
	} else {
		logrus.Error("Wrong keystone version configuration")
		os.Exit(1)
	}

	// TODO: v2 and v3 authtentication flow
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: ks.authUrl(),
		Username: ks.userName(),
		Password: ks.password(),
		TenantName: ks.projectName(),
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	client := openstack.NewIdentityV2(provider)
	return client		
}
