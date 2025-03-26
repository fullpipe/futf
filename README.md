# force-utf8

### Windows1251

#### Scenario:

- original Windows-1251 text has been mis-interpreted as Windows-1252
- then those mis-interpreted characters are themselves turned into their UTF-8 equivalents

#### Example:

`Мировая фанстастика от А до Я` -> `Ìèðîâàÿ ôàíñòàñòèêà îò À äî ß`


## Usage

```go
import github.com/fullpipe/futf

futf.ToUTF8("Ìèðîâàÿ ôàíñòàñòèêà îò À äî ß") // Мировая фанстастика от А до Я
```

