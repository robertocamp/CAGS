# AIR

## What it is
[Air](https://github.com/cosmtrek/air) is a lightweight and powerful live reloading tool specifically designed for Go applications. It watches for changes in your source code and automatically reloads your application whenever changes are detected. This is especially useful for web development, where rapid feedback and quick iterations can significantly boost productivity.

## What it does
When you’re developing a Go web project, every time you modify code, you typically need to stop the server, recompile, and restart it to test your changes. With Air, this process becomes seamless:
- **Watches for Changes**: Air monitors specified directories and files (usually source code files) for any changes.
- **Rebuilds Automatically**: Upon detecting a change, Air recompiles the application automatically.
- **Restarts the Server**: Air then restarts your application server, allowing you to see the changes instantly in your browser or API client.

In short, Air removes the repetitive steps of manual rebuilding and restarting, making development smoother and faster.

## Installation

### 1. Install Air
To install Air globally, run:
```bash
go install github.com/cosmtrek/air@latest
```
This will install the `air` binary, which you can verify with:
```bash
air -v
```

### 2. air init
#### 1. **Initialize Air**:
    In the root directory of your Go project, you can generate a default configuration file:
    ```bash 
    air init
    ```

    This will create an `.air.toml` file where you can customize the settings if needed.

    2. **Run Air**:
    To start using Air for live reloading, simply run:
    ```bash
    air
    ```

    This will start the Air tool, which will watch your source code for changes and automatically rebuild and restart your application.

    ** in other words, the "air" command actually runs your project locally! **

### Example Configuration (`.air.toml`):
    Here’s a basic example of what the `.air.toml` file might look like:
    ```toml
    # Config file for air

    [build]
    cmd = "go build -o ./tmp/main"
    bin = "./tmp/main"
    full_bin = ""
    include_ext = ["go", "tpl", "tmpl", "html"]
    exclude_dir = ["assets", "tmp"]
    exclude_file = []
    follow_symlink = true
    args_bin = []

    [log]
    color = "auto"
    time = false

    [serve]
    root = "."
    cmd = "./tmp/main"
    delay = 1000
    grace = 5000
    ignore = ["assets", "tmp"]
    ignore_file = []
    watch_dir = ["."]
    watch_ext = []
    exclude_dir = []
    include_dir = []
    ```

    With this setup, you should be ready to start developing your Go project with the convenience of live reloading provided by Air.

### 3. Running Air
Start Air by running the following command from your project’s root directory:
```bash
air
```

This will start watching for file changes, rebuild, and restart the server as needed.

---

## summary

`Air` is a **development tool** rather than a package you import into your Go code. It operates externally to your application code, monitoring for changes and handling automatic rebuilds and restarts. 

Once installed, `Air` acts as a wrapper around your Go project during development, so you just start it up in the terminal. It doesn’t require any imports in your actual Go code, nor does it affect your production environment—it’s purely for local development productivity.



## transcripts

go install github.com/cosmtrek/air@latest                        ▦ CAGS 🔀  air⎪●◦⎥ go ∩ v1.22.5  aws ▲   us-east-2 2024-11-03 13:37

go: github.com/cosmtrek/air@v1.61.1 requires go >= 1.23; switching to go1.23.2
go: github.com/cosmtrek/air@latest: version constraints conflict:
        github.com/cosmtrek/air@v1.61.1: parsing go.mod:
        module declares its path as: github.com/air-verse/air
                but was required as: github.com/cosmtrek/air

887ms robertc ▻ ⊘ air -v                                                     ▦ CAGS 🔀  air⎪●◦⎥ go ∩ v1.22.5  aws ▲   us-east-2 2024-11-03 13:37

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.22.5


