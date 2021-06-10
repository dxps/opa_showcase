package products_enablement

# By default, respond as negative.
default subject_has_product = false

# Tell if the subject's products contains the product used in the context.
subject_has_product {
	input.subject.products[_] == input.context.product
}
