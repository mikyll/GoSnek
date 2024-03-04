<div align="center">

  [![GPL-3.0 License][license-shield]][license-url]
  [![Size][size-shield]][size-url]
  [![Issues][issues-shield]][issues-url]
  [![Stars][stars-shield]][stars-url]\
  [![Go][go-shield]][go-url]
  [![VS Code][vs-code-shield]][vs-code-url]

# GoSnek 🐍
<p>
  Go CLI implementation of Snake game, using channels.<br/>
  NB: this code works only with ANSI terminals.

</div>
          
<h2 align="center">Demo</h2>
<p align="center">
  <img alt="Demo" src="https://github.com/mikyll/GoSnek/blob/main/gfx/cli-snake.gif"/><br/>
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

<div align="center">

[![LinkedIn][linkedin-shield]][linkedin-url]
[![GitHub followers][github-shield]][github-url]
  
</div>

[downloads-shield]: https://img.shields.io/github/downloads/mikyll/GoSnek/total
[downloads-url]: https://github.com/mikyll/GoSnek/releases/latest
[license-shield]: https://img.shields.io/github/license/mikyll/GoSnek
[license-url]: https://github.com/mikyll/GoSnek/blob/main/LICENSE
[size-shield]: 	https://img.shields.io/github/repo-size/mikyll/GoSnek
[size-url]: https://github.com/mikyll/GoSnek
[issues-shield]: https://img.shields.io/github/issues/mikyll/GoSnek
[issues-url]: https://github.com/mikyll/GoSnek/issues
[stars-shield]: https://custom-icon-badges.herokuapp.com/github/stars/mikyll/GoSnek?logo=star&logoColor=yellow&style=flat
[stars-url]: https://github.com/mikyll/GoSnek/stargazers

[go-shield]: https://img.shields.io/badge/Go-%2300ADD8.svg?logo=go&logoColor=white
[go-url]: https://go.dev/
[vs-code-shield]:  https://img.shields.io/badge/VS%20Code-0078d7.svg?logo=visual-studio-code&logoColor=white
[vs-code-url]: https://code.visualstudio.com/

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?logo=linkedin&colorB=0077B5
[linkedin-url]: https://www.linkedin.com/in/michele-righi/?locale=en_US
[github-shield]: https://img.shields.io/github/followers/mikyll.svg?style=social&label=Follow
[github-url]: https://github.com/mikyll
