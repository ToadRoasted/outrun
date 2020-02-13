package netobj

import (
	"github.com/Mtbcooler/outrun/logic/roulette"
)

type ChaoRouletteGroup struct {
	ChaoWheelOptions ChaoWheelOptions `json:"ORN_lastChaoWheelOptions" db:"last_chao_wheel_options"` // actual wheel options for this wheel
	WheelChao        []string         `json:"ORN_wheelChao" db:"wheel_chao"`                         // what Chao/characters are in this wheel
	ChaoRouletteInfo RouletteInfo     `json:"ORN_chaoRouletteInfo" db:"chao_roulette_info"`          // may not be needed
}

type SqlChaoRouletteGroup struct {
	ChaoWheelOptions []byte `json:"ORN_lastChaoWheelOptions" db:"last_chao_wheel_options"` // actual wheel options for this wheel
	WheelChao        []byte `json:"ORN_wheelChao" db:"wheel_chao"`                         // what Chao/characters are in this wheel
	ChaoRouletteInfo []byte `json:"ORN_chaoRouletteInfo" db:"chao_roulette_info"`          // may not be needed
}

func DefaultChaoRouletteGroup(playerState PlayerState, allowedCharacters, allowedChao []string) ChaoRouletteGroup {
	chaoWheelOptions := DefaultChaoWheelOptions(playerState)
	//wheelChao, err := roulette.GetRandomChaoRouletteItems(chaoWheelOptions.Rarity, exclusions) // populate based on given rarities
	wheelChao, newRarity, err := roulette.GetRandomChaoRouletteItems(chaoWheelOptions.Rarity, allowedCharacters, allowedChao) // populate based on given rarities
	if err != nil {
		panic(err) // TODO: Find a better way to handle error. Hard to manage since the player creators don't already output errors
	}
	// newRarity is rarity but with any modifications that need to be made
	chaoWheelOptions.Rarity = newRarity
	chaoRouletteInfo := DefaultRouletteInfo()
	return ChaoRouletteGroup{
		chaoWheelOptions,
		wheelChao,
		chaoRouletteInfo,
	}
}
