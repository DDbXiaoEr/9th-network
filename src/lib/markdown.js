import { marked } from 'marked'

const renderer = {
  html() {
    return ''
  },
  link({ href, title, text }) {
    const safe = sanitizeHref(href)
    if (!safe) return text
    const t = title ? ` title="${escapeAttr(title)}"` : ''
    return `<a href="${escapeAttr(safe)}"${t} target="_blank" rel="noopener noreferrer">${text}</a>`
  },
  image({ href, title, text }) {
    const safe = sanitizeHref(href)
    if (!safe) return escapeAttr(text)
    const t = title ? ` title="${escapeAttr(title)}"` : ''
    return `<img src="${escapeAttr(safe)}" alt="${escapeAttr(text)}"${t}>`
  }
}

marked.use({ renderer, gfm: true })

export function renderMarkdown(source) {
  return marked.parse(source || '', { async: false })
}

function sanitizeHref(href) {
  if (!href) return ''
  const trimmed = String(href).trim()
  if (trimmed.startsWith('#') || (trimmed.startsWith('/') && !trimmed.startsWith('//'))) {
    return trimmed
  }
  try {
    const url = new URL(trimmed, 'http://localhost')
    if (url.protocol === 'http:' || url.protocol === 'https:' || url.protocol === 'mailto:') {
      return trimmed
    }
  } catch {
    return ''
  }
  return ''
}

function escapeAttr(value) {
  return String(value)
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
}
