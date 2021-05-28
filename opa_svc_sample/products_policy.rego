package products.policy

import data.products_acl

# By default, deny requests.
default allow = false

allow {
	user_has_product
}

# Tell if the user (from `input.user`) has the provided `product`.
user_has_product {
	user_prods := products_acl[input.user]

	some i

	# product is the `i`-th element in the user->products mappings for the provided user.
	user_prods[i] == input.product
}
