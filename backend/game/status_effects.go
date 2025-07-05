package game

type StatusEffect struct {
	ID          string // Only needed when visible
	Name        string // Only needed when visible
	Description string // Only needed when visible
	Visible     bool

	OnHit func(current *Character, from *Character, action *Action) // Apply the status effect when the character gets hit
}
