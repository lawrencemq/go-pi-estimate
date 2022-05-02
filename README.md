# go-pi-estimate

Estimates Pi using a Monte-Carlo method and Go Routines

Idea and mathematics based on [this TikTok video](https://www.tiktok.com/foryou?_r=1&_t=8RwOVSKaLIU&feed_mode=v1&is_from_webapp=v1&item_id=7086552077370985770#/@mathletters/video/7086552077370985770).

We can estimate Pi by applying a normal distribution over the area of a circle and a square. The ratio that a point randomly falls inside the circle multiplied by the (known) area of the square yields an estimate.

The more iterations added to the distribution, over time, should give a more accurate estimation.

## Usage

```
usage: go-pi-estimate [-h|--help] [-i|--iterations <integer>] [-t|--threads
                      <integer>] [-v|--verbose]

                      Estimates Pi through a Monte Carlo method

Arguments:

  -h  --help        Print help information
  -i  --iterations  the maximum number of iterations to use to estimate Pi per
                    thread. Default: 10000
  -t  --threads     the number of threads to use. Default: 4
  -v  --verbose     print verbose output during calculation. Default: false
```
