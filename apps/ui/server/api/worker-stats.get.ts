// TODO: move to Go server
export default defineEventHandler(() => {
  return {
    goroutines: Math.floor(Math.random() * 3) + 1,
  }
})
