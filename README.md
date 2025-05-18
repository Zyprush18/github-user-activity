# Github User Activity

A simple command-line interface (CLI) application to fetch and display the recent GitHub activity of any user directly in your terminal.

This project is ideal for practicing:

* Working with REST APIs
* Handling JSON data
* Building simple CLI tools

## 🚀 Features

* Accepts a GitHub username as a command-line argument.
* Fetches recent public activity of the user using GitHub's API.
* Displays readable activity summaries in the terminal.
* Gracefully handles errors (e.g., invalid username or API failures).
* Does **not** use any external libraries or frameworks for HTTP requests or JSON parsing.

## 📆 Example Usage

```bash
$ github-activity kamranahmedse
```

### Sample Output

```
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
- Created kamranahmedse/awesome-tools
```

## 🧹 Supported Activity Types

The CLI currently supports and formats the following GitHub event types:

* **PushEvent** → `Pushed <n> commits to <repo>`
* **IssuesEvent (opened)** → `Opened a new issue in <repo>`
* **WatchEvent (starred)** → `Starred <repo>`
* **CreateEvent (repository)** → `Created <repo>`

Other event types are ignored for simplicity.

## 🔗 GitHub API Endpoint Used

This CLI uses the following GitHub API endpoint:

```
https://api.github.com/users/<username>/events
```

Example:

```
https://api.github.com/users/kamranahmedse/events
```

Note: GitHub's public API is rate-limited. If you encounter errors, it might be due to exceeding the rate limit.

## 🛠️ Installation & Build

1. Clone the repository:

```bash
git clone https://github.com/yourusername/github-activity.git
cd github-activity
```

2. Build the project:

**If using Go:**

```bash
go build -o github-activity main.go
```

3. Run the CLI:

```bash
./github-activity <github-username>
```

## Project Idea

This project is inspired by the <a href="https://roadmap.sh/projects/github-user-activity">Github User Activity Project on Roadmap.sh</a>, designed to practice building CLI apps using Go.
