package tui

import "the9thnet/config-manager/internal/model"

func servicesRoot(ctx *Ctx, s *model.Services) screen {
	return newFormScreen(ctx, "services.json", []formRow{
		openRow("板块文案", func() string { return s.Section.Title + s.Section.Highlight }, func() screen {
			return servicesSectionForm(ctx, &s.Section)
		}),
		openRow("分类", func() string { return itemCount(len(s.Categories)) }, func() screen {
			return newStringList(ctx, "服务分类", &s.Categories)
		}),
		openRow("服务列表", func() string { return itemCount(len(s.Services)) }, func() screen {
			return objectList(ctx, "服务列表", "", &s.Services,
				func(_ int, v *model.Service) (string, string) { return v.Title, v.Category },
				serviceForm, func() model.Service { return model.Service{} })
		}),
	})
}

func servicesSectionForm(ctx *Ctx, v *model.ServicesSection) screen {
	return newFormScreen(ctx, "服务板块文案", []formRow{
		rowText("eyebrow", &v.Eyebrow),
		rowText("标题", &v.Title),
		rowText("高亮", &v.Highlight),
		rowLong("副标题", &v.Subtitle),
		rowText("全部分类名", &v.AllLabel),
		rowText("提示文案", &v.Hint),
		rowText("提示操作", &v.HintAction),
	})
}

func serviceForm(ctx *Ctx, v *model.Service) screen {
	return newFormScreen(ctx, "服务", []formRow{
		rowText("id", &v.ID),
		rowText("图标", &v.Icon),
		rowText("名称", &v.Title),
		rowText("分类", &v.Category),
		rowLong("简介", &v.Desc),
		openRow("标签", func() string { return itemCount(len(v.Tags)) }, func() screen {
			return newStringList(ctx, "标签", &v.Tags)
		}),
		rowText("状态", &v.Status),
		boolRow("热门", &v.Hot),
		openRow("访问方式", func() string { return accessSummary(v.Access) }, func() screen {
			if v.Access == nil {
				v.Access = &model.ServiceAccess{Type: "modal"}
			}
			return serviceAccessForm(ctx, v.Access)
		}),
	})
}

func accessSummary(a *model.ServiceAccess) string {
	if a == nil || a.Type == "" {
		return "弹窗提示"
	}
	if a.Type == "link" {
		if a.URL == "" {
			return "链接跳转（未填链接）"
		}
		return "链接跳转 " + a.URL
	}
	return "弹窗提示"
}

func serviceAccessForm(ctx *Ctx, v *model.ServiceAccess) screen {
	return newFormScreen(ctx, "访问方式", []formRow{
		choiceRow("方式", &v.Type, opt("modal", "弹窗提示"), opt("link", "链接跳转")),
		rowText("按钮文案", &v.Label),
		rowText("链接地址", &v.URL),
		rowText("弹窗标题", &v.Title),
		rowLong("弹窗内容", &v.Note),
	})
}
