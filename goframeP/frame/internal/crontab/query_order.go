package crontab

import "context"

type Cron struct {
	ctx context.Context
}

func NewCron(ctx context.Context) *Cron {
	return &Cron{
		ctx: ctx,
	}
}

func (c *Cron) QueryOrder() {

}
