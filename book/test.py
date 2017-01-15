#encoding=utf8

import commands

result = commands.getstatusoutput("./helloworld")
print result[1]