package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Breadcrumb(year *cal.Year) string {
	return header.Items{
		header.NewIntItem(year.Number),
		header.NewTextItem("Goals").Ref(true),
	}.Table(true)
}

func Goals(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	return page.Modules{{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Year":       year,
			"Breadcrumb": Breadcrumb(year),
			"Extra":      header.Items{}.WithTopRightCorner(cfg.ClearTopRightCorner),
		},
	}}, nil
}
