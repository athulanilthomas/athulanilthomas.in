
import type { CachedEventHandlerOptions } from 'nitropack'

interface RepoDTO {
    name: string
    full_name: string
    description: string
    stargazers_count: number
    html_url: string
    avatar_url: string
    primary_language: string
}

const GITHUB_USERNAME = 'athulanilthomas'

const CACHE_CONFIG: CachedEventHandlerOptions = {
    getKey: () => 'repos',
    maxAge: 60 * 60 * 24 * 7 // 1 Week
}

export default defineCachedEventHandler(async (event) => {
    const config = useRuntimeConfig(event)
    const repos = await $fetch<RepoDTO[]>(`/github/repos`, { baseURL: config.apiBase })

    const repositories = repos.map((repo, index) => {
        const owner = repo.full_name.split('/')[0] ?? ''

        return {
            id: String(index + 1),
            nameWithOwner: repo.full_name,
            description: repo.description,
            url: repo.html_url,
            stargazerCount: repo.stargazers_count,
            primaryLanguage: repo.primary_language || null,
            ownerAvatarUrl: repo.avatar_url,
            isOwned: owner.toLowerCase() === GITHUB_USERNAME.toLowerCase()
        }
    })

    return { repositories }
}, CACHE_CONFIG)