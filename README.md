# assistant

Assitant built with ChatGPT!

## Quick start

1. Clone the repo

```shell
git clone https://github.com/bjwswang/assistant
```

2. Install dependencies

```shell
cd assistant
go mod tidy
```

3. Set configuration in `assistant.json`

see [Configurations](## Configurations) for more details

4. Build the assistant server and CLI

```shell
make build
```

When build is done,you will get two binaries in `bin` directory.

- `assistant` is the assistant server
- `acli` is the assistant CLI which can interact with the assistant server

5. Start the assistant server

```shell
./bin/assistant --config assistant.json
```

6. Test the assistant

```shell
curl -XGET http://localhost:9999
```

Output should be

```text
Welcome to AI Assistant 👋!
```

## Configurations

| Parameter                                 | Description                                   | Default                                                 |
|-------------------------------------------|-----------------------------------------------|---------------------------------------------------------|
| `addr`                       | The address which assistant server will watch                           | `:9999`                                                |
| `assistant.api_key`                       | OpenAI api key                                | `sk-xxx`                                                |
| `assistant.chat.xxx`                         | OpenAI model configuration for `Chat`                                | `model:gpt-3.5-turbo` `temperature:0.5` `max_tokens:100`                                         |
| `assistant.unit_test.xxx`                         | OpenAI model configuration for generating unit tests                                | `model:gpt-3.5-turbo` `temperature:0.5` `max_tokens:100`                             |
| `fiber.xxx`                               | [Fiber](https://gofiber.io/) related parameters    |         see the official document                  |

## APIs

1. Normal chat

- path: `/chat`
- method: `post`
- paramters:
  - `question`: the question you want to chat with the assistant

2. Generate unit test

- path: `/ut`
- method: `post`
- paramters:
  - `code`: the code you want to used for generating unit tests

## CLI

1. Chat with assistant

```shell
./bin/acli --server http://localhost:9999 chat --question "What is AI assistant in 10 words?"
```

2. Generate unit tests

```shell
./bin/acli --server http://localhost:9999 ut --file {filepath_to_source_code_}"
```

## Contribute to assistant

Welcom to contirbute to this AI assistant!
