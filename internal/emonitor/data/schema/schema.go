package schema

import orderData "vine-template-rpc/internal/order/data/schema"

var Schemas = []interface{}{
	&Site{},
	&Equipment{},

	// order
	&orderData.Order{},
}
