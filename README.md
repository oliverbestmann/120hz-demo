# 120hz-demo

This demo checks whether your display runs at 120 Hz.

It draws four circles at different positions, updating their locations every 1/120th of a second.

On a 120 Hz display, the circle appears at all four positions with similar brightness.
On a 60 Hz display, every other position is skipped, so the circle appears in only two locations.

## Run

Run the demo:
```sh
go run github.com/oliverbestmann/120hz-demo@latest
```

Check https://ebitengine.org/en/documents/install.html if you're missing dependencies at compile time.
