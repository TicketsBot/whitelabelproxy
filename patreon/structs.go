package patreon

type(
	Pledge struct {
		Data Data
		Included []map[string]interface{}
	}

	Data struct {
		Attributes DataAttributes
	}

	DataAttributes struct {
		Amount int `json:"pledge_amount_cents"`
	}

	User struct {
		Attributes UserAttributes
	}

	UserAttributes struct {
		DiscordId string `json:"discord_id"`
		Email string
	}
)
