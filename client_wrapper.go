package dictributedConfigLibrary

import (
	"dictributedConfigLibrary/internal"
	"dictributedConfigLibrary/internal/delivery/proto"
	"dictributedConfigLibrary/internal/entity"
	"log"

	"google.golang.org/grpc"
	"strconv"
)

type ConfigServiceWrapper struct {
	client *internal.ConfigServiceClient
}

func NewConfigServiceWrapper() *ConfigServiceWrapper {
	return &ConfigServiceWrapper{}
}

func (c *ConfigServiceWrapper) Connect(port int, address string) error {
	conn, err := grpc.Dial(address+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to config service: %v", err)
		return err
	}
	client := proto.NewConfigServiceClient(conn)
	c.client = &internal.ConfigServiceClient{Client: client}
	log.Println("Successfully connected to config service")
	return nil
}

func (c *ConfigServiceWrapper) CreateConfig(serviceName string, data map[string]string) (*entity.Config, error) {
	config := &proto.Config{
		ServiceName: serviceName,
		Data:        data,
	}
	response, err := c.client.CreateConfig(config)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Name:      response.Config.ServiceName,
		Data:      response.Config.Data,
		CreatedAt: response.CreatedAt.AsTime(),
		Version:   response.Version,
	}, nil
}

func (c *ConfigServiceWrapper) GetConfig(serviceName string) (*entity.Config, error) {
	response, err := c.client.GetConfig(serviceName)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Name:      response.Config.ServiceName,
		Data:      response.Config.Data,
		CreatedAt: response.CreatedAt.AsTime(),
		Version:   response.Version,
	}, nil
}

func (c *ConfigServiceWrapper) GetConfigByVersion(serviceName string, version int64) (*entity.Config, error) {
	response, err := c.client.GetConfigByVersion(serviceName, version)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Name:      response.Config.ServiceName,
		Data:      response.Config.Data,
		CreatedAt: response.CreatedAt.AsTime(),
		Version:   response.Version,
	}, nil
}

func (c *ConfigServiceWrapper) UpdateConfig(serviceName string, data map[string]string) (*entity.Config, error) {
	config := &proto.Config{
		ServiceName: serviceName,
		Data:        data,
	}
	response, err := c.client.UpdateConfig(config)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Name:      response.Config.ServiceName,
		Data:      response.Config.Data,
		CreatedAt: response.CreatedAt.AsTime(),
		Version:   response.Version,
	}, nil
}

func (c *ConfigServiceWrapper) DeleteConfig(serviceName string) (string, error) {
	response, err := c.client.DeleteConfig(serviceName)
	return response.Message, err
}

func (c *ConfigServiceWrapper) DeleteConfigVersion(serviceName string, version int64) (string, error) {
	response, err := c.client.DeleteConfigVersion(serviceName, version)
	return response.Message, err
}

func (c *ConfigServiceWrapper) GetConfigHistory(serviceName string) ([]*entity.Config, error) {
	stream, err := c.client.ListConfig(serviceName)
	if err != nil {
		return nil, err
	}
	var configs []*entity.Config
	for {
		response, err := stream.Recv()
		if err != nil {
			break
		}
		configs = append(configs, &entity.Config{
			Name:      response.Config.ServiceName,
			Data:      response.Config.Data,
			CreatedAt: response.CreatedAt.AsTime(),
			Version:   response.Version,
		})
	}
	return configs, nil
}

func (c *ConfigServiceWrapper) SetRelevantConfig(serviceName string, version int64) (*entity.Config, error) {
	response, err := c.client.SetRelevantConfig(serviceName, version)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Name:      response.Config.ServiceName,
		Data:      response.Config.Data,
		CreatedAt: response.CreatedAt.AsTime(),
		Version:   response.Version,
	}, nil
}
