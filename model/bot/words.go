package bot

var monologues = map[string][]string{
	"Yudachi_kcw": {
		"ぽい？",
		"ぽい！",
		"ぽい〜♪",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kaga_kcw": {
		"やりました",
		"ここは譲れません",
		"何か相談？いいけれど",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Akagi_kcw": {
		"ごっはん〜ごっはん〜...、あ違ったw ごめんなさいね",
		"慢心してはだめ",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shioi_kcw": {
		"どぼ〜ん",
		"んがとか！んがとか！！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shimakaze_kcw": {
		"おっ",
		"おっそーい！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Hibiki_kcw": {
		"不死鳥の名はだてじゃない",
		"ありがとう、спасибо",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Akatsuki_kcw": {
		"暁の水平線の...なんだっけ？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Ikazuchi_kcw": {
		"いかずちよ！かみなりじゃないわ！そこんとこもよろしくね！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Inazuma_kcw": {
		"いなずまのほんきをみるのです！",
		"なのです！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Makigumo_kcw": {
		"おやくだちです！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Mogami_kcw": {
		"今度は衝突しないって！ホントだよ！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Sakawa_kcw": {
		"ぴゃん♪",
		"ぴゃっ！？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Aoba_kcw": {
		"青葉取材！いえ出撃します！",
		"なになに？なんの話ですか？",
		"司令官、青葉、見ちゃいました！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kinugasa_kcw": {
		"衣笠さんにおまかせ！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Nagato_kcw": {
		"ビッグ7の力、侮るなよ。",
		"提督、いつもお疲れ様だな。今日くらいは一緒に飲もう。",
		"いいだろう。",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kikuduki_kcw": {
		"うぅっ……なんなのさ、一体……",
		"礼は言わぬ……この力、何に使うか……",
		"私は前線に居なくて良いのか……？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Sazanami_kcw": {
		"ktkr♪",
		"ご主人様、調子に乗ると、ぶっとばしますよ♪",
		"ほいさっさ～♪",
		"徹底的にやっちまうのねっ！",
		"はにゃ～っ！",
		"うっっくぅ～、なんもいえねぇ～…",
		"あぁ～、今回も～やられてしまいましたが～…あっ、それは違うの？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Ushio_kcw": {
		"ひやぁ！…も、もう、構わないでください…",
		"もしかして…呼びましたか？",
		"私でも…お役に立てたのでしょうか…ああっ、みんな見ないでください…恥ずかしいよぉ…",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Akebono_kcw": {
		"何で触るの？ウザイなぁ。",
		"敵ぃ～？ フフン♪そうこなくっちゃね！",
		"こっち見んな！この糞提督！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Oboro_kcw": {
		"提督のそういうとこ、キライ･･･じゃ、ないです。",
		"がんばる！",
		"実はお腹…空いてました…",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kongo_kcw": {
		"バーニングラアアアアーブ",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Haruna_kcw": {
		"はい、榛名は大丈夫です！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kirishima_kcw": {
		"よく出来ましたっ！",
		"ご命令を、司令。",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Hiei_kcw": {
		"ひえー",
		"ひえ〜",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Ryujo_kcw": {
		"いってみよう！",
		"ほっほー……ウチのこと、大切に思ってくれてるん？　それはちょっち嬉しいなぁ♪",
		"ありがとう♪これで赤城や加賀にも負けないかな？って……そりゃ無理か、あはははは……",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Yukikaze_kcw": {
		"しれぇ！",
		"ぜったい、だいじょうぶ！",
		"ゆきかじぇはしずみません！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Tenryu_kcw": {
		"フフフ、怖いか？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Amatsukaze_kcw": {
		"やだ、髪は触んないでよ…。吹流しが取れちゃうでしょ",
		"私の連装砲くんのほうが可愛いに決まってるでしょ？",
		"いい風",
		"お風呂よ、お風呂！汗臭いのはいや！",
		"はぁ～疲れた。帰っていい？",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shigure_kcw": {
		"いい雨だね",
		"雨は、いつかやむさ",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Murasame_kcw": {
		"はいは〜い",
		"むらさめだよー",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shiratsuyu_kcw": {
		"いっちばーん！",
		"はい、一番艦です！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kagerou_kcw": {
		"やっと会えた！陽炎よ。よろしくねっ！",
		"なーに？お話したいの？",
		"ちょっあまり触ると怒るわよ、もう",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shiranui_kcw": {
		"不知火です。ご指導ご鞭撻、よろしくです",
		"お呼びですか指令",
		"何でしょうか？不知火に落ち度でも？",
		"つまらないわね。もっと骨のある敵はいないの",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kuma_kcw": {
		"くまー",
		"くまー！",
		"くま？",
		"くま♪",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Tama_kcw": {
		"たまー",
		"たまー！",
		"たま？",
		"たま♪",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kiso_kcw": {
		"きそー",
		"きそー！",
		"きそ？",
		"きそ♪",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shokaku_kcw": {
		"瑞鶴、あまり提督の邪魔しちゃダメよ？",
		"提督…？あの、なんでしょう？",
		"瑞鶴…いいの？",
		"ちょっとお風呂に…すぐ出るから",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Zuikaku_kcw": {
		"提督さん…なに？作戦？",
		"翔鶴姉、なに？……って、提督さんじゃん？！なにやってんの！？爆撃されたいの！？",
		"全機爆装、準備出来しだい発艦！　目標、母港執務室の提督、ヤッちゃって！",
		"七面鳥ですって！？冗談じゃないわ！！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Yubari_kcw": {
		"何でしょう？試し撃ち、ご所望ですか？",
		"あぁ、提督もやっぱりそう思います？ここに兵装まだ載りそうよね‥、うん‥。",
		"出撃よ！ってやだ、私が一番遅いって…お、置いてかないでよぉ！",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Zuiho_kcw": {
		"瑞鳳です。軽空母ですが、錬度があがれば、正規空母並の活躍をおみせできます。",
		"九九艦爆は、脚が可愛いのよ、脚が。",
		"彗星は彗星で悪くないんだけれど、整備大変なのよー、整備が。",
		"天山はー…って、あれ？んっ、提督？格納庫まさぐるの止めてくれない？んっ。っていうか、邪魔っ",
		"いいかもね！",
		"やられちゃったなー。温泉、入りたいな。",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
}
