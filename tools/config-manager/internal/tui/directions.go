package tui

import "the9thnet/config-manager/internal/model"

func directionsRoot(ctx *Ctx, d *model.Directions) screen {
	return newFormScreen(ctx, "directions.json", []formRow{
		openRow("板块文案", func() string { return d.Section.Title + d.Section.Highlight }, func() screen {
			return directionsSectionForm(ctx, &d.Section)
		}),
		openRow("兴趣方向", func() string { return itemCount(len(d.Directions)) }, func() screen {
			return objectList(ctx, "兴趣方向", "", &d.Directions,
				func(_ int, v *model.Direction) (string, string) { return v.Name, v.ID },
				directionForm, func() model.Direction { return model.Direction{} })
		}),
		openRow("CTF 战队", func() string { return d.CTF.Title }, func() screen {
			return ctfForm(ctx, &d.CTF)
		}),
	})
}

func directionsSectionForm(ctx *Ctx, v *model.DirectionsSection) screen {
	return newFormScreen(ctx, "方向板块文案", []formRow{
		rowText("eyebrow", &v.Eyebrow),
		rowText("标题", &v.Title),
		rowText("高亮", &v.Highlight),
		rowLong("副标题", &v.Subtitle),
	})
}

func directionForm(ctx *Ctx, v *model.Direction) screen {
	return newFormScreen(ctx, "兴趣方向", []formRow{
		rowText("id", &v.ID),
		rowText("图标", &v.Icon),
		rowText("名称", &v.Name),
		rowText("标语", &v.Tagline),
		rowLong("描述", &v.Desc),
		openRow("亮点", func() string { return itemCount(len(v.Highlights)) }, func() screen {
			return newStringList(ctx, "亮点", &v.Highlights)
		}),
		openRow("技术栈", func() string { return itemCount(len(v.Stack)) }, func() screen {
			return newStringList(ctx, "技术栈", &v.Stack)
		}),
	})
}

func ctfForm(ctx *Ctx, v *model.CTF) screen {
	return newFormScreen(ctx, "CTF 战队", []formRow{
		boolRow("启用", &v.Enabled),
		rowText("eyebrow", &v.Eyebrow),
		rowText("标题", &v.Title),
		rowLong("副标题", &v.Subtitle),
		openRow("历史最好成绩", func() string { return v.BestRank.Value + v.BestRank.Unit + " " + v.BestRank.Title }, func() screen {
			return bestRankForm(ctx, &v.BestRank)
		}),
		openRow("战队信息", func() string { return v.Team.Name }, func() screen {
			return teamForm(ctx, &v.Team)
		}),
		openRow("队员", func() string { return itemCount(len(v.Members)) }, func() screen {
			return objectList(ctx, "队员", "", &v.Members,
				func(_ int, m *model.Member) (string, string) { return m.Name, m.Role },
				memberForm, func() model.Member { return model.Member{} })
		}),
		openRow("比赛记录", func() string { return itemCount(len(v.Competitions)) }, func() screen {
			return objectList(ctx, "比赛记录", "", &v.Competitions,
				func(_ int, c *model.Competition) (string, string) { return c.Name, c.Result },
				competitionForm, func() model.Competition { return model.Competition{Type: "国际赛"} })
		}),
	})
}

func bestRankForm(ctx *Ctx, v *model.BestRank) screen {
	return newFormScreen(ctx, "历史最好成绩", []formRow{
		rowText("标签", &v.Label),
		rowText("数值", &v.Value),
		rowText("单位", &v.Unit),
		rowText("标题", &v.Title),
		rowLong("备注", &v.Note),
	})
}

func teamForm(ctx *Ctx, v *model.Team) screen {
	return newFormScreen(ctx, "战队信息", []formRow{
		rowText("名称", &v.Name),
		rowText("全称", &v.FullName),
		rowText("口号", &v.Slogan),
		rowText("主页 URL", &v.URL),
		openRow("信息字段", func() string { return itemCount(len(v.Fields)) }, func() screen {
			return objectList(ctx, "信息字段", "", &v.Fields,
				func(_ int, f *model.TeamField) (string, string) { return f.Label, f.Value },
				teamFieldForm, func() model.TeamField { return model.TeamField{} })
		}),
	})
}

func teamFieldForm(ctx *Ctx, v *model.TeamField) screen {
	return newFormScreen(ctx, "信息字段", []formRow{
		rowText("字段名", &v.Label),
		rowText("值", &v.Value),
	})
}

func memberForm(ctx *Ctx, v *model.Member) screen {
	return newFormScreen(ctx, "队员", []formRow{
		rowText("姓名", &v.Name),
		rowText("ID", &v.Handle),
		rowText("角色", &v.Role),
		rowText("方向", &v.Direction),
	})
}

func competitionForm(ctx *Ctx, v *model.Competition) screen {
	return newFormScreen(ctx, "比赛记录", []formRow{
		rowText("年份", &v.Date),
		rowText("赛事名称", &v.Name),
		rowText("类型", &v.Type),
		rowText("成绩", &v.Result),
		rowLong("备注", &v.Note),
	})
}
