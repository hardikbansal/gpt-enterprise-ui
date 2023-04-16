# gpt-enterprise-ui

Aim of this repository is to provide on-premise server for chat-gpt rather than local server. 
It helps companies/organisation to provide access to Chat-Gpt to their employees without them(employees) actually purchasing the GPT subscription individually.
It stores the usage data at the server in postgres backend. It also helps analyze the data of how people are using the chat-gpt at the workspace to optimize the workflow. 
I recommend to use this only for work purpose and not for personal use as this stores data regarding user usage in plain text (for now).

## Prerequisites

Running this project requires following dependencies

1. Docker (to orchestrate container)
2. Docker Compose (to manage containers)
3. Google Auth client id
4. Open AI key

## Steps  to install

1. Update the `.env` file with three parameters
   1. `API_KEY`: OpenAi Key
   2. `GOOGLE_LOGIN_CLIENT_ID` : Read here to get this. This is used to facilitate login via gmail.
   3. `EMAIL_DOMAIN`: Email for your organisation for example 
2. Run `docker compose up -d`. Wait for some time to get it live. 
3. For production purpose make `IS_DEBUG=false` in `.env` file

## Features supported

1. Feature to store conversation with chat-gpt.
2. Template feature with any number of variables.

### Features to be supported

1. Stream tokens
2. Loader screen after submitting query. For now it takes time to load response, so one has to wait for the response without any indication.
3. Other login methods to be supported
4. Important langchain features to be added.