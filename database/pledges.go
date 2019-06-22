package database

type Pledge struct {
	Email     string `gorm:"column:EMAIL;type:VARCHAR(255);unique;primary_key"`
	Tier      int    `gorm:"column:TIER"`
	DiscordID *string `gorm:"column:DISCORDID;type:VARCHAR(21);nullable"`
}

func Update(email string, tier int, discordId *string) {
	var node Pledge
	Db.Where(Pledge{Email: email}).Assign(Pledge{Tier: tier, DiscordID: discordId}).FirstOrCreate(&node)
}

func GetPledge(email string, ch chan int) {
	var node Pledge
	Db.Where(Pledge{Email: email}).Take(&node)
	ch <- node.Tier
}

func GetDiscordID(email string, ch chan *string) {
	var node Pledge
	Db.Where(Pledge{Email: email}).Take(&node)
	ch <- node.DiscordID
}

func IsEmailPledged(email string, ch chan bool) {
	var count int
	Db.Where(Pledge{Email: email}).Count(&count)
	ch <- count > 0
}

func IsDiscordIDPledged(discordID *string, ch chan bool) {
	var count int
	Db.Where(Pledge{DiscordID: discordID}).Count(&count)
	ch <- count > 0
}
