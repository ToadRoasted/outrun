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

// ID 0 = Story Mode league
// ID 2 = Timed Mode league
func LeagueOperatorInformation(id, highRanking, totalRanking, startDate, endDate, newLeague, oldLeague, numLeagueMembers, sendPresent int64) OperatorInformation {
	return NewOperatorInformation(
		id,
		"-1,-1,"+strconv.Itoa(int(highRanking))+","+strconv.Itoa(int(totalRanking))+",-1,"+strconv.Itoa(int(startDate))+","+strconv.Itoa(int(endDate))+","+strconv.Itoa(int(newLeague))+","+strconv.Itoa(int(oldLeague))+","+strconv.Itoa(int(numLeagueMembers))+","+strconv.Itoa(int(sendPresent)),
	)
}

func EventOperatorInformation(eventid string, ranking, numMembers, sendPresent int64) OperatorInformation {
	return NewOperatorInformation(
		1,
		eventid+","+strconv.Itoa(int(ranking))+","+strconv.Itoa(int(numMembers))+","+strconv.Itoa(int(sendPresent)),
	)
}
