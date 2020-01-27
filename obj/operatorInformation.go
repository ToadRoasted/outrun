package obj

import "strconv"

// TODO: create a param object that makes it easier to create informations

type OperatorInformation struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func NewOperatorInformation(id int64, content string) OperatorInformation {
	return OperatorInformation{
		id,
		content,
	}
}

func LeagueOperatorInformation(highRanking, totalRanking, startDate, endDate, newLeague, oldLeague, numLeagueMembers, sendPresent int64) OperatorInformation {
	return NewOperatorInformation(
		0,
		"-1,-1,"+strconv.Itoa(int(highRanking))+","+strconv.Itoa(int(totalRanking))+",-1,"+strconv.Itoa(int(startDate))+","+strconv.Itoa(int(endDate))+","+strconv.Itoa(int(newLeague))+","+strconv.Itoa(int(oldLeague))+","+strconv.Itoa(int(numLeagueMembers))+","+strconv.Itoa(int(sendPresent)),
	)
}
