from flask import Flask
import random
 
app = Flask(__name__)
 
@app.route("/")
def hello():
    # Generate random color code
    color = "%06x" % random.randint(0, 0xFFFFFF)
 
    # Generate random emoji
    emojis = ['😀', '😁', '😂', '🤣', '😃', '😄', '😅', '😆', '😉', '😊', '😋', '😎', '😍', '😘', '😗', '😙', '😚', '🙂', '🤗', '🤔', '😌', '😛', '😜', '😝', '😲', '🤪']
    emoji = random.choice(emojis)
 
    # Return HTML with random color background, big emoji in center, and title
    return f'<html><head><title>Random Emoji Generator</title></head><body style="background-color:#{color};display:flex;justify-content:center;align-items:center;height:100vh"><div style="font-size:20em">{emoji}</div></body></html>'
 
if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=5000)
