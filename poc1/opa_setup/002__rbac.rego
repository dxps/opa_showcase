package rbac

# By default, respond as negative.
default subject_has_support_role = false

# Tell if the subject has the 'support' role based on a specific group membership.
subject_has_support_role {
	support_group := concat("@", ["support", input.context.product])
	trace(concat("=", ["support_group", support_group]))
	input.subject.memberOf[_] == support_group
}
