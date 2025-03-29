package compose

import (
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Breadcrumb(year *cal.Year, i int) string {
	return header.Items{
		header.NewIntItem(year.Number),
		header.NewTextItem("Goals").RefText("Goals"),
		header.NewTextItem("Goal " + strconv.Itoa(i+1)).Ref(true),
	}.Table(true)
}

func BreadcrumbIndex(year *cal.Year) string {
	return header.Items{
		header.NewIntItem(year.Number),
		header.NewTextItem("Goals").Ref(true),
	}.Table(true)
}

func VisionBreadcrumb(year *cal.Year, label string) string {
	return header.Items{
		header.NewIntItem(year.Number),
		header.NewTextItem("Goals").RefText("Goals"),
		header.NewTextItem(label).Ref(true),
	}.Table(true)
}

func Goals(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	modules = append(modules, page.Module{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Year":       year,
			"Breadcrumb": BreadcrumbIndex(year),
			"Extra":      header.Items{}.WithTopRightCorner(cfg.ClearTopRightCorner),
		},
	})

	for i := 0; i < 3; i++ {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[1],
			Body: map[string]interface{}{
				"Year":       year,
				"Breadcrumb": Breadcrumb(year, i),
				"Extra":      header.Items{}.WithTopRightCorner(cfg.ClearTopRightCorner),
			},
		})
	}

	visionpages := []string{"10 Year Vision", "5 Year Vision", "1 Year Vision"}

	for idx := range visionpages {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[2],
			Body: map[string]interface{}{
				"Year":       year,
				"Breadcrumb": VisionBreadcrumb(year, visionpages[idx]),
				"Extra":      header.Items{}.WithTopRightCorner(cfg.ClearTopRightCorner),
			},
		})
	}

	return modules, nil
}
