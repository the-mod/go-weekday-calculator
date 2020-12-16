# go-weekday-calculator

Calculates the Date for a given Weekday based on the current week and formates it according to the given format.

## Examples:
Current Calenderweek is `CW51`
`GetFormatedDateString("Monday", "dd.MM.YYYY")` => `14.12.2020`
`GetFormatedDateString("Sunday", "dd.MM.YYYY")` => `20.12.2020`

## Supported Formats:
* `dd` for days
* `MM` for months
* `yy` or `yyyy` for years

The delimeter is up to you.
So e.g. it works with `yyyy-MM-dd` or `MM/dd/yy`.