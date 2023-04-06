# RateRocket.io Slack Bot

RateRocket Slack Bot is a simple Slack bot that fetches the current exchange rate for Bitcoin to a specified currency using the RateRocket.io API.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.17 or higher)
- A Slack workspace with administrative access to create a new bot

### Installing

Clone the repository:

```bash
git clone https://github.com/microwaves/raterocket-slack.git
cd raterocket-slack
```

Create a Slack app and bot:

- Go to the Slack API website and sign in to your workspace.
- Click the "Create New App" button.
- Give your app a name and select your workspace, then click "Create App".
- Go to the "Slash Commands" section and click "Create New Command".
- Set the command to /raterocket, and set the Request URL to the URL where you will host the bot. Add a short description and usage hint, then click "Save".
- Go to the "Install App" section and click "Install App to Workspace". Authorize the app.

Set the environment variables:

- `SLACK_VERIFICATION_TOKEN`: The Slack verification token for your app, found in the "Basic Information" section of your app's settings.
- `PORT`: The port on which the bot should listen for incoming HTTP requests.

Example:

```bash
export SLACK_VERIFICATION_TOKEN="your-token-here"
export PORT=8080
```

### Running the Bot

To run the bot locally:

```bash
go run main.go
```

The bot will now listen for incoming HTTP requests on the specified port.

### Running the Tests

To run the tests:

```bash
go test
```

## Deployment

To deploy the bot, you can use a platform like Heroku or any other platform that supports Go applications and allows you to set environment variables.

## Usage

Once the bot is running, you can use the `/raterocket` command in your Slack workspace, followed by a currency code, to fetch the current exchange rate for Bitcoin to the specified currency.

Example:

```
/raterocket USD
```

This command will return the current exchange rate for Bitcoin to US dollars.

## Maintainers

Stephano Zanzin Ferreira - [@microwaves](https://github.com/microwaves)

## License

RateRocket Slack Bot is released under the BSD license. See [LICENSE](https://github.com/microwaves/MarketHistorian/blob/main/LICENSE).
