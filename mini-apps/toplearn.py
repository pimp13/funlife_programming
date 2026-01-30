import pprint
from typing import Any, List, Dict


def main():
  print('hi, my name is pouya i\'m developer i love golang python and linux')
  
  # my_list = [1, 2, 3, 4, 5, 6]
  # doubles = map(lambda n: n * 2, my_list)
  # print(doubles)
  # print(list(doubles))
  
  # people: List[Dict[str, Any]] = [
  #   {'first_name': 'pouya', 'last_name': 'gh', 'age': 23},
  #   {'first_name': 'ali', 'last_name': 'puma', 'age': 28},
  #   {'first_name': 'mohammad', 'last_name': 'hassani', 'age': 32},
  #   {'first_name': 'sara', 'last_name': 'moradi', 'age': 18},
  # ]
  # result = [person for person in people if person.get('age') > 18]
  # pprint.pprint(result)
  # a = map(lambda person: person.get('last_name'), people)
  # pprint.pprint(people)
  # bozorg_tar_az_18 = filter(lambda person: person.get('age') > 18, people)
  # pprint.pprint(list(bozorg_tar_az_18))
  
  # formating = pprint.pformat(people)
  # print("formatting -->", formating)
  
  # print(all([1,2,1,0]))
  
  def do_twice(func):
    def wrapper_do_twice():
      func()
      func()
    return wrapper_do_twice
  
  
  @do_twice
  def print_name():
    print('Hello my name is pouya')
    
  print_name()
    
if __name__ == '__main__':
  main()