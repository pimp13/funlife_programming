
# def colorize(text: str, color: str) -> str:
#   if type(text) is not str:
#     raise TypeError('invalid type text')
#   return f"hi, this is {text} with color {color}"
# tc = colorize(3242, "red")
# print (tc)



"""
try, except, else and finally 
"""

def get_number_from_input():
  try:
    num = int(input('enter your code number: '))
  # except ErrorType or khali
  # except ValueError:
  except:
    return 'please enter a number, your input value is not a number!'
  else:
    # agar be error nakhord in run mishe, agar vaarede except nashe else run mishe   
    return num
  finally:
    print ('man hamishe run misham agar error bashe ya nabashe!')

v = get_number_from_input()
print (v)