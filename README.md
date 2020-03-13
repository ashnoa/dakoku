# Dakoku

A CLI tool for your time management.

## Description

`dakoku` is a CLI tool that show your work time for each tasks. It create database to store your tasks and datas, and show it on your terminal.


## Features

- Initialize SQLite3 Database for your tasks.
- Create tasks.
- Log your work times.


## Requirement

- Go
- SQLite3


## Usage

First, you have to initialize your database at current directory. It makes `db.sql` file.

```
    $ dakoku init
```

And then, you can create your task.

```
    $ dakoku create TASK_TITLE
```

Check your today's tasks with below command. It shows `TASK_ID | TASK_TITLE | WORK_TIME | ("Now Doing" if task is processing) `.

```
    $ dakoku show
```

`show` has some options.

```
    $ dakoku show --days 1 # Show tasks from 1 day ago. You can also use short option -d.
    $ dakoku show --days 0 # This is equal to `dakoku show`.
    $ dakoku show --all # Show all tasks. You can also use short option -a.
```

Start a task.

```
    $ dakoku start TASK_ID
```

Stop a task.

```
    $ dakoku stop TASK_ID
```


## Installation

```
    $ go get https://github.com/ashnoa/dakoku
```


## Author

[ashnoa](https://twitter.com/ashnoa)


## Licence

[MIT](https://github.com/ashnoa/dakoku/blob/master/LICENSE)
