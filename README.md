# gitlab-hooks-listener
simple golang application for receiving gitlab hooks and parse them and send them to telegram channel 

### Quick start
- copy and create new config file `./config` folder from `./config/config_example.yml` and name it `config.yml`
- run `docker-compose up -d`
- create a project in config file 
- set gitlab webhook on route `/projects/project-name/events`

### Config 
file `config/config.yml` 
- `http-server.port` http server port 
- `projects.*` register your projects here 
- `project.*.gitlab-secret` [gitlab secret token](https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#validate-payloads-by-using-a-secret-token) 
- `project.*.telegram-token` [telegram bot api token ](https://core.telegram.org/bots/api#authorizing-your-bot)
- `project.*.telegram-chat` [your telegram chat id](https://core.telegram.org/bots/api#chat) 
- `project.*.events` enable or disable events that you want to receive here (push,merge-request,...)

### Gitlab available events  
- comment
- deployment
- feature-flag
- group
- issue
- job
- merge-request
- pipline
- push ✅
- release
- sub-group
- tag
- wiki-page 

