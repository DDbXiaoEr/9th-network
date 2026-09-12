import { ref } from 'vue'
import { activitiesConfig } from '../config/activities'

export const parseDate = (value) => {
  if (!value) return NaN
  const match = String(value).match(
    /(\d{4})[-/.](\d{1,2})(?:[-/.](\d{1,2}))?(?:[ T](\d{1,2}):(\d{2}))?/
  )
  if (!match) return Date.parse(value)
  return new Date(
    Number(match[1]),
    Number(match[2]) - 1,
    Number(match[3] || 1),
    Number(match[4] || 0),
    Number(match[5] || 0)
  ).getTime()
}

const inferStatus = (date) => {
  const time = parseDate(date)
  if (Number.isNaN(time)) return 'past'
  return time >= Date.now() ? 'upcoming' : 'past'
}

const normalize = (list) =>
  (Array.isArray(list) ? list : []).map((item, index) => ({
    id: item.id || `activity-${index}`,
    title: item.title || item.theme || '未命名活动',
    date: item.date || item.time || '',
    location: item.location || item.place || '',
    status: item.status || inferStatus(item.date || item.time),
    type: item.type || '活动',
    summary: item.summary || item.desc || ''
  }))

const pickList = (data) => {
  if (Array.isArray(data)) return data
  if (Array.isArray(data?.activities)) return data.activities
  if (Array.isArray(data?.data)) return data.data
  if (Array.isArray(data?.list)) return data.list
  return []
}

export function useActivities() {
  const activities = ref([])
  const loading = ref(false)
  const error = ref('')
  const source = ref('config')

  const load = async () => {
    const { api, activities: local } = activitiesConfig
    if (!api) {
      activities.value = normalize(local)
      source.value = 'config'
      return
    }

    loading.value = true
    error.value = ''
    try {
      const response = await fetch(api, { headers: { Accept: 'application/json' } })
      if (!response.ok) throw new Error(`HTTP ${response.status}`)
      const data = await response.json()
      activities.value = normalize(pickList(data))
      source.value = 'api'
    } catch (err) {
      error.value = err?.message || '活动数据加载失败'
      activities.value = normalize(local)
      source.value = 'config'
    } finally {
      loading.value = false
    }
  }

  load()

  return { activities, loading, error, source, reload: load }
}
