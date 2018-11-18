package services

import (
	"encoding/json"
	"errors"

	"github.com/heart/play-with-go/members"
	"github.com/heart/play-with-go/request"
)

//BNKService
type BNKService struct{}

func NewBNKService() *BNKService {
	return &BNKService{}
}

func (s *BNKService) GetBNKMembers(r request.IRequester) ([]*members.Member, error) {
	body, err := r.Get("https://raw.githubusercontent.com/whs/bnk48json/master/bnk48.json")
	if err != nil {
		return nil, errors.New("Error Request:" + err.Error())
	}

	mems := make([]*members.Member, 0)
	err = json.Unmarshal([]byte(body), &mems)
	if err != nil {
		return nil, errors.New("Error JSON:" + err.Error())
	}

	return mems, nil
}
