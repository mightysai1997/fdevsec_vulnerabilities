# Instead of this …
#cursor.execute(f"SELECT admin FROM users WHERE username = '{username}'");
# ...do this...
#cursor.execute("SELECT admin FROM users WHERE username = %(username)s", {'username': username}); 
import cursor
import yaml
cursor.execute(f"SELECT admin FROM users WHERE username = '{username}'");
cursor.execute(f"SELECT admin FROM users WHERE username = '{username}'");
cursor.execute(f"SELECT admin FROM users WHERE username = '{username}'");
Data = yaml.load(input_file, Loader=yaml.SafeLoader);
