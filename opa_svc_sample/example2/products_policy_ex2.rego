package products.policy.ex2

# By default, respond as negative.
default subject_has_product = false

# Tell if the subject's products (provided as `input.subject.products`) 
# contains the product used in the context (provided as `input.context.product`).
subject_has_product {
	input.subject.products[_] == input.context.product
}
