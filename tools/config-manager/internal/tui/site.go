package tui

import (
	"fmt"

	"the9thnet/config-manager/internal/model"
)

func siteRoot(ctx *Ctx, s *model.Site) screen {
	return newFormScreen(ctx, "site.json", []formRow{
		openRow("品牌信息", func() string { return s.Brand.Name }, func() screen {
			return brandForm(ctx, &s.Brand)
		}),
		openRow("导航链接", func() string { return itemCount(len(s.NavLinks)) }, func() screen {
			return objectList(ctx, "导航链接", "", &s.NavLinks,
				func(_ int, v *model.NavLink) (string, string) { return v.Label, v.Kind },
				navLinkForm, func() model.NavLink { return model.NavLink{Kind: "section"} })
		}),
		openRow("页脚分组", func() string { return itemCount(len(s.FooterLinks)) }, func() screen {
			return objectList(ctx, "页脚分组", "", &s.FooterLinks,
				func(_ int, v *model.FooterGroup) (string, string) { return v.Title, itemCount(len(v.Items)) },
				footerGroupForm, func() model.FooterGroup { return model.FooterGroup{} })
		}),
		openRow("首屏轮播", func() string { return itemCount(len(s.HeroSlides)) }, func() screen {
			return objectList(ctx, "首屏轮播", "", &s.HeroSlides,
				func(_ int, v *model.HeroSlide) (string, string) { return v.Title, v.Tag },
				heroSlideForm, func() model.HeroSlide { return model.HeroSlide{} })
		}),
		openRow("统计数据", func() string { return itemCount(len(s.Stats)) }, func() screen {
			return objectList(ctx, "统计数据", "", &s.Stats,
				func(_ int, v *model.Stat) (string, string) { return v.Label, v.Value + v.Suffix },
				statForm, func() model.Stat { return model.Stat{} })
		}),
		openRow("社团介绍", func() string { return itemCount(len(s.AboutTabs)) }, func() screen {
			return objectList(ctx, "社团介绍", "", &s.AboutTabs,
				func(_ int, v *model.AboutTab) (string, string) { return v.Label, v.Key },
				aboutTabForm, func() model.AboutTab { return model.AboutTab{} })
		}),
		openRow("资源分组", func() string { return itemCount(len(s.ResourceGroups)) }, func() screen {
			return objectList(ctx, "资源分组", "", &s.ResourceGroups,
				func(_ int, v *model.ResourceGroup) (string, string) { return v.Title, v.Accent },
				resourceGroupForm, func() model.ResourceGroup { return model.ResourceGroup{} })
		}),
		openRow("资源文章", func() string { return itemCount(len(s.ResourceArticles)) }, func() screen {
			return objectList(ctx, "资源文章", "", &s.ResourceArticles,
				func(_ int, v *model.ResourceArticle) (string, string) { return v.Title, v.Category },
				resourceArticleForm, func() model.ResourceArticle { return model.ResourceArticle{} })
		}),
		openRow("加入信息", func() string { return s.JoinInfo.Title + s.JoinInfo.Highlight }, func() screen {
			return joinInfoForm(ctx, &s.JoinInfo)
		}),
	})
}

func brandForm(ctx *Ctx, v *model.Brand) screen {
	return newFormScreen(ctx, "品牌信息", []formRow{
		rowText("名称", &v.Name),
		rowText("英文名", &v.EnName),
		rowText("英文全称", &v.FullEnName),
		rowText("口号", &v.Slogan),
		rowText("学校", &v.School),
		rowText("Logo 路径", &v.Logo),
	})
}

func navLinkForm(ctx *Ctx, v *model.NavLink) screen {
	return newFormScreen(ctx, "导航链接", []formRow{
		rowText("id", &v.ID),
		rowText("名称", &v.Label),
		rowText("类型 section/route", &v.Kind),
		rowText("路径 to", &v.To),
	})
}

func footerGroupForm(ctx *Ctx, v *model.FooterGroup) screen {
	return newFormScreen(ctx, "页脚分组", []formRow{
		rowText("标题", &v.Title),
		openRow("链接", func() string { return itemCount(len(v.Items)) }, func() screen {
			return objectList(ctx, "页脚链接", "", &v.Items,
				func(_ int, it *model.FooterItem) (string, string) { return it.Label, it.To },
				footerItemForm, func() model.FooterItem { return model.FooterItem{} })
		}),
	})
}

func footerItemForm(ctx *Ctx, v *model.FooterItem) screen {
	return newFormScreen(ctx, "页脚链接", []formRow{
		rowText("名称", &v.Label),
		rowText("目标 to", &v.To),
	})
}

