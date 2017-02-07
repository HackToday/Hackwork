import json

# will result in bar being a unicode object; without ensure_ascii=False, bar is a str.
# source: http://stackoverflow.com/questions/20833772/unicode-string-display-on-django-template
bar = json.dumps(foo, ensure_ascii=False)
