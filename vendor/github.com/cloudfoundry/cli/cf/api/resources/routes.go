package resources

import "github.com/cloudfoundry/cli/cf/models"

type RouteResource struct {
	Resource
	Entity RouteEntity
}

type RouteEntity struct {
	Host            string                  `json:"host"`
	Domain          DomainResource          `json:"domain"`
	Path            string                  `json:"path"`
	Port            int                     `json:"port"`
	Space           SpaceResource           `json:"space"`
	Apps            []ApplicationResource   `json:"apps"`
	ServiceInstance ServiceInstanceResource `json:"service_instance"`
}

func (resource RouteResource) ToFields() (fields models.Route) {
	fields.Guid = resource.Metadata.Guid
	fields.Host = resource.Entity.Host
	return
}

func (resource RouteResource) ToModel() (route models.Route) {
	route.Host = resource.Entity.Host
	route.Path = resource.Entity.Path
	route.Port = resource.Entity.Port
	route.Guid = resource.Metadata.Guid
	route.Domain = resource.Entity.Domain.ToFields()
	route.Space = resource.Entity.Space.ToFields()
	route.ServiceInstance = resource.Entity.ServiceInstance.ToFields()
	for _, appResource := range resource.Entity.Apps {
		route.Apps = append(route.Apps, appResource.ToFields())
	}
	return
}
