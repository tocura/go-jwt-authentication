package enum

type EnumRole string

const (
	RolePremium EnumRole = "premium"
	RoleNormal  EnumRole = "normal"
)

func MapToEnumRole(role string) EnumRole {
	m := map[string]EnumRole{
		"premium": RolePremium,
		"normal":  RoleNormal,
	}

	return m[role]
}
