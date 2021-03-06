# taue is reporting developer activity (GitHub, GitLab, etc...) to Slack

* Go 1.8~
* It is assumed that work on Heroku

## Usage
1. Add user information to `taue/resources/Users.json`
    * id int (required): user id
    * name string (required): user name (sample Twitter)
    * slackName string (required): Slack user name
    * githubName string (required): GitHub user name (https://github.com/xxxxxxx)
    * githubToken string (optional): [GitHub Personal access token](https://github.com/settings/tokens)
    * gitlabId int (optional): GitLab user id
    * gitlabToken string (optional): [GitLab Personal access token](https://gitlab.com/profile/personal_access_tokens)

```javascript
[
  {
    "id": 0,
    "name": "yutailang0119",
    "slackName": "yutaro",
    "githubName": "yutailang0119",
    "githubToken": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "gitlabId": 01234567,
    "gitlabToken": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  },
  {
    "name": "d_date",
    "slackName": "d_date",
    "githubName": "d-date",
  }
]
```

2. Add [Slack WebHook URL](https://api.slack.com/incoming-webhooks)Endpoint (https://hooks.slack.com/services/xxxxxxxxxxxxxxxx) to env
**Example for Heroku**

```bash
$ heroku config:set SLACK_WEBHOOK_ENDPOINT="xxxxxxxxxxxxxxxx" --app "app_name"
```

3. Set a scheduler

## taue support private git repositories on GitHub and GitLab

### GitHub
Please get your [private access token](https://github.com/settings/tokens) and checked `repo:status`.
Add this token  to `githubToken` on Users.json.

### GitLab
GitLab API is a little unique.
First please get your [private token](https://gitlab.com/profile/personal_access_tokens).
* `Expires at`: token expired date. If it blank is indefinitely.
* `Scopes`: check `api Access your API`
Second call this API (https://gitlab.com/api/v3/projects?private_token=xxxxxxxxxxxxx) with token.
```javascript
[
    {
    ...
        "owner": {
            "name": "",
            "username": "",
            "id": 123456,
            "state": "active",
            "avatar_url": "",
            "web_url": ""
        }
    ...
    },
    { ... }
]
```
This id is `gitlabId`.
Finally add this id and token to `gitlabId`, `gitlabToken` on Users.json.

## License
[taue](https://github.com/yutailang0119/taue) is released under the [Apache License 2.0](LICENSE).
