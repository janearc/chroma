# contents

- `colourways/` colourways
- `cmd/paratune` adjusts a colourway live
- `libtheme-css` a comically comprehensive color and radiation library
- `readme` lib and tool for maintaining a profile for your eyesight
- `cmd/colourway` allows you to diff spectra
- `cmd/spectra` generate colourways from spectra

# build & use

i use `game` to build and distribute software for a variety of reasons. all of
the commands that are built have quite substantial help and docs. i'd love
to hear from you if you use this.

```bash
$ sh game/bootstrap.sh     # bootstrap
$ ./game/bin/game build    # compile
$ ./game/bin/game check    # tests
$ ./game/bin/game clean    # clean
$ ./game/bin/game install  # install
```
