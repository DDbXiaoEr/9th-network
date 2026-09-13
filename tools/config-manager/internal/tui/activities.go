package tui

import "the9thnet/config-manager/internal/model"

func activitiesRoot(ctx *Ctx, a *model.Activities) screen {
	return newFormScreen(ctx, "activities.json", []formRow{
		rowText("在线接口 api（留空用本地数据）", &a.API),
		openRow("板块文案", func() string { return a.Section.Title + a.Section.Highlight }, func() screen {
			return activitiesSectionForm(ctx, &a.Section)
		}),
		openRow("活动列表", func() string { return itemCount(len(a.Activities)) }, func() screen {
			return objectList(ctx, "活动列表", "", &a.Activities,
				func(_ int, v *model.Activity) (string, string) { return v.Title, v.Date },
				activityForm, func() model.Activity { return model.Activity{Status: "upcoming", Type: "活动"} })
		}),
	})
}

func activitiesSectionForm(ctx *Ctx, v *model.ActivitiesSection) screen {
	return newFormScreen(ctx, "活动板块文案", []formRow{
		rowText("eyebrow", &v.Eyebrow),
		rowText("标题", &v.Title),
		rowText("高亮", &v.Highlight),
		rowLong("副标题", &v.Subtitle),
		rowText("即将开始文案", &v.UpcomingLabel),
		rowText("往期回顾文案", &v.PastLabel),
		rowText("空状态文案", &v.EmptyText),
	})
}

func activityForm(ctx *Ctx, v *model.Activity) screen {
	return newFormScreen(ctx, "活动", []formRow{
		rowText("id", &v.ID),
		rowText("主题", &v.Title),
		rowText("时间（如 2026-09-20 19:00）", &v.Date),
		rowText("地点", &v.Location),
		rowText("状态 upcoming/past", &v.Status),
		rowText("类型", &v.Type),
		rowLong("简介", &v.Summary),
	})
}
