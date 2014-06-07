package bot

import "tweet-kcwidget/model/tweet"

var roster = map[int]map[int]string{
	tweet.TypeMission: {
		0:  "Yudachi_kcw",
		1:  "Kaga_kcw",
		2:  "Akagi_kcw",
		3:  "Shioi_kcw",
		4:  "Shimakaze_kcw",
		5:  "Hibiki_kcw",
		6:  "Inazuma_kcw",
		7:  "Ikazuchi_kcw",
		8:  "Akatsuki_kcw",
		9:  "Makigumo_kcw",
		10: "Ryujo_kcw",
		11: "Tenryu_kcw",
		12: "Shigure_kcw",
		13: "Murasame_kcw",
		14: "Shiratsuyu_kcw",
		15: "Kagerou_kcw",
		16: "Shiranui_kcw",
		17: "Kuma_kcw",
		18: "Tama_kcw",
		19: "Kiso_kcw",
		20: "Shokaku_kcw",
		21: "Zuikaku_kcw",
		22: "Yubari_kcw",
		23: "Zuiho_kcw",
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
