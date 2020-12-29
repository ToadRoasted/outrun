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
		"DefaultMaintenanceMessage": "Sonic Runners Revival is currently in maintenance mode!\nPlease check our social media for more information!\n\n" +
			"Our website: https://sonic.runner.es/\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"DefaultRewardLabel": "A gift from the Revival Team.",
		"DeviceSuspensionNotice": "This device has been blocked from accessing the Sonic Runners Revival game server.\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"FirstLoginBonusRewardLabel":    "A Start Dash Login Bonus.",
		"LeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League High Score Ranking: %v.",
		"LeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Story Mode.",
		"LeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Total Score Ranking: %v.",
		"LoginBonusRewardLabel":         "A Login Bonus.",
		"NewAccountsDisabledNotice": "Registration of new accounts is disabled at the moment. For more information, please visit our Twitter or Discord.\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"NewAccountsDisabledBetaNotice": "Registration of new accounts is not possible on this beta server. You will need to contact someone on the Revival Team so they can create a transfer ID and password for you.\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"QuickLeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League Timed Mode High Score Ranking: %v.",
		"QuickLeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Timed Mode.",
		"QuickLeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Timed Mode Total Score Ranking: %v.",
		"Rank999RewardLabel":                 "A reward for reaching Rank 999. Very few players are dedicated enough to reach this rank!",
		"SuspensionNotice_Temporary": "The account on this device has been temporarily suspended from Sonic Runners Revival for %s\n" +
			"You will be able to play the game again at: %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionNotice_Permanent": "The account on this device has been permanently banned from Sonic Runners Revival for %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionReason_0": "an unspecified reason.",
		"SuspensionReason_1": "cheating in-game.\nPlease don't do that!",
		"SuspensionReason_2": "packet manipulation or exploiting a server glitch.\nPlease stop.",
		"SuspensionReason_3": "attempting to mess up the economy through various methods.",
		"SuspensionReason_4": "an unknown reason (RID 4).",
		"SuspensionReason_5": "an unknown reason (RID 5).",
		"SuspensionReason_6": "an unknown reason (RID 6).",
		"SuspensionReason_7": "an unknown reason (RID 7).",
		"SuspensionReason_8": "an unknown reason (RID 8).",
		"SuspensionReason_9": "an unknown reason (RID 9).",
		"UpdateGameNotice": "Please update Sonic Runners Revival; this version is no longer supported on Sonic Runners Revival!\n\n" +
			"Our website: https://sonic.runner.es/\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"WatermarkTicker_1": "This server is powered by [ff0000]Outrun for Revival!",
		"WatermarkTicker_2": "User ID: [0000ff]%s",
		"WatermarkTicker_3": "High score (Timed Mode): [0000ff]%v",
		"WatermarkTicker_4": "High score (Story Mode): [0000ff]%v",
		"WatermarkTicker_5": "Total distance ran (Story Mode): [0000ff]%v",
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
				result = "ERR: Key \"" + key + "\" does not exist or is empty!"
			}
		} else {
			result = "ERR: Key \"" + key + "\" does not exist in " + LanguageEnumToLanguageNameTable[language] + "!"
		}
	}
	return result
}
