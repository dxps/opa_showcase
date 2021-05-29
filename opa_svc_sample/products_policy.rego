package products.policy

import data.products.acl

# By default, deny requests.
default user_has_product = false

# Tell if the user (as `input.user`) has the product (as `input.product`).
user_has_product {
	user_prods := acl[input.user]

	user_prods[_] == input.product
}
