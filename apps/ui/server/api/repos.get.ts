export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig(event)
    const result = await $fetch(`/github/repos`, { baseURL: config.apiBase })

    return { result }
})