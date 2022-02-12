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
To prevent the fruit from spawning over the snake, we first get a pair of random indexes ```f.x``` and ```f.y``` and we check if that location is already occupied by the snake.
If so, we save those indexes and iterate over the columns (```f.y```) and rows (```f.x```) of the board, starting from ```f.y+1``` and ```f.x+1```.
At each cycle we check if the new location is still occupied by the snake:
- If not we've found our fruit coordinates;
- otherwise we increase ```f.x``` by one (or ```f.y```, for the external loop).
When the index reaches the border (```BL``` or ```BH```), we reset it to 1, so it can make a complete cycle, and stop when they get to ```f.x``` or ```f.y``` (remember we started from ```f.x+1``` and ```f.y+1```).

This way we got a random starting point for our fruit spawn location and, in case it's not valid, we iteratively check for free coordinates.
```go
f.x = rand.Intn(BL-2) + 1
f.y = rand.Intn(BH-2) + 1

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
