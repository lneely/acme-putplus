# acme-put+

Save files with elevated privileges from the Acme editor using PolicyKit.

## Usage

1. Open a file with restricted privileges (e.g., `/etc/hosts`) in Acme
2. Make your edits in the Acme window
3. Add `Put+` to the window's tag line
4. Middle-click `Put+` to save with elevated privileges
5. You'll be prompted for your password via PolicyKit

## Requirements

- PolicyKit (`pkexec` command)
- A graphical PolicyKit authentication agent
- Acme text editor from Plan 9 from User Space

## Installation

```
mk install
```

This installs the `Put+` binary to `$HOME/bin/`.

## How it works

When executed, `Put+` reads the current Acme window content and filename, then uses `pkexec tee` to write the file with elevated privileges. PolicyKit handles the authentication prompt through your desktop environment's authentication agent.