## Logging test postmortem

### Introduced bug
When setting up the user session as a user performs an login, we return an error when trying to unmarshal the user object. This happens as we have flipped the `==` operator in the `setUserSession()` function. The bug can be seen here:
```Go
if err == nil {
		return errors.New("Could not marshal json")
	}
```

Our CI/CD setup didn't let us deploy this bug to the production code, as our tests catch the bug when the pull request is created (as seen [here](https://github.com/Lindharden/DevOps/actions/runs/4519248472/jobs/7959629022)). Even when forcefully merging the PR, the code would not be sent to production as our CD script would fail and not deploy the changes. This shows that our CI/CD setup works as intended, however, it's a problem when we intentionally try to implement a bug. Therefore the other team must find the bug in a locally running instance of Kibana.

### Find bug in logs
Logging in to the minitwit application gave an error. Looking in Kibana we see the following log:
```
Mar 25, 2023 @ 14:09:38.054	{"level":"error","timestamp":"2023-03-25T13:09:38Z","caller":"controllers/loginController.go:130","msg":"Could not session","user":"zoinks","stacktrace": ...}
```

The log provided an stacktrace which shows us exactly where the error came from. This means we could easily isolate the problematic component.
