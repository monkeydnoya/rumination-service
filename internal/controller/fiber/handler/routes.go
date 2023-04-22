package handler

func (s Server) SetupRoutes() {
	api := s.App.Group("/api/blogs")

	api.Get("/", s.GetBlogs)
	api.Get("/:id", s.GetBlogById)

	api.Use(s.DeserializeUserRemote()) //TODO: Rename Deserialize to Middleware
	api.Post("/add", s.CreateBlog)
	api.Patch("/edit", s.UpdateBlog)

	api.Delete("/:id", s.DeleteBlog)
}
