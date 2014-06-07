package bot

import "tweet-kcwidget/model/tweet"

var roster = map[int]map[int][]string{
	// 狗肉の策です
	tweet.TypeMission: {
		0:  []string{"Yudachi_kcw", "Kaga_kcw", "Akagi_kcw"},
		1:  []string{"Kaga_kcw", "Akagi_kcw", "Shioi_kcw"},
		2:  []string{"Akagi_kcw", "Shioi_kcw", "Shimakaze_kcw"},
		3:  []string{"Shioi_kcw", "Shimakaze_kcw", "Hibiki_kcw"},
		4:  []string{"Shimakaze_kcw", "Hibiki_kcw", "Inazuma_kcw"},
		5:  []string{"Hibiki_kcw", "Inazuma_kcw", "Ikazuchi_kcw"},
		6:  []string{"Inazuma_kcw", "Ikazuchi_kcw", "Akatsuki_kcw"},
		7:  []string{"Ikazuchi_kcw", "Akatsuki_kcw", "Makigumo_kcw"},
		8:  []string{"Akatsuki_kcw", "Makigumo_kcw", "Ryujo_kcw"},
		9:  []string{"Makigumo_kcw", "Ryujo_kcw", "Tenryu_kcw"},
		10: []string{"Ryujo_kcw", "Tenryu_kcw", "Shigure_kcw"},
		11: []string{"Tenryu_kcw", "Shigure_kcw", "Murasame_kcw"},
		12: []string{"Shigure_kcw", "Murasame_kcw", "Shiratsuyu_kcw"},
		13: []string{"Murasame_kcw", "Shiratsuyu_kcw", "Kagerou_kcw"},
		14: []string{"Shiratsuyu_kcw", "Kagerou_kcw", "Shiranui_kcw"},
		15: []string{"Kagerou_kcw", "Shiranui_kcw", "Kuma_kcw"},
		16: []string{"Shiranui_kcw", "Kuma_kcw", "Tama_kcw"},
		17: []string{"Kuma_kcw", "Tama_kcw", "Kiso_kcw"},
		18: []string{"Tama_kcw", "Kiso_kcw", "Shokaku_kcw"},
		19: []string{"Kiso_kcw", "Shokaku_kcw", "Zuikaku_kcw"},
		20: []string{"Shokaku_kcw", "Zuikaku_kcw", "Yubari_kcw"},
		21: []string{"Zuikaku_kcw", "Yubari_kcw", "Zuiho_kcw"},
		22: []string{"Yubari_kcw", "Zuiho_kcw", "Yudachi_kcw"},
		23: []string{"Zuiho_kcw", "Yudachi_kcw", "Kaga_kcw"},
	},
	tweet.TypeNyukyo: {
		0:  []string{"Mogami_kcw"},
		4:  []string{"Sakawa_kcw"},
		8:  []string{"Aoba_kcw"},
		12: []string{"Kinugasa_kcw"},
		16: []string{"Nagato_kcw"},
		20: []string{"Kikuduki_kcw"},
	},
	tweet.TypeCreateship: {
		0:  []string{"Kongo_kcw"},
		6:  []string{"Kirishima_kcw"},
		12: []string{"Haruna_kcw"},
		18: []string{"Hiei_kcw"},
	},
	tweet.TypeSortie: {
		0:  []string{"Sazanami_kcw"},
		4:  []string{"Ushio_kcw"},
		8:  []string{"Akebono_kcw"},
		12: []string{"Oboro_kcw"},
		16: []string{"Yukikaze_kcw"},
		20: []string{"Amatsukaze_kcw"},
	},
}
