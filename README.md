# Snippetbox

A web application for sharing text snippets.

## Project Structure

```
.
├── cmd/
│   └── web/          # Application entry point and HTTP handlers
├── ui/
│   ├── html/         # HTML templates
│   │   ├── pages/    # Page templates
│   │   └── partials/ # Reusable template partials
│   └── static/       # Static assets
│       ├── css/
│       ├── js/
│       └── img/
```

## Environment Variables

Create a `.env` file in the project root with the following variables:

```env
ADDR=:4000
DSN=web:pass@/snippetbox
```

### Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `ADDR` | HTTP server address and port | `:4000` |
| `DSN` | MySQL data source name | `root:123@/snippetbox` |

### DSN Format

The MySQL DSN follows this format:

```
[username]:[password]@[protocol]([host]:[port])/[database]
```

## Contributing

### Commit Messages

This project follows commit message prefix conventions. See [docs/devops/commit-message.md](docs/devops/commit-message.md) for the full list of prefixes.
