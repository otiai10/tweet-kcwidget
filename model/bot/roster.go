package bot

import "tweet-kcwidget/model/tweet"

var roster = map[int]map[int]string{
	tweet.TypeMission: {
		0:  "Yudachi_kcw",
		3:  "Kaga_kcw",
		6:  "Akagi_kcw",
		9:  "401_kcw",
		12: "Shimakaze_kcw",
		15: "Hibiki_kcw",
		18: "Ryujo_kcw",
		21: "Tenryu_kcw",
	},
	tweet.TypeNyukyo: {
		0:  "Mogami_kcw",
		4:  "Sakawa_kcw",
		8:  "Aoba_kcw",
		12: "Kinugasa_kcw",
		16: "Nagato_kcw",
		20: "Kikuduki_kcw",
	},
	tweet.TypeCreateship: {
		0:  "Kongo_kcw",
		6:  "Kirishima_kcw",
		12: "Haruna_kcw",
		18: "Hiei_kcw",
	},
	tweet.TypeSortie: {
		0:  "Sazanami_kcw",
		4:  "Ushio_kcw",
		8:  "Akebono_kcw",
		12: "Oboro_kcw",
		16: "Yukikaze_kcw",
		20: "Amatsukaze_kcw",
	},
}
