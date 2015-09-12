# Coup-sim

Coup-sim is a card-game simulator designed for specialized strategy exploration. It is intended to be a method by which the true game theory champions may reveal themselves.

## Making your own Strategy
In order to make a strategy, follow these steps.

1. Copy the "RandomStrategy" folder in `model/strategies/`. Paste it as a new folder into the strategies directory and rename it to (yourstrategyname)Strategy. Also rename the file inside to (yourstrategyname)Strategy.go
2. Modify the GetStrategyName() method in your strategy so that it returns your chosen strategy alias. This can also be set in the constuctor as it is for the RandomStrategy.
3. Add your strategy alias to the method in `coup/strategyFactory.go`. This string will be used to select your strategy on the command line. Just copy one of the existing case statements and substitute in your strategy name. Note that the compiler won't be able to find your strategy unless you import it at the top. Import it just like the others.

Did it work?

Run the program. Choose the number of players as 2. Type your strategy name alias when asked to choose a strategy. If the processor was able to find your strategy, you will be asked to give an in-game name. Choose a name. Repeat the strategy and name choice for the second player. Choose to play the game 1 time.

No errors? Yay! It worked! Any changes made to the strategy file will be reflected by this strategy's choices in-game.

Make sure to run `go install` to compile your changes to the strategy code. Also consider writing some go tests if your strategy gets significantly complicated.

** Players CANNOT have the same in-game name

## Strategy Rules
- A strategy must be able to implement the Strategy interface. (Go will not compile if it does not)
- A strategy may not import any twoodhouse packages aside from log and deck
- A strategy will be disqualified in-game if it does not comply with the data return requirements specified in strategy.go

## Tips

- Use the "console" strategy to play 1 on 1 with your strategy for functional testing. It is just a strategy which polls the command line for all its actions.
- Check out the log.go file for how to parse JSON log files
- Strategies are passed by value in the controller, so you should implement each decision as independent of all others it may have previously made. Use data from the log if you want to consider past actions.
