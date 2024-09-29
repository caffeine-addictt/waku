# Setup

!!! FAILURE "Not implemented"

    Setup is a planned feature and currently
    serves no purpose.

Setup describes a post-setup script that
is optionally ran for each operating
system.

!!! INFO

    Only paths relative to the directory
    containing `template.json` are allowed.

```json
{
  "setup": {
    // path to a shell script or binary
    // (optional)
    "linux": "",

    // path to a shell script or binary
    // (optional)
    "darwin": "",

    // path to a executable or binary
    // (optional)
    "windows": "",

    // This is the fallback script for unknown
    // operating systems.
    // (optional)
    "*": ""
  }
}
```
