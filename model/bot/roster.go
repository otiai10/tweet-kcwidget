package bot

import "tweet-kcwidget/model/tweet"

var roster = map[int]map[int]string{
	tweet.TypeMission: {
		0:  "Yudachi_kcw",
		4:  "Kaga_kcw",
		8:  "Akagi_kcw",
		12: "401_kcw",
		16: "Shimakaze_kcw",
		20: "Hibiki_kcw",
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
		0:  "Sazanami_kcw",
		4:  "Ushio_kcw",
		8:  "Akebono_kcw",
		12: "Makigumo_kcw",
		16: "Yukikaze_kcw",
		20: "Amatsukaze_kcw",
	},
	tweet.TypeSortie: {
		0:  "Kongo_kcw",
		4:  "Kirishima_kcw",
		8:  "Haruna_kcw",
		12: "Hiei_kcw",
		16: "Ryujo_kcw",
		20: "Tenryu_kcw",
	},
}
