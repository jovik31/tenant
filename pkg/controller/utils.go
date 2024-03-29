package controller

import (
	"fmt"
	"log"
	"github.com/jovik31/tenant/pkg/apis/jovik31.dev/v1alpha1"
	tenantType "github.com/jovik31/tenant/pkg/apis/jovik31.dev/v1alpha1"
)


type EventObject struct {

	eventType string
	key string
	newObj interface{}
	oldObj interface{}
}


func existsNode(nodeList []tenantType.Node, currentNodeName string) bool {

	for _, element := range nodeList {
		if currentNodeName == element.Name {
			fmt.Print("Node is part of tenant")

			return true
		}
	}
	return false

}

func (c *Controller) refreshTenant(namespace string, name string) (*v1alpha1.Tenant, error){

	tenant, err:=c.tenantLister.Tenants(namespace).Get(name)
	if err!=nil{
		log.Println("Failed retrieving newest tenant", err.Error())
	}
	return tenant, nil
}
