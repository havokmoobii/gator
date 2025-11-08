# Gator: A Simple Blog Aggregator.

## Requirements:
This application requires Postgres v15 or later and Go v1.25 or later.

## Installation:
Gator can be installed with the command:
```
go install https://github.com/havokmoobii/gator@latest
```

## Setup:
Start the Postgres server in the background
- **Mac**: brew services start postgresql@15
- **Linux**: sudo service postgresql start

Enter the psql shell:
- **Mac**: psql postgres
- **Linux**: sudo -u postgres psql

Run the command:

```
CREATE DATABASE gator;
```

Set the user password (Linux only):
```
ALTER USER YourUserNameHere PASSWORD 'YourPasswordHere';
```

You can now enter "exit" to leave the Postgres CLI.

The start command will need to be run again on a system restart.

Create a file in your home directory named ".gatorconfig.json" and paste the following into it:
```
{"db_url":"postgres://HavokMoobii@localhost:5432/gator?sslmode=disable","current_user_name":}
```

Change "HavokMoobii" to your username from the "ALTER USER" command earlier or, if applicable "YourUsername:YourPassword".

## Usage:
```
gator
```
Runs the program, but will not do anything without a secondary command.

### Commands:
Many commands require a user to be logged in to the database to function.

Before a user can log in they must be registered.

#### register:
```
gator register <username>
```
Registers a new user and log them in.

#### login:
```
gator login <username>
```
Logs a user in if their username is registered in the database.

#### reset:
```
gator reset
```
Removes all entries from the database.

#### users:
```
gator users
```
Lists all registered usernames.

#### addfeed:
```
gator addfeed <feed_name> <feed_url>
```
Adds the feed to the database. The logged in user will also follow the feed.

#### feeds:
```
gator feeds
```
Lists all feeds currently registered in the database.

#### follow:
```
gator follow <url>
```
Follows the given feed for the current user if it exists in the database.

#### unfollow:
```
gator unfollow <url>
```
Unfollows the current feed if it exists in the database and if the current user is following it.

#### following:
```
gator following
```
Lists all feeds the current user is following.

#### agg:
```
gator agg <interval>
```
Checks the least recently updated feed for new posts every "interval".

Interval must be a unit of time (EG: 4s, 10m, 1h).

This command will run forever unless stopped with "CTRL + C". It is intended to be run on a second terminal in the background.

#### browse:
```
gator browse <max_number_of_posts>
```

Lists up to the given number of posts (default 2) for feeds that the current user follows. 

The posts will be ordered from newest to oldest with titles and urls listed for each post.