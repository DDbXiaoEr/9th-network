package tui

import (
	"fmt"
	"strings"

	"the9thnet/config-manager/internal/model"
)

func articlesRoot(ctx *Ctx, a *model.Articles) screen {
	return newFormScreen(ctx, "articles.json", []formRow{
		openRow("资源分组", func() string { return itemCount(len(a.Groups)) }, func() screen {
			return objectList(ctx, "资源分组", "", &a.Groups,
				func(_ int, v *model.ResourceGroup) (string, string) { return v.Title, v.Accent },
				resourceGroupForm, func() model.ResourceGroup { return model.ResourceGroup{} })
		}),
		openRow("资源文章", func() string { return itemCount(len(a.Articles)) }, func() screen {
			return objectList(ctx, "资源文章", "", &a.Articles,
				func(_ int, v *model.ResourceArticle) (string, string) { return v.Title, v.Category },
				resourceArticleForm, func() model.ResourceArticle { return model.ResourceArticle{} })
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
		rowText("文件", &v.File),
		openRow("标签", func() string { return itemCount(len(v.Tags)) }, func() screen {
			return newStringList(ctx, "标签", &v.Tags)
		}),
		openRow("正文", func() string { return markdownSummary(v.Markdown) }, func() screen {
			if strings.TrimSpace(v.File) == "" && strings.TrimSpace(v.ID) != "" {
				v.File = "/articles/" + v.ID + ".md"
			}
			return newTextScreen(ctx, "正文 Markdown", &v.Markdown)
		}),
	})
}

func markdownSummary(s string) string {
	if strings.TrimSpace(s) == "" {
		return "（空）"
	}
	return fmt.Sprintf("%d 行", strings.Count(s, "\n")+1)
}
