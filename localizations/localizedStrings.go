package localizations

import (
	"github.com/Mtbcooler/outrun/enums"
)

var LocalizedStrings = map[string]map[string]string{
	"en": map[string]string{
		"DailyBattleWinRewardLabel":          "A reward for winning a daily battle.",
		"DailyBattleWinStreakRewardLabel":    "A reward for winning %v daily battles in a row.",
		"DailyChallengeRewardLabel":          "A Daily Challenge Reward.",
		"DefaultAnnouncementMessage":         "Welcome to Sonic Runners Revival, powered by OUTRUN!",
		"DefaultLoginRouletteMessage":        "Earn some items to help you get a high score, and maybe even a top place in the rankings!",
		"DefaultMaintenanceMessage":          "The server is currently undergoing scheduled maintenance. Check our social media, especially our Discord server, for more information.",
		"DefaultRewardLabel":                 "A gift from the Revival Team.",
		"FirstLoginBonusRewardLabel":         "A Debut Dash Login Bonus.", // TODO: Should this be corrected to "Start Dash Login Bonus" for consistency?
		"LeagueHighRankingRewardLabel":       "A reward for getting the following position in the Runners' League High Score Ranking: %v.",
		"LeaguePromotionRewardLabel":         "Runners' League Promotion Reward. Story Mode.",
		"LeagueTotalRankingRewardLabel":      "A reward for getting the following position in the Runners' League Total Score Ranking: %v.",
		"LoginBonusRewardLabel":              "A Login Bonus.",
		"QuickLeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League Timed Mode High Score Ranking: %v.",
		"QuickLeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Timed Mode.",
		"QuickLeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Timed Mode Total Score Ranking: %v.",
		"TestString":                         "The quick blue hedgehog jumped over the lazy egg 1234567890 times.",
		"WatermarkTicker_1":                  "This server is powered by [ff0000]OUTRUN!",
		"WatermarkTicker_2":                  "ID: [0000ff]%s",
		"WatermarkTicker_3":                  "High score (Timed Mode): [0000ff]%v",
		"WatermarkTicker_4":                  "High score (Story Mode): [0000ff]%v",
		"WatermarkTicker_5":                  "Total distance ran (Story Mode): [0000ff]%v",
	},
}

var LanguageEnumToLanguageCodeTable map[int64]string{
	enums.LangJapanese: "ja",
	enums.LangEnglish: "en",
	enums.LangChineseZH: "zh",
	enums.LangChineseZHJ: "zhj",
	enums.LangKorean: "ko",
	enums.LangFrench: "fr",
	enums.LangGerman: "de",
	enums.LangSpanish: "es",
	enums.LangPortuguese: "pt",
	enums.LangItalian: "it",
	enums.LangRussian: "ru",
}

func GetStringByLanguage(language int64, key string) string {
	return "Not Yet Implemented"
}
