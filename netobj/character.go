package netobj

import (
	"github.com/Mtbcooler/outrun/config/gameconf"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
)

/*
Notes:
  - I believe stars are used as "prestige" for the characters, if all skills are maxed out
    - starMax may be the max prestige
*/

type Character struct { // Can also be used as PlayCharacter
	obj.Character
	Status            int64          `json:"status"` // value from enums.CharacterStatus*
	Level             int64          `json:"level"`
	Exp               int64          `json:"exp"`
	Star              int64          `json:"star"`
	StarMax           int64          `json:"starMax"`
	LockCondition     int64          `json:"lockCondition"` // value from enums.LockCondition*
	CampaignList      []obj.Campaign `json:"campaignList"`
	AbilityLevel      []int64        `json:"abilityLevel"`      // levels for each ability
	AbilityNumRings   []int64        `json:"abilityNumRings"`   // possibly unused?
	AbilityLevelUp    []int64        `json:"abilityLevelup"`    // which ability ID(s) leveled up during a run
	AbilityLevelUpExp []int64        `json:"abilityLevelupExp"` // exp to level up corresponding abilityLevelup ability?
}

var tick = 0

func DefaultCharacter(char obj.Character) Character {
	status := int64(enums.CharacterStatusUnlocked)
	level := int64(0)
	exp := int64(0)
	star := int64(0)     // Limit breaks
	starMax := int64(10) // Max number of limit breaks?
	lockCondition := int64(enums.LockConditionOpen)
	campaignList := []obj.Campaign{}
	abilityLevel := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // index 0 is a dummy entry not used by the game apparently???
	abilityNumRings := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	abilityLevelUp := []int64{}
	abilityLevelUpExp := []int64{}
	return Character{
		char,
		status,
		level,
		exp,
		star,
		starMax,
		lockCondition,
		campaignList,
		abilityLevel,
		abilityNumRings,
		abilityLevelUp,
		abilityLevelUpExp,
	}
}

func DefaultLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionRingOrRedRing)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

func DefaultStageLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionMileageEpisode)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

func DefaultRouletteLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionRoulette)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

// UnlockedCharacterState is a default CharacterState
func UnlockedCharacterState() []Character { // every character
	// TODO: It looks like the game only wants 300000-300020, otherwise an index error is created. Investigate this in game!
	return []Character{
		DefaultCharacter(constobjs.CharacterSonic),
		DefaultCharacter(constobjs.CharacterTails),
		DefaultCharacter(constobjs.CharacterKnuckles),
		DefaultCharacter(constobjs.CharacterAmy),
		DefaultCharacter(constobjs.CharacterShadow),
		DefaultCharacter(constobjs.CharacterBlaze),
		DefaultCharacter(constobjs.CharacterRouge),
		DefaultCharacter(constobjs.CharacterOmega),
		DefaultCharacter(constobjs.CharacterBig),
		DefaultCharacter(constobjs.CharacterCream),
		DefaultCharacter(constobjs.CharacterEspio),
		DefaultCharacter(constobjs.CharacterCharmy),
		DefaultCharacter(constobjs.CharacterVector),
		DefaultCharacter(constobjs.CharacterSilver),
		DefaultCharacter(constobjs.CharacterMetalSonic),
		DefaultCharacter(constobjs.CharacterClassicSonic),
		DefaultCharacter(constobjs.CharacterWerehog),
		DefaultCharacter(constobjs.CharacterSticks),
		DefaultCharacter(constobjs.CharacterTikal),
		DefaultCharacter(constobjs.CharacterMephiles),
		DefaultCharacter(constobjs.CharacterPSISilver),
	}
}

func DefaultCharacterState() []Character {
	if gameconf.CFile.AllCharactersUnlocked {
		return UnlockedCharacterState()
	}
	return []Character{
		DefaultCharacter(constobjs.CharacterSonic),
		DefaultStageLockedCharacter(constobjs.CharacterTails),    // Episode 11
		DefaultStageLockedCharacter(constobjs.CharacterKnuckles), // Episode 17
		DefaultRouletteLockedCharacter(constobjs.CharacterAmy),
		DefaultRouletteLockedCharacter(constobjs.CharacterShadow),
		DefaultLockedCharacter(constobjs.CharacterBlaze),
		DefaultRouletteLockedCharacter(constobjs.CharacterRouge),
		DefaultRouletteLockedCharacter(constobjs.CharacterOmega),
		DefaultRouletteLockedCharacter(constobjs.CharacterBig),
		DefaultRouletteLockedCharacter(constobjs.CharacterCream),
		DefaultRouletteLockedCharacter(constobjs.CharacterEspio),
		DefaultRouletteLockedCharacter(constobjs.CharacterCharmy),
		DefaultRouletteLockedCharacter(constobjs.CharacterVector),
		DefaultRouletteLockedCharacter(constobjs.CharacterSilver),
		DefaultLockedCharacter(constobjs.CharacterMetalSonic),
		DefaultLockedCharacter(constobjs.CharacterClassicSonic),
		DefaultLockedCharacter(constobjs.CharacterWerehog),
		DefaultLockedCharacter(constobjs.CharacterSticks),
		DefaultLockedCharacter(constobjs.CharacterTikal),
		DefaultLockedCharacter(constobjs.CharacterMephiles),
		DefaultLockedCharacter(constobjs.CharacterPSISilver),
		// other characters will be added to the CharacterState as they are obtained on the roulette
	}
}

// CheckForLevelAbnormalities fixes any abnormalities in the ability levels and the character level
// Returns: Corrected Character, and if abnormalities were detected
func CheckForLevelAbnormalities(char Character) (Character, bool) {
	abnormalitiesDetected := false
	if char.Star > char.StarMax {
		char.Star = char.StarMax
		abnormalitiesDetected = true
	}
	realLevel := int64(0)
	for index, aL := range char.AbilityLevel {
		if (index == 0 && aL != 0) || aL < 0 { // does ability index 0 have a non-zero level, or is ability level negative?
			char.AbilityLevel[index] = 0
			aL = 0
			abnormalitiesDetected = true
		}
		if aL > 10 { // is ability level above 10?
			char.AbilityLevel[index] = 10
			aL = 10
			abnormalitiesDetected = true
		}
		realLevel += aL
	}
	if realLevel != char.Level {
		char.Level = realLevel
		abnormalitiesDetected = true
	}
	return char, abnormalitiesDetected
}
