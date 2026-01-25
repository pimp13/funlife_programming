import sys
import requests
from bs4 import BeautifulSoup

if len(sys.argv) < 2:
  print("No arguments provided")
  exit(1)

word = sys.argv[1]
url = f'https://www.google.com/search?sca_esv=09a4dc9e07e27686&sxsrf=ANbL-n5TvPz9gQvQlMwtdM8MjnlqHaOWTw:1769170109522&udm=2&fbs=ADc_l-aN0CWEZBOHjofHoaMMDiKpaEWjvZ2Py1XXV8d8KvlI3vWUtYx0DZdicpfE1faGYenqWn-q4MFiFFtvJjTKeAVxBf9XF8ByrMpEedseJb6C24e7QdJQdIE3TPpl5mEwf0EBCRLNeEhQy5amsEYIcoTye3rrZrd3IP3OYha6_rH_GIVOU8GK5eecabclKqwVhxmAmgM5&q={word}'

html = requests.get(url).text
soup = BeautifulSoup(html, 'html.parser')

try:
  img = soup.find_all('img')[2]['src']
  print (img)
except:
  print('can not get img from google photos')