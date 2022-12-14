package server

import (
  "cloud.google.com/go/firestore"
  "github.com/go-chi/chi/v5"
  "github.com/robbailey3/website-api/activities"
  "github.com/robbailey3/website-api/blog"
  "github.com/robbailey3/website-api/config"
  "github.com/robbailey3/website-api/image"
  "github.com/robbailey3/website-api/photos"
  "github.com/robbailey3/website-api/tasks"
)

func setupRoutes(db *firestore.Client, router chi.Router) {
  router.Route("/api", func(r chi.Router) {
    // TODO: Standardise these
    blog.SetupBlogRoutes(db, r)
    photos.InitPhotoRoutes(db, r)
    config.SetupConfigRoutes(r)
    tasks.InitTasksRoutes(r, db)
    image.InitImageRoutes(r, db)
    activities.SetupActivityRoutes(r, db)
  })
}
