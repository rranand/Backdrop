# ✅ Backdrop - Weekly TODO

## Week 1: Setup & Auth
- [x] Define all DB models (User, Task, Token)
- [x] Set up Go project structure (`cmd`, `internal`, `pkg`)
- [x] Initialize Go modules (`go mod tidy`)
- [x] Create routes and basic HTTP server
- [x] Implement user login + token generation
- [x] Store login tokens in database

---

## Week 2: Task Handling & Processing
- [ ] API to request new task (generate upload URL)
- [ ] File upload endpoint
- [ ] Goroutine worker pool to simulate processing
- [ ] Save task status to database
- [ ] Implement cancelation logic (invalidate URL)
- [ ] Add basic task polling endpoint

---

## Week 3: Contexts, Retry & Stability
- [ ] Add context timeout + cancel support to workers
- [ ] Add Redis or caching if needed (optional)
- [ ] Retry failed tasks logic (optional)
- [ ] Validate task transitions (pending → processing → done/cancelled)
- [ ] Add middleware for logging & input validation

---

## Week 4: Cleanup & Wrap-up
- [ ] Test all routes (manual/Postman)
- [ ] Write helpful error messages & logs
- [ ] Create Postgres schema migration script
- [ ] Finalize README with setup instructions
- [ ] Optional: Add Dockerfile for app + db
- [ ] Review codebase and push to GitHub

---
