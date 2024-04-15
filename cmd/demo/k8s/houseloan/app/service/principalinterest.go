package service

import "github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"

type PrincipalInterest struct {
	ctx context.Context
}

func NewPrincipalInterest(ctx context.Context, principal, interest float64, periods int) *PrincipalInterest {
	return &PrincipalInterest{
		ctx: ctx,
	}
}
