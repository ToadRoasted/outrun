package localizations

import (
	"github.com/Mtbcooler/outrun/enums"
)

var LocalizedStrings = map[string]map[string]string{
	"en": map[string]string{
		"DailyBattleWinRewardLabel":       "A reward for winning a daily battle.",
		"DailyBattleWinStreakRewardLabel": "A reward for winning %v daily battles in a row.",
		"DailyChallengeRewardLabel":       "A Daily Challenge Reward.",
		"DefaultAnnouncementMessage":      "Welcome to Sonic Runners Revival!",
		"DefaultLoginRouletteMessage":     "Earn some items to help you get a high score, and maybe even a top place in the rankings!",
		"DefaultMaintenanceMessage":       "Sonic Runners Revival is currently in maintenance mode!\nPlease check our social media for more information!",
		"DefaultRewardLabel":              "A gift from the Revival Team.",
		"DeviceSuspensionNotice": "This device has been blocked from accessing the Sonic Runners Revival game server, likely due to repeated incidents of cheating.\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/revivalsonic\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"FirstLoginBonusRewardLabel":         "A Debut Dash Login Bonus.", // TODO: Should this be corrected to "Start Dash Login Bonus" for consistency?
		"LeagueHighRankingRewardLabel":       "A reward for getting the following position in the Runners' League High Score Ranking: %v.",
		"LeaguePromotionRewardLabel":         "Runners' League Promotion Reward. Story Mode.",
		"LeagueTotalRankingRewardLabel":      "A reward for getting the following position in the Runners' League Total Score Ranking: %v.",
		"LoginBonusRewardLabel":              "A Login Bonus.",
		"QuickLeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League Timed Mode High Score Ranking: %v.",
		"QuickLeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Timed Mode.",
		"QuickLeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Timed Mode Total Score Ranking: %v.",
		"SuspensionNotice_Temporary": "The account on this device has been temporarily suspended from Sonic Runners Revival until the following date: %s\n" +
			"Reason: %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/revivalsonic\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionNotice_Permanent": "The account on this device has been permanently banned from Sonic Runners Revival.\n" +
			"Reason: %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/revivalsonic\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionReason_0": "(no reason specified)",
		"SuspensionReason_1": "Cheating",
		"SuspensionReason_2": "Packet manipulation",
		"SuspensionReason_3": "Account selling",
		"SuspensionReason_4": "Reserved for future use (Reason ID 4)",
		"SuspensionReason_5": "Reserved for future use (Reason ID 5)",
		"SuspensionReason_6": "Reserved for future use (Reason ID 6)",
		"SuspensionReason_7": "Reserved for future use (Reason ID 7)",
		"SuspensionReason_8": "Reserved for future use (Reason ID 8)",
		"SuspensionReason_9": "Reserved for future use (Reason ID 9)",
		"WatermarkTicker_1":  "This server is powered by [ff0000]Outrun for Revival!",
		"WatermarkTicker_2":  "ID: [0000ff]%s",
		"WatermarkTicker_3":  "High score (Timed Mode): [0000ff]%v",
		"WatermarkTicker_4":  "High score (Story Mode): [0000ff]%v",
		"WatermarkTicker_5":  "Total distance ran (Story Mode): [0000ff]%v",
	},
}

var LanguageEnumToLanguageCodeTable = map[int64]string{
	enums.LangJapanese:   "ja",
	enums.LangEnglish:    "en",
	enums.LangChineseZH:  "zh",
	enums.LangChineseZHJ: "zhj",
	enums.LangKorean:     "ko",
	enums.LangFrench:     "fr",
	enums.LangGerman:     "de",
	enums.LangSpanish:    "es",
	enums.LangPortuguese: "pt",
	enums.LangItalian:    "it",
	enums.LangRussian:    "ru",
}

var LanguageEnumToLanguageNameTable = map[int64]string{
	enums.LangJapanese:   "Japanese",
	enums.LangEnglish:    "English",
	enums.LangChineseZH:  "Chinese",
	enums.LangChineseZHJ: "Chinese",
	enums.LangKorean:     "Korean",
	enums.LangFrench:     "French",
	enums.LangGerman:     "German",
	enums.LangSpanish:    "Spanish",
	enums.LangPortuguese: "Portuguese",
	enums.LangItalian:    "Italian",
	enums.LangRussian:    "Russian",
}

func GetStringByLanguage(language int64, key string, fallBackToEnglish bool) string {
	result := LocalizedStrings[LanguageEnumToLanguageCodeTable[language]][key]
	if result == "" {
		if fallBackToEnglish {
			result = LocalizedStrings["en"][key]
			if result == "" {
				result = "ERROR: Key \"" + key + "\" does not exist or is empty!"
			}
		} else {
			result = "ERROR: Key \"" + key + "\" does not exist for the " + LanguageEnumToLanguageNameTable[language] + " language!"
		}
	}
	return result
}
