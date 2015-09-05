Example:

```go
import "github.com/Joey92/gomasks"

const (
	USER_ROLE_A gomasks.Bitmask = 1 << iota
	USER_ROLE_B
	USER_ROLE_C
)

func BitmaskInAction() {

	var Roles gomasks.Bitmask

	// Add flags
	Roles.AddFlag(USER_ROLE_A)
	Roles.AddFlag(USER_ROLE_B)
	Roles.AddFlag(USER_ROLE_C)

	fmt.Println(Roles) // 7

	// Check flags
	Roles.HasFlag(USER_ROLE_A) // true
	Roles.HasFlag(USER_ROLE_B) // true

	// Remove flags
	Roles.RemoveFlag(USER_ROLE_C)

	Roles.HasFlag(USER_ROLE_C) // false

	// All stored in an uint32!
	fmt.Println(Roles) // 3

}
```