package patreon

// Value = price, in cents. Patreon only sends us price, not tier name / ID
type Tier int

const(
	Premium Tier = 150
	WhitelabelBasic Tier = 300
)

var Tiers = []Tier{
	Premium,
	WhitelabelBasic,
}

func GetTierIndex(tier Tier) int {
	for i, t := range Tiers {
		if t == tier {
			return i
		}
	}

	return -1
}
