package associate_service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) ListStimusWords(ctx context.Context, _ *emptypb.Empty) (*desc.ListStimusWordsResponse, error) {
	res, err := s.stimusWordService.ListStimuses(ctx)
	if err != nil {
		return nil, err
	}

	resp := &desc.ListStimusWordsResponse{
		Words: make([]*desc.StimusWord, 0, len(res)),
	}
	for _, word := range res {
		resp.Words = append(resp.Words, &desc.StimusWord{
			Id:   word.ID,
			Name: word.Name,
		})
	}

	return resp, nil
}
