package internal

import (
	"cfglib/internal/delivery/proto"
	"context"
)

type ConfigServiceClient struct {
	Client proto.ConfigServiceClient
}

func NewConfigServiceClient(client proto.ConfigServiceClient) *ConfigServiceClient {
	return &ConfigServiceClient{
		Client: client,
	}
}

func (c *ConfigServiceClient) CreateConfig(config *proto.Config) (*proto.ConfigResponse, error) {
	return c.Client.CreateConfig(context.Background(), config)
}

func (c *ConfigServiceClient) GetConfig(serviceName string) (*proto.ConfigResponse, error) {
	return c.Client.GetConfig(context.Background(), &proto.ConfigName{ServiceName: serviceName})
}

func (c *ConfigServiceClient) GetConfigByVersion(serviceName string, version int64) (*proto.ConfigResponse, error) {
	return c.Client.GetConfigByVersion(context.Background(), &proto.ConfigNameAndVersion{ServiceName: serviceName, Version: version})
}

func (c *ConfigServiceClient) UpdateConfig(config *proto.Config) (*proto.ConfigResponse, error) {
	return c.Client.UpdateConfig(context.Background(), config)
}

func (c *ConfigServiceClient) DeleteConfig(serviceName string) (*proto.DeleteResponse, error) {
	return c.Client.DeleteConfig(context.Background(), &proto.ConfigName{ServiceName: serviceName})
}

func (c *ConfigServiceClient) DeleteConfigVersion(serviceName string, version int64) (*proto.DeleteResponse, error) {
	return c.Client.DeleteConfigVersion(context.Background(), &proto.ConfigNameAndVersion{ServiceName: serviceName, Version: version})
}

func (c *ConfigServiceClient) ListConfig(serviceName string) (proto.ConfigService_ListConfigsClient, error) {
	return c.Client.ListConfigs(context.Background(), &proto.ListRequest{ServiceName: serviceName})
}

func (c *ConfigServiceClient) SetRelevantConfig(serviceName string, version int64) (*proto.ConfigResponse, error) {
	return c.Client.SetRelevantConfig(context.Background(), &proto.ConfigNameAndVersion{ServiceName: serviceName, Version: version})
}
