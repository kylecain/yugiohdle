package model

type Card struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	FrameType string `json:"frameType"`
	Desc      string `json:"desc"`
	Atk       int    `json:"atk"`
	Def       int    `json:"def"`
	Level     int    `json:"level"`
	Race      string `json:"race"`
	Attribute string `json:"attribute"`
	Card_sets []struct {
		Set_name        string `json:"set_name"`
		Set_code        string `json:"set_code"`
		Set_rarity      string `json:"set_rarity"`
		Set_rarity_code string `json:"set_rarity_code"`
		Set_price       string `json:"set_price"`
	} `json:"card_sets"`
	Card_images []struct {
		Id                int    `json:"id"`
		Image_url         string `json:"image_url"`
		Image_url_small   string `json:"image_url_small"`
		Image_url_cropped string `json:"image_url_cropped"`
	} `json:"card_images"`
	Card_prices []struct {
		Cardmarket_price   string `json:"cardmarket_price"`
		Tcgplayer_price    string `json:"tcgplayer_price"`
		Ebay_price         string `json:"ebay_price"`
		Amazon_price       string `json:"amazon_price"`
		Coolstuffinc_price string `json:"coolstuffinc_price"`
	} `json:"card_prices"`
}
