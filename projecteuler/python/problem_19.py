'''
You are given the following information, but you may prefer to do some research for yourself.
  - 1 Jan 1900 was a Monday. 
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine. 
  - A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
'''

rom datetime import date

def is_sunday(day):
    return day.weekday() == 6

def main(initial_date, final_date):
    sundays = 0
    for y in range(initial_date.year, final_date.year+1):
        sundays += sum([1 for m in range(1, 13) if is_sunday(date(y, m, 1))])
    print(sundays)


if __name__ == '__main__':
    initial_date = date(1901, 1, 1)
    final_date = date(2000, 12, 31)

    main(initial_date, final_date)
