REQUIREMENTS : pandas ewm library

The underlying market is modeled like sinusoidal function

Quantopian platform

Brock et al. (1992)

market inefficiency in which prices takes swings from their fundamental values and (2) markets are efficient and 
the predictable variation can be explained by time-varying equilibrium returns.

Bauer and Dahlquist (2001)

Naved and Srivastava (2015) -> moving average(simple, triangular, exponential, variable, and weighted) , 3 rules (Direction of the moving average, price and moving average crossover and 
crossover of two moving averages with different periods on stocks)
-> profitable ,  simple moving average performed better 


Park and Irwin (2007)

Quantopian 
platform (Algorithm IDE), which is a Python development environment designed to help for 
coding trading strategies using the Algorithm API.



: Its best to use these while using moving average exponential -> span(n-day moving average exponential),centre of mass(C=(S−1)/2) ,
span provides an intuitive way to think about how many observations are significantly weighted in the EWMA. Directly tied to the idea of a "lookback period.
com gives control over the relative weighting of recent vs. past data.Balances the influence of older vs. newer data by specifying where the "center of mass" of the weights lies.
half life Defines the time it takes for the weight of a data point to reduce to half its original value.Intuitive when thinking in terms of "how quickly" the influence of old data decays.
alpha 𝜆  Directly specifies the smoothing constant, controlling the decay rate of older data.