func heroSlideForm(ctx *Ctx, v *model.HeroSlide) screen {
	return newFormScreen(ctx, "首屏轮播", []formRow{
		rowText("图片", &v.Image),
		rowText("标签", &v.Tag),
		rowText("标题", &v.Title),
		rowText("副标题", &v.Subtitle),
		rowLong("描述", &v.Desc),
	})
}

func statForm(ctx *Ctx, v *model.Stat) screen {
	return newFormScreen(ctx, "统计数据", []formRow{
		rowText("key", &v.Key),
		rowText("数值", &v.Value),
		rowText("后缀", &v.Suffix),
		rowText("标签", &v.Label),
	})
}

func aboutTabForm(ctx *Ctx, v *model.AboutTab) screen {
	return newFormScreen(ctx, "社团介绍", []formRow{
		rowText("key", &v.Key),
		rowText("名称", &v.Label),
		openRow("配图", func() string { return itemCount(len(v.Images)) }, func() screen {
			return newStringList(ctx, "配图", &v.Images)
		}),
		openRow("段落（简介）", func() string { return itemCount(len(v.Paragraphs)) }, func() screen {
			return newStringList(ctx, "段落", &v.Paragraphs)
		}),
		openRow("成员分组", func() string { return itemCount(len(v.Groups)) }, func() screen {
			return objectList(ctx, "成员分组", "", &v.Groups,
				func(_ int, g *model.MemberGroup) (string, string) { return g.Role, itemCount(len(g.Names)) },
				memberGroupForm, func() model.MemberGroup { return model.MemberGroup{} })
		}),
		openRow("条目（发展）", func() string { return itemCount(len(v.Items)) }, func() screen {
			return newStringList(ctx, "条目", &v.Items)
		}),
	})
}

func memberGroupForm(ctx *Ctx, v *model.MemberGroup) screen {
	return newFormScreen(ctx, "成员分组", []formRow{
		rowText("角色", &v.Role),
		openRow("成员", func() string { return itemCount(len(v.Names)) }, func() screen {
			return newStringList(ctx, "成员姓名", &v.Names)
		}),
	})
}

func resourceGroupForm(ctx *Ctx, v *model.ResourceGroup) screen {
	return newFormScreen(ctx, "资源分组", []formRow{
		rowText("标题", &v.Title),
		rowText("图标", &v.Icon),
		rowText("主题色 accent", &v.Accent),
		openRow("链接", func() string { return itemCount(len(v.Links)) }, func() screen {
			return newStringList(ctx, "链接", &v.Links)
		}),
	})
}

func resourceArticleForm(ctx *Ctx, v *model.ResourceArticle) screen {
	return newFormScreen(ctx, "资源文章", []formRow{
		rowText("id", &v.ID),
		rowText("分类", &v.Category),
		rowText("标题", &v.Title),
		rowText("作者", &v.Author),
		rowLong("摘要", &v.Summary),
		rowText("封面", &v.Cover),
		openRow("标签", func() string { return itemCount(len(v.Tags)) }, func() screen {
			return newStringList(ctx, "标签", &v.Tags)
		}),
		openRow("正文", func() string { return itemCount(len(v.Content)) }, func() screen {
			return newStringList(ctx, "正文段落", &v.Content)
		}),
	})
}

func joinInfoForm(ctx *Ctx, v *model.JoinInfo) screen {
	return newFormScreen(ctx, "加入信息", []formRow{
		rowText("eyebrow", &v.Eyebrow),
		rowText("标题", &v.Title),
		rowText("高亮", &v.Highlight),
		rowText("副标题", &v.Subtitle),
		openRow("须知", func() string { return v.Notice.Title }, func() screen {
			return noticeForm(ctx, &v.Notice)
		}),
		openRow("流程步骤", func() string { return itemCount(len(v.Steps)) }, func() screen {
			return objectList(ctx, "流程步骤", "", &v.Steps,
				func(_ int, st *model.Step) (string, string) { return st.Title, st.Desc },
				stepForm, func() model.Step { return model.Step{} })
		}),
		rowLong("备注", &v.Note),
		openRow("按钮 CTA", func() string { return fmt.Sprintf("%s → %s", v.CTA.Label, v.CTA.To) }, func() screen {
			return ctaForm(ctx, &v.CTA)
		}),
	})
}

func noticeForm(ctx *Ctx, v *model.Notice) screen {
	return newFormScreen(ctx, "加入须知", []formRow{
		rowText("标题", &v.Title),
		rowLong("描述", &v.Desc),
	})
}

func stepForm(ctx *Ctx, v *model.Step) screen {
	return newFormScreen(ctx, "流程步骤", []formRow{
		rowText("标题", &v.Title),
		rowLong("描述", &v.Desc),
	})
}

func ctaForm(ctx *Ctx, v *model.CTA) screen {
	return newFormScreen(ctx, "按钮 CTA", []formRow{
		rowText("文案", &v.Label),
		rowText("目标 to", &v.To),
	})
}
