package products

import data.products.acl
import input

# default access = false

allow {
	user := input.user
	enablement := acl[user]
	count(enablement[input.product]) > 0
}
