package patreon

import (
	"fmt"
	"github.com/TicketsBot/whitelabelproxy/database"
	"github.com/apex/log"
	"github.com/mitchellh/mapstructure"
)

func HandleCreate(pledge Pledge) {
	var user User
	for _, include := range pledge.Included {
		if include["type"] == "user" {
			if err := mapstructure.Decode(include, &user); err != nil {
				log.Error(err.Error())
				return
			}
		}
	}

	var tier Tier
	for _, t := range Tiers {
		if int(t) == pledge.Data.Attributes.Amount {
			tier = t
			break
		}
	}

	if tier == 0 {
		log.Error(fmt.Sprintf("Invalid pledge amount: %d", pledge.Data.Attributes.Amount))
		return
	}

	switch tier {
	case Premium: // TODO: Premium
	case WhitelabelBasic: {
		go database.Update(user.Attributes.Email, GetTierIndex(WhitelabelBasic), &user.Attributes.DiscordId)
	}
	}
}
