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

3. Set api key in `assistant.json`

4. Run the assistant server

```shell
go run .
```

5. Test the assistant

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
| `assistant.api_key`                       | OpenAI api key                                | `sk-xxx`                                                |
| `assistant.model`                         | OpenAI model                                  | `gpt-3.5-turbo`                                         |
| `fiber.xxx`                               | [Fiber](https://gofiber.io/) related parameters    |         see the official document                  |
