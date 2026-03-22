// About Section
export interface AboutData {
  about?: string
  email?: string
  location?: string
}

// Experience Section
export interface Experience {
  id: string
  company: string
  position: string
  location?: string | null
  startDate: string
  endDate?: string | null
  description?: string | null
  logoUrl: string
}

export interface ExperienceData {
  experiences: Experience[]
}

// Skills Section
export interface Skill {
  id: string
  title: string
  yearsOfExperience: number
  logoUrl: string
}

export interface SkillsData {
  skills: Skill[]
  maxYears: number
}

// Education Section
export interface Education {
  id: string
  school: string
  degree?: string | null
  areaOfStudy: string
  startDate: string
  endDate?: string | null
  description?: string | null
}

export interface EducationData {
  educations: Education[]
}

// Projects Section
export interface Repository {
  id: string
  nameWithOwner: string
  description: string | null
  url: string
  stargazerCount: number
  primaryLanguage: string | null
  ownerAvatarUrl: string
  isOwned: boolean
}

export interface ProjectsData {
  repositories: Repository[]
}
