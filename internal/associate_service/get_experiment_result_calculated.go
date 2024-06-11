package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	pb "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) GetExperimentCalculated(ctx context.Context, req *pb.GetExperimentCalculatedRequest) (*pb.GetExperimentCalculatedResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(req.GetNames()))
	for _, name := range req.GetNames() {
		if name != "" {
			names = append(names, name)
		}
	}

	res, err := s.experimentResultService.GetCalculatedResulsts(ctx, req.GetId(), req.GetExperimentResultsPagination().GetNumber(), req.GetExperimentResultsPagination().GetLimit(), names)
	if err != nil {
		return nil, err
	}
	resLongestWeighetst, err := s.experimentResultService.GetLongestAndWeighetst(ctx, req.GetId(), names)
	if err != nil {
		return nil, err
	}

	return &pb.GetExperimentCalculatedResponse{
		ExperimentGrahp: BuildGrahpCalcualted(res),
		Nodes:           BuildNodesCalculated(res),
		LongestChains:   BuildChain(resLongestWeighetst.TopLongest),
		WeidghtetChains: BuildChain(resLongestWeighetst.TopWeight),
		FindWords: buildFind(resLongestWeighetst.FindWolrds),
	}, nil
}

func buildFind(FindWolrds map[string][][]string) map[string]*pb.Chains {
	result := make(map[string]*pb.Chains, len(FindWolrds))

	for key, chain := range FindWolrds {
		result[key] = &pb.Chains{
			Chains: BuildChain(chain),
		}
	}

	return result
}

func BuildNodesCalculated(d []*datastruct.ExperimentResultCalculated) []*pb.Node {
	results := make([]*pb.Node, 0, len(d))
	mapp := make(map[string]int64, 0)
	for _, r := range d {
		mapp[r.StimuesWord] = 0
		mapp[r.AssotiationWord] = 0
	}
	x := int64(1000)
	y := int64(1000)
	for key := range mapp {
		if x > 100000 {
			y += 1000
			x = int64(1000)
		}
		results = append(results, &pb.Node{
			Name: key,
			X:    x,
			Y:    y,
		})
		x += 1000
	}
	return results
}

func BuildGrahpCalcualted(d []*datastruct.ExperimentResultCalculated) []*pb.ExperimentGrahp {
	result := make([]*pb.ExperimentGrahp, 0, len(d))

	for _, dd := range d {
		result = append(result, &pb.ExperimentGrahp{
			StimusWord:      dd.StimuesWord,
			AssotiationWord: dd.AssotiationWord,
			Amount:          dd.Amount,
		})
	}
	return result
}

func BuildChain(ddd [][]string) []*pb.Chain {
	result := make([]*pb.Chain, 0, len(ddd))

	for _, d := range ddd {
		chain := &pb.Chain{
			Words: make([]string, 0),
			Total: int64(len(d)),
		}
		for _, dd := range d {
			chain.Words = append(chain.Words, dd)
		}
		result = append(result, chain)
	}
	return result
}
