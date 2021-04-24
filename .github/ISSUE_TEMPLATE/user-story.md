---
name: User Story
about: Create a user story to represents some value to an end user
title: 'Authenticates user credential '
labels: 'Type: User Stories'
assignees: rasatmaja

---

**User story**
> ``As a [type of user], I want [an action] so that [a benefit/a value]``

***As a*** developer who use AAA service to autheticate user, ***i want*** to receive a JWT Token from server when i send user email and his password, ***so that*** i can use those JWT token to navigate user into secure resources 

**Acceptance Criteria**
- Should be able to send user email and his password int the API request. *These fields are mandatory* 
- Should return a JSON object with field **access_token** and **refresh_token**

**Non-Functional Requirements**
- The API should support a load of 50 API requests per second without failing.
- API should process the request and respond within maximum 10 seconds in a cloud-to-cloud setup.

**Issue type** :  

- [ ] 🐞 Possible Bug
- [x] 🦌 New Feature
- [ ] 🤴 Code style update (formatting, renaming)
- [ ] 🏇🏼 Refactoring (no functional changes, no api changes)
- [ ] 🏅 Build related changes
- [ ] 🦧 Documentation content changes

**Descriptions** :  

- Impementation of new RS512 signing method
- Impementation of RSA key
- Unit testing 

**Task** :
- [ ] Generate RS512 pair keys
- [ ] Unit testing
- [ ] Benchmark testing
