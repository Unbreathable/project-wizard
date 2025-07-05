package game

type Element string

// All the elements
const (
	ElementFire  Element = "fire"
	ElementWater Element = "water"
	ElementAir   Element = "air"
	ElementEarth Element = "earth"
	ElementDark  Element = "dark"
	ElementLight Element = "light"
	ElementNone  Element = "none"
)

func (e Element) Ptr() *Element {
	return &e
}

// Get the damage multiplier for when another element is attacking with e2
func (e1 Element) GetDamageMultiplierFor(e2 Element) float64 {
	switch e1 {
	// Fire <- e2
	case ElementFire:
		switch e2 {
		case ElementWater:
			return 1.5
		case ElementAir:
			return 0.8
		case ElementEarth:
			return 1.0
		case ElementDark:
			return 1.1
		case ElementLight:
			return 1.1
		}

	// Water <- ...
	case ElementWater:
		switch e2 {
		case ElementFire:
			return 0.5
		case ElementEarth:
			return 0.8
		case ElementDark:
			return 1.1
		case ElementLight:
			return 1.1
		}

	// Air <- ...
	case ElementAir:
		switch e2 {
		case ElementFire:
			return 0.8
		case ElementDark:
			return 1.1
		case ElementLight:
			return 1.1
		}

	// Earth <- ...
	case ElementEarth:
		switch e2 {
		case ElementFire:
			return 1.1
		case ElementAir:
			return 0.8
		case ElementDark:
			return 1.1
		case ElementLight:
			return 1.1
		}

	// Dark <- ...
	case ElementDark:
		switch e2 {
		case ElementLight:
			return 1.5
		}

	// Light <- ...
	case ElementLight:
		switch e2 {
		case ElementDark:
			return 1.5
		}
	}

	return 1.0
}
