<h1 align="center">go-snake 🐍</h1>
<p align="center">
  Go CLI implementation of Snake game, using channels.<br/>
  NB: this code works only with ANSI terminals.
</p>
          
<h2 align="center">Demo</h2>
<p align="center">
  <img alt="Demo" src="https://github.com/mikyll/go-snake/blob/main/gfx/cli-snake.gif"/><br/>
  Resize the terminal window to change the "difficulty".
</p>

### Fruit Spawn Algorithm
```go
if b.xy[f.x][f.y] == HEAD || b.xy[f.x][f.y] == BODY {
  if f.x == 1 || f.x == BL-1 {
    f.x = BL / 2
  }
  if f.y == 1 || f.y == BH-1 {
    f.y = BH / 2
  }
  found := false
  j := f.y
  i := f.x
  for f.y += 1; f.y != j; f.y++ {
    for f.x += 1; f.x != i; f.x++ {
      if b.xy[f.x][f.y] != HEAD && b.xy[f.x][f.y] != BODY && b.xy[f.x][f.y] != BORDER {
        found = true
        break
      }
      if f.x == BL-1 {
        f.x = 1
      }
    }
    if found {
      break
    }
    if f.y == BH-1 {
      f.y = 1
    }
  }
}
```

### References
- [Read character from stdin without pressing enter](https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter/): [Windows](https://stackoverflow.com/a/70627571), [UNIX](https://stackoverflow.com/a/17278776)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
