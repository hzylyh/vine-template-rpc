package schema

import (
	orderData "vine-template-rpc/internal/order/data/schema"
	systemData "vine-template-rpc/internal/system/data/schema"
)

var Schemas = []interface{}{
	&Site{},
	&Equipment{},

	// order
	&orderData.Order{},

	// system
	&systemData.User{},
}
