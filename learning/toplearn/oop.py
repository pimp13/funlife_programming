
from pprint import pprint
from typing import Any, Dict, List


class MyClass:
  timsar: str
  activeUserCount: int = 0
  
  def __init__(self, timsar: str) -> None:
    self.timsar = timsar
    self.__key = "this-is-key"
    """
    self baraye har object ke az class sakhte mishe 
    amma ClassName. baraye hameye object haye class hast va moshtarak hast
    """
    MyClass.activeUserCount += 1
    
mc = MyClass("dev pouya gh")
tow = MyClass('hello')
# print (mc.__key)
# print(dir(mc))
# print(mc.activeUserCount)
# print(tow.activeUserCount)


# --------------------------------------------------------

class ChatRoom:
  __activeRoomCount = 0
  
  def __init__(self, username: str, password: str) -> None:
    try:
      validateResult = self.__validatePassword(password)
      for validate in validateResult:
        if not validate.get('ok'):
          pprint(validateResult)
          raise ValueError(validate.get('message', 'validation is failed!'))
      self.username = username
      self.password = password
      ChatRoom.__activeRoomCount += 1  
    except ValueError as err:
      print(f'error is: {err}')
    
    
  def getActiveRoomCount(self) -> int:
    return ChatRoom.__activeRoomCount
  
  def getUserInfo(self) -> str:
    return f'{self.username}'
  
  def __validatePassword(self, passwd: str):
    hasNumber = False
    for char in passwd:
      if char.isdigit():
        hasNumber = True
        break
    
    hasStdLen = False
    if len(passwd) >= 8:
      hasStdLen = True
    
    pprint(f'hasStdLen: {hasStdLen} , hasNumber: {hasNumber}')
    
    validationStates: List[Dict[str, Any]] = []
    if not hasNumber:
      validationStates.append({'ok': False, 'message': 'password shooma bayad daraye number va string bashe'})
    
    if not hasStdLen:
      validationStates.append({'ok': False, 'message': 'password shooma bayad 8 ya bishtar character bashe'})
    
    if len(validationStates) == 0:
      validationStates.append({'ok': True, 'message': 'validation is ok'})
            
    return validationStates

# newChatRoom = ChatRoom('dev.pouya.gh', 'sd3fdsfds')
# newChatRoom_2 = ChatRoom('ali.hassani', '87654321')
# print (newChatRoom.getActiveRoomCount())


# --------------------------------------------------------


class User:
  def __init__(self, username, family) -> None:
    self.username = username
    self.family = family
    
  """ return instance sakhte shode az class """
  def __repr__(self) -> str:
    return f'{self.username} {self.family}'
    
  @classmethod
  def get_class_method(self):
    print(self)
    
  def test(self):
    print(self)
    
  @classmethod
  def from_string(cls, name_string: str):
    fullname = name_string.split(',')
    return cls(fullname[0], fullname[1])
    
# user = User.from_string('pouya,gh')
# print(f'username: {user.username} , family: {user.family}')
# user1 = User('devpouyagh')
# User.getClassMethod()
# user1.test()
print(User("pouya", "gh"))