export default defineEventHandler(async (event) => {
  const api = createApiClient(event)
  const result = await api('/spotify/now-playing', {
    ignoreResponseError: true
  })
  return result
})
