# Create a function which answers the question "Are you playing banjo?".
# If your name starts with the letter "R" or lower case "r", you are playing banjo!

# The function takes a name as its only argument, and returns one of the following strings:

# name + " plays banjo" 
# name + " does not play banjo"
# Names given are always valid strings.

def are_you_playing_banjo(name):
    if name[0] == 'r' or name[0] == 'R':
        return name + " plays banjo" 
    else:
        return name + " does not play banjo"

print(are_you_playing_banjo("dom"))
print(are_you_playing_banjo("rom"))

# https://www.codewars.com/kata/53af2b8861023f1d88000832/train/python

# Rather than checking the characters separately, make it lower case and just check 'r'

def areYouPlayingBanjo(name):
    return name + (' plays' if name[0].lower() == 'r' else ' does not play') + " banjo";