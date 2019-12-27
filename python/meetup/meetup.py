import calendar


class MeetupDayException(Exception):
    pass


def meetup(year, month, week, day_of_week):
    days_of_week = dict(zip(list(calendar.day_name), range(7)))

    cal = calendar.Calendar()

    days = [
        day for day in cal.itermonthdates(year, month)
        if day.weekday() == days_of_week[week] and day.month == month
        ]

    if day_of_week == 'teenth':
        meetup = [d for d in days if d.day >= 13 and d.day <= 19][0]
    elif day_of_week == 'last':
        meetup = days[-1]
    else:
        meetup = days[int(day_of_week[0])-1]

    return meetup
