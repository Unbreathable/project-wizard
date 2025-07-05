package game

type Action struct {
	ID          uint // Per character
	Name        string
	Description string
	Element     *Element // Element of the action (may be nil)
	Damage      int      // The damage it deals
	ManaCost    int      // The mana it costs to cast
	Oversight   bool     // If the action is allowed to be executed from characters other than the one attached to

	// For easier time coding simulation
	originCharacter *Character
	targetCharacter *Character
	latestResult    ActionResult

	Before  func(current *Character, target *Character)              // Gets called before the simulation runs
	Execute func(current *Character, target *Character) ActionResult // Gets called to actually execute the action
}

// Create a new action that deals damage on execution.
func NewDamageAction(id uint, name string, description string, damage int, element *Element) Action {
	return Action{
		ID:          id,
		Name:        name,
		Description: description,
		Element:     element,
		Damage:      damage,
		Oversight:   false,
		Execute: func(current, target *Character) ActionResult {
			return ActionResult{
				DamageToCharacter: &damage,
			}
		},
	}
}

type ActionResult struct {
	DamageToCharacter *int
}
