# JobSearch.wtf

> Because the job market is BS

A comprehensive platform empowering software engineers to effortlessly manage their job search journey. From tracking job applications and interview schedules to capturing insightful notes and conducting retrospectives. It streamlines the entire process, enabling engineers to navigate their career path with clarity and confidence.

## TODO:

### Stack

- [x] Go
- [x] Chi
- [ ] HTMX
- [ ] TailwindCSS
- [ ] Turso SQLite
- [ ] Deployed to VPS - Dockerfile

### Features

- [ ] User login and auth (oauth: google/github)
- [ ] AI: analysis of interview reviews and notes for improvements
- [ ] UUID? Is that still best practices?
- [ ] Figure out turso db migrations on deploy so I dont have to do it locally
- [ ] CRUD: companies, jobs, interviews, notes, contact info, scheduling, etc...
- [ ] Add company table
  - [ ] adjust db to associate relation with company on applications
  - [ ] Ranking of company for that job application
    - company_id
    - job_application_id
    - composite index on company_id + job_application_id
    - rank
- [ ] Add jobs table
  - [ ] description
  - [ ] comp
  - [ ] company foreign key
  - [ ] Notes
  - [ ] equity
  - [ ] benefits
  - [ ] Ranking of company
  - [ ] Recruiter Contact
  - [ ] Interview Dates
  - [ ] Expected Timeline (next contact? hiring timeline?)
- Add Interview table
  - [ ] Notes
  - [ ] Retrospective
  - [ ] Areas you did well
  - [ ] Areas you could improve
  - [ ] Interviewers Names & Contact Info
