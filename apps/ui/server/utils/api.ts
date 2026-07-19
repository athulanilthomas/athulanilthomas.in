export function createApiClient(event: Parameters<typeof useRuntimeConfig>[0]) {
  const config = useRuntimeConfig(event)

  return $fetch.create({
    baseURL: config.public.apiBase,
    headers: {
      'X-API-Secret': config.apiSecret
    }
  })
}