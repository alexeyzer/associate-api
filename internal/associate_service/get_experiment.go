package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *AssociateApiServiceServer) GetExperiment(ctx context.Context, req *desc.GetExperimentRequest) (*desc.GetExperimentResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.experimentService.GetExperiment(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return datastructExperimentToPB(res), nil
}

func datastructExperimentToPB(d *datastruct.Experiment) *desc.GetExperimentResponse {
	response := &desc.GetExperimentResponse{
		Id:                     d.ID,
		Name:                   d.Name,
		ExperimentStimuses:     make([]*desc.ExperimentStimuses, 0, len(d.ExperimentStimuses)),
		ExperimentAssotiations: make([]*desc.ExperimentAssociations, 0, len(d.ExperimentReulsts)),
		ExperimentGrahp:        BuildGrahp(d.ExperimentReulsts),
		Nodes:                  BuildNodes(d.ExperimentReulsts),
	}
	for _, result := range d.ExperimentReulsts {
		response.ExperimentAssotiations = append(response.ExperimentAssotiations, ExperimentResultTopb(result))
	}
	for _, stimus := range d.ExperimentStimuses {
		response.ExperimentStimuses = append(response.ExperimentStimuses, ExperimentStimusTopb(stimus))
	}

	return response
}

func BuildNodes(d []*datastruct.ExperimentResultList) []string {
	results := make([]string, 0, len(d))
	mapp := make(map[string]int64, 0)
	for _, r := range d {
		mapp[r.StimusWord] = 0
		mapp[r.AssotiationWord] = 0
	}
	for key := range mapp {
		results = append(results, key)
	}
	return results
}

func BuildGrahp(d []*datastruct.ExperimentResultList) []*desc.ExperimentGrahp {
	grahp := make(map[string]map[string]int64, 0)

	for _, r := range d {
		_, ok := grahp[r.StimusWord]
		if !ok {
			grahp[r.StimusWord] = make(map[string]int64, 0)
		}
		grahp[r.StimusWord][r.AssotiationWord] = grahp[r.StimusWord][r.AssotiationWord] + 1
	}
	result := make([]*desc.ExperimentGrahp, 0, len(grahp))

	for key, mapp := range grahp {
		for associate, amount := range mapp {
			result = append(result, &desc.ExperimentGrahp{
				StimusWord:      key,
				AssotiationWord: associate,
				Amount:          amount,
			})
		}
	}
	return result
}

func ExperimentResultTopb(d *datastruct.ExperimentResultList) *desc.ExperimentAssociations {
	res := &desc.ExperimentAssociations{
		AssotiationWordId: d.AssotiationWordID,
		AssotiationWord:   d.AssotiationWord,
		StimusWordId:      d.StimusWordID,
		StimusWord:        d.StimusWord,
	}
	return res
}

func ExperimentStimusTopb(d *datastruct.ExperimentStimus) *desc.ExperimentStimuses {
	stimus := &desc.ExperimentStimuses{
		Id:   wrapperspb.Int64(d.StimusID),
		Name: wrapperspb.String(d.Name),
	}
	if d.LimitedResponseTime != nil {
		stimus.LimitedResponseTime = wrapperspb.Int64(*d.LimitedResponseTime)
	}
	return stimus
}
