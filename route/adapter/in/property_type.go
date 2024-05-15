package in

import (
	"github.com/Vioneta/VionetaOS-MessageBus/codegen"
	"github.com/Vioneta/VionetaOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
