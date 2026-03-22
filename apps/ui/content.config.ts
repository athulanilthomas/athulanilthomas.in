import { defineContentConfig, defineCollection, z } from '@nuxt/content'

const baseSchema = z.object({
  type: z.string(),
  title: z.string(),
  id: z.string(),
})

const aboutSchema = baseSchema.extend({
  about: z.string(),
  email: z.string().email(),
  location: z.string(),
})

const educationSchema = baseSchema.extend({
  educations: z.array(
    z.object({
      id: z.string(),
      school: z.string(),
      degree: z.string().nullable(),
      areaOfStudy: z.string(),
      startDate: z.string(),
      endDate: z.string(),
      description: z.string(),
    })
  ),
})

const experienceSchema = baseSchema.extend({
  experiences: z.array(
    z.object({
      id: z.string(),
      company: z.string(),
      position: z.string(),
      location: z.string(),
      startDate: z.string(),
      endDate: z.string().nullable(),
      logoUrl: z.string().url(),
    })
  ),
})

const projectsSchema = baseSchema.extend({
  repositories: z.array(
    z.object({
      id: z.string(),
      nameWithOwner: z.string(),
      description: z.string(),
      url: z.string().url(),
      stargazerCount: z.number(),
      primaryLanguage: z.string(),
      ownerAvatarUrl: z.string().url(),
      isOwned: z.boolean(),
    })
  ),
})

const skillsSchema = baseSchema.extend({
  maxYears: z.number(),
  skills: z.array(
    z.object({
      id: z.string(),
      title: z.string(),
      yearsOfExperience: z.number(),
      logoUrl: z.string().url(),
    })
  ),
})

export default defineContentConfig({
  collections: {
    about: defineCollection({
      type: 'data',
      source: 'portfolio/about.yml',
      schema: aboutSchema,
    }),
    education: defineCollection({
      type: 'data',
      source: 'portfolio/education.yml',
      schema: educationSchema,
    }),
    experience: defineCollection({
      type: 'data',
      source: 'portfolio/experience.yml',
      schema: experienceSchema,
    }),
    projects: defineCollection({
      type: 'data',
      source: 'portfolio/projects.yml',
      schema: projectsSchema,
    }),
    skills: defineCollection({
      type: 'data',
      source: 'portfolio/skills.yml',
      schema: skillsSchema,
    })
  }
})