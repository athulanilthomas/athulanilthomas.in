package dto

import "github.com/google/go-github/github"

type RepoDTO struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	StargazersCount int    `json:"stargazers_count"`
	HTMLURL         string `json:"html_url"`
	AvatarURL       string `json:"avatar_url"`
	PrimaryLanguage string `json:"primary_language"`
}

func toRepoDTO(repo *github.Repository) *RepoDTO {
	if repo == nil {
		return nil
	}

	return &RepoDTO{
		Name:            repo.GetName(),
		FullName:        repo.GetFullName(),
		Description:     repo.GetDescription(),
		StargazersCount: repo.GetStargazersCount(),
		HTMLURL:         repo.GetHTMLURL(),
		AvatarURL:       repo.GetOwner().GetAvatarURL(),
		PrimaryLanguage: repo.GetLanguage(),
	}
}

func ToRepoDTOs(repos []*github.Repository) []*RepoDTO {
	rs := make([]*RepoDTO, 0, len(repos))

	for _, r := range repos {
		rs = append(rs, toRepoDTO(r))
	}

	return rs
}
