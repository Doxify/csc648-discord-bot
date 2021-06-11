<p align="center">
  <h3 align="center">csc648-discord-bot</h3>
  <p align="center"> 
    A bot used in my software engineering team's Discord server.
  </p>
  <p align="center">
    <img src="https://img.shields.io/badge/go%20report-A-green.svg?style=flat"/>  
  </p>
</p>

---

## I. Built with
* Go 1.16
* [discordgo](https://pkg.go.dev/github.com/bwmarrin/discordgo)

## II. Motivation
I wanted to create something in Go and creating a discord bot in order to make my team's lives easier was a good motivator!

## III. Purpose
The purpose of the bot is to return important information related to the class. It currently only returns the database instance assigned to each user in `data.json`. 

See `.example.data.json` for what `data.json` should look like.

## IV. Commands
These commands will work in any channels the bot has access to.

|Command|Description|
|--|--|
|!db|Returns the connection details of the database instance assigned to you|

## V. Usage

### Prerequisites
* Create a Discord bot [here](https://discord.com/developers/applications) and invite it to your server.
* Copy the contents of `.example.data.json` to a file called `data.json` and fill it out.

### Running the bot
1. Build the application
  ```go
    go build
  ```
2. Launch the application
  ```bash
    ./csc648-discord-bot -t <DISCORD_BOT_TOKEN>
  ```