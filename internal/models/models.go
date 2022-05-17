package models

type Champion struct {
	ID                     string   `db:"id" json:"id"`
	TitleHolder            string   `db:"title_holder" json:"titleHolder"`
	TitleHolderNumber      int32    `db:"title_holder_number" json:"titleHolderNumber"`
	DateWon                NullTime `db:"date_won" json:"dateWon"`
	DateLost               NullTime `db:"date_lost" json:"dateLost"`
	Show                   string   `db:"show" json:"show"`
	CurrentChampion        bool     `db:"current_champion" json:"currentChampion"`
	TitleHolderOrderNumber int32    `db:"title_holder_order_number" json:"titleHolderOrderNumber"`
	WrestlerID             int32    `db:"wrestler_id" json:"wrestler_id"`
}

type Champions []*Champion

type ChampionList struct {
	Champions Champions
}

type CreateChampionInput struct {
	TitleHolder            string
	TitleHolderNumber      int32
	TitleHolderOrderNumber int32
	WrestlerID             int32
	DateWon                NullTime
	Show                   string
}

type UpdateChampionInput struct {
	ID              string
	CurrentChampion bool
	DateLost        NullTime
}

type YearInput struct {
	StartDate NullTime
	EndDate   NullTime
}
