package user

import (
	"errors"
	"github.com/alexeyval/test-task-makves-group/internals/app/models"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func NewUser(record []string) (*models.User, error) {
	if len(record) != models.CountField {
		return nil, errors.New("len(record) != CountField")
	}
	id, err := strconv.Atoi(record[models.ID])
	if err != nil {
		return nil, err
	}

	return &models.User{
		UserID:                    record[models.UserID],
		ID:                        int64(id),
		UID:                       record[models.UID],
		Domain:                    record[models.Domain],
		CN:                        record[models.CN],
		Department:                record[models.Department],
		Title:                     record[models.Title],
		Who:                       record[models.Who],
		LogonCount:                record[models.LogonCount],
		NumLogons7:                record[models.NumLogons7],
		NumShare7:                 record[models.NumShare7],
		NumFile7:                  record[models.NumFile7],
		NumAd7:                    record[models.NumAd7],
		NumN7:                     record[models.NumN7],
		NumLogons14:               record[models.NumLogons14],
		NumShare14:                record[models.NumShare14],
		NumFile14:                 record[models.NumFile14],
		NumAd14:                   record[models.NumAd14],
		NumN14:                    record[models.NumN14],
		NumLogons30:               record[models.NumLogons30],
		NumShare30:                record[models.NumShare30],
		NumFile30:                 record[models.NumFile30],
		NumAd30:                   record[models.NumAd30],
		NumN30:                    record[models.NumN30],
		NumLogons150:              record[models.NumLogons150],
		NumShare150:               record[models.NumShare150],
		NumFile150:                record[models.NumFile150],
		NumAd150:                  record[models.NumAd150],
		NumN150:                   record[models.NumN150],
		NumLogons365:              record[models.NumLogons365],
		NumShare365:               record[models.NumShare365],
		NumFile365:                record[models.NumFile365],
		NumAd365:                  record[models.NumAd365],
		NumN365:                   record[models.NumN365],
		HasUserPrincipalName:      record[models.HasUserPrincipalName],
		HasMail:                   record[models.HasMail],
		HasPhone:                  record[models.HasPhone],
		FlagDisabled:              record[models.FlagDisabled],
		FlagLockout:               record[models.FlagLockout],
		FlagPasswordNotRequired:   record[models.FlagPasswordNotRequired],
		FlagPasswordCantChange:    record[models.FlagPasswordCantChange],
		FlagDontExpirePassword:    record[models.FlagDontExpirePassword],
		OwnedFiles:                record[models.OwnedFiles],
		NumMailboxes:              record[models.NumMailboxes],
		NumMemberOfGroups:         record[models.NumMemberOfGroups],
		NumMemberOfIndirectGroups: record[models.NumMemberOfIndirectGroups],
		MemberOfIndirectGroupsIds: toSliceInt(record[models.MemberOfIndirectGroupsIds]),
		MemberOfGroupsIds:         toSliceInt(record[models.MemberOfGroupsIds]),
		IsAdmin:                   record[models.IsAdmin],
		IsService:                 record[models.IsService],
	}, nil
}

func toSliceInt(s string) []int64 {
	if strings.TrimSpace(s) == "" {
		return []int64{}
	}
	split := strings.Split(s, ";")

	ints := make([]int64, 0, len(split))
	for _, v := range split {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Errorln("В файле некорректные данные у "+
				"MemberOfIndirectGroupsIds или MemberOfGroupsIds:", s)
		}
		ints = append(ints, num)
	}

	return ints
}
