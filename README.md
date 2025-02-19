<img style="margin-inline: auto;" alt="golang icon" height="80px" width="auto" src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg" />

# Build Your Own Shell
Questo progetto è un punto di partenza per soluzioni in Go alla sfida "Build Your Own Shell". L'obiettivo è costruire una shell conforme a POSIX capace di interpretare comandi shell, eseguire programmi esterni e comandi builtin come `cd`, `pwd`, `echo` e altro ancora. Durante il percorso, imparerai a gestire il parsing dei comandi shell, REPLs, comandi builtin e molto altro.

## Funzionalità

### Comandi Builtin
- `echo`: Stampa gli argomenti passati.
- `exit`: Termina la shell.
- `pwd`: Mostra la directory corrente.
- `cd`: Cambia la directory corrente.
- `type`: Verifica se un comando è builtin o esterno.

### Esecuzione di Comandi Esterni
La shell è in grado di eseguire comandi esterni presenti nel PATH di sistema.

### Redirezione di Output
Supporta la redirezione dell'output standard e degli errori:
- `>`: Redirezione dell'output standard.
- `>>`: Append dell'output standard.
- `2>`: Redirezione degli errori.
- `2>>`: Append degli errori.

### Autocompletamento
Supporta l'autocompletamento dei comandi builtin e dei comandi esterni.

### Parsing degli Argomenti
Gestisce correttamente argomenti con spazi, virgolette singole e doppie, e caratteri di escape.

## Struttura del Progetto

- `cmd/myshell/main.go`: Punto di ingresso della shell.
- `cmd/myshell/cmds.go`: Implementazione dei comandi builtin.
- `cmd/myshell/helpers.go`: Funzioni di supporto per l'esecuzione dei comandi e la gestione dell'input.
- `cmd/myshell/legacy.go`: Codice legacy commentato per riferimento.
- `cmd/myshell/helpers_test.go`: Test per le funzioni di supporto.