package products.dashboard.policy1

# By default, respond as negative.
default subject_has_product = false

# Tell if the subject's products contains the product used in the context.
subject_has_product {
	input.subject.products[_] == input.context.product
}

default subject_is_support = false

subject_is_support {
	input.subject.memberOf[_] == input.context.group
}
