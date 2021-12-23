# chess-cli
chess-cli is a cli for playing chess against an uci compatible engine written in go

## Usage

    $ chess-cli --help
        Chess-cli is a CLI to play chess against an UCI engine of your choice with the ability to 
        specify depth
        Usage:
          chess-cli [command]

        Available Commands:
          analyze     Get lichess analyze urls on your games
          delete      Delete your chess games from database
          help        Help about any command
          list        List all your games (default will only show on-going games)
          play        Play/Continue a chess game
          puzzle      Play a daily lichess puzzle
          resign      Resign on your chess games

        Flags:
          -h, --help     help for chess-cli
          -t, --toggle   Help message for toggle

## Examples
### Playing against an engine
    $ chess-cli play engine --help
        Start a game against an engine

        Usage:
          chess-cli play engine [flags]

        Flags:
              --color string   choose your color: white/black (default random)
          -d, --depth int      Set the engine depth to search x piles only (default 3)
          -h, --help           help for engine
              --name string    Set the name of the game (default random)
          -p, --path string    Set the UCI chess engine path (required)
  #### Example:
    $ chess-cli play engine -p "D:\Programming\stockfish.exe" --depth 5 --color white

### List games 
    $ chess-cli list --help
      List all your games (default will only show on-going games)

      Usage:
        chess-cli list [flags]

      Flags:
        -a, --all      show all games with board
        -e, --engine   show games' engine configuration
        -h, --help     help for list
        

### Resign games
    $ chess-cli resign --help
      Resign on your chess games

      Usage:
        chess-cli resign [game-names...] [flags]

      Flags:
        -h, --help   help for resign

    $ chess-cli resign tZHpL fromsettingNameFlag getGameNamesFromTheListCommand thisiscasesensitiveAlso
      You resigned on Game "tZHpL" Status: 0-1, Resignation
      You resigned on Game "fromsettingNameFlag" Status: 1-0, Resignation
      Game "getGameNamesFromTheListCommand" doesn't exist.
      Game "thisiscasesensitiveAlso" doesn't exist.


### Delete games
    $ chess-cli delete --help
      Delete your chess games from database

      Usage:
        chess-cli delete [game-names...] [flags]

      Flags:
        -h, --help   help for delete
        
    $ chess-cli delete wzfhz tZHpL 
      Game "wzfhz" is deleted.
      Game "tZHpL" is deleted.
      

### Get lichess analysis url for games
    $ chess-cli analyze --help
      Get lichess analyze urls on your games

      Usage:
        chess-cli analyze [game-names...] [flags]

      Flags:
        -h, --help   help for analyze
    
    $ chess-cli analyze wzfhz tZHpL blahblah
        Analyze Game "wzfhz" on lichess: https://lichess.org/L7vPPR5e
        Analyze Game "tZHpL" on lichess: https://lichess.org/SP6l4EFe
        Game "blahblah" doesn't exist.


### Play a daily puzzle from lichess
    $ chess-cli puzzle -h
      Play a daily lichess puzzle

      Usage:
        chess-cli puzzle [flags]

      Flags:
        -h, --help   help for puzzle

